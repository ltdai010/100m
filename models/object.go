package models

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

type Object struct {
	ObjectId   int32  `json:"object_id"`
	ObjectName string `json:"object_name"`
	Score      int32  `json:"score"`
	UserID     int32  `json:"user_id"`
}

const OBJECT = "object"
const OBJECT_COUNTER = "object_counter"

func AddOne(object Object) (string, error) {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, uint32(object.ObjectId))
	b, err := json.Marshal(&object)
	it := generic.TItem{
		Key:   bs,
		Value: b,
	}
	_, err = BigsetIf.BsPutItem2(OBJECT, &it)
	if err != nil {
		return "", err
	}
	obScore := &ObjectObjectName{
		ObjectID: object.ObjectId,
		ObjectName: object.ObjectName,
	}
	err = obScore.AddCToP()
	if err != nil {
		return "", err
	}
	obUser := &ObjectUser{
		ObjectID: object.ObjectId,
		UserID:   object.UserID,
	}
	err = obUser.AddCToP()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%012d", object.ObjectId), nil
}

func AddMultipleObject(list []Object) error {
	var err error
	listIt := []*generic.TItem{}
	listOn := []*generic.TItem{}
	listOu := []*generic.TItem{}
	for _, o := range list {
		b, err := json.Marshal(&o)
		if err != nil {
			return err
		}
		bs := make([]byte, 4)
		binary.BigEndian.PutUint32(bs, uint32(o.ObjectId))
		it := generic.TItem{
			Key:   bs,
			Value: b,
		}
		listIt = append(listIt, &it)
		//add to object name
		on := ObjectObjectName{
			ObjectID:   o.ObjectId,
			ObjectName: o.ObjectName,
		}
		oit := generic.TItem{
			Key:   on.Key(),
			Value: bs,
		}
		listOn = append(listOn, &oit)
		//add to user
		ou := ObjectUser{
			ObjectID: o.ObjectId,
			UserID:   o.UserID,
		}
		nit := generic.TItem{
			Key:   ou.Key(),
			Value: bs,
		}
		listOu = append(listOu, &nit)
	}
	_, err = BigsetIf.BsMultiPut2(OBJECT, listIt)
	_, err = BigsetIf.BsMultiPut2(OBJECT_OBJECT_NAME, listOn)
	_, err = BigsetIf.BsMultiPut2(OBJECT_USER, listOu)
	return err
}

func GetOne(ObjectId string) (object *Object, err error) {
	i, err := BigsetIf.BsGetItem2(OBJECT, generic.TItemKey(ObjectId))
	if err != nil {
		return nil, err
	}
	var ob Object
	err = json.Unmarshal(i.GetValue(), &ob)
	return &ob, err
}

func GetPaginateIn(UserID string, counter int32) ([]*Object, error) {
	listit, err := BigsetIf.BsGetSliceFromItem2(USER, generic.TItemKey(UserID+"-"), counter)
	if err != nil {
		return nil, err
	}
	listOb := []*Object{}
	for _, i := range listit {
		var ob Object
		err = json.Unmarshal(i.GetValue(), &ob)
		if err != nil {
			return nil, err
		}
		listOb = append(listOb, &ob)
	}
	return listOb, nil
}
