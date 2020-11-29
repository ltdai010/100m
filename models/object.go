package models

import (
	"encoding/binary"
	"encoding/json"
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

func AddOne(object Object) (int32, error) {
	//convert int32 to byte
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, uint32(object.ObjectId))
	b, err := json.Marshal(&object)
	it := generic.TItem{
		Key:   bs,
		Value: b,
	}
	//put to database
	_, err = BigsetIf.BsPutItem2(OBJECT, &it)
	if err != nil {
		return 0, err
	}
	obScore := &ObjectObjectName{
		ObjectID: object.ObjectId,
		ObjectName: object.ObjectName,
	}
	//index object name
	err = obScore.AddCToP()
	if err != nil {
		return 0, err
	}
	obUser := &ObjectUser{
		ObjectID: object.ObjectId,
		UserID:   object.UserID,
	}
	//index objectID and userID
	err = obUser.AddCToP()
	if err != nil {
		return 0, err
	}
	return object.ObjectId, nil
}

func AddMultipleObject(list []Object) error {
	var err error
	//create list item
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
	//put list item
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

func GetPaginateObjectWithObjectName(ObjectName string, page, counter int32) ([]*Object, error) {
	//get slice form key obname-
	listit, err := BigsetIf.BsGetSliceFromItem2(OBJECT_OBJECT_NAME, generic.TItemKey(ObjectName + "-"), counter)
	if err != nil {
		return nil, err
	}
	if len(listit) == 0 {
		return nil, nil
	}
	key := listit[len(listit) - 1].GetKey()
	var i int32
	//roll page
	for i = 0; i < page; i++ {
		listit, err = BigsetIf.BsGetSliceFromItem2(OBJECT_OBJECT_NAME, key, counter)
		if err != nil {
			return nil, err
		}
		key = listit[len(listit) - 1].GetKey()
	}
	listO := []*Object{}
	//get object value
	for _, i := range listit {
		o := &Object{}
		it, err := BigsetIf.BsGetItem2(OBJECT, i.GetValue())
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(it.GetValue(), o)
		if err != nil {
			return nil, err
		}
		listO = append(listO, o)
	}
	return listO, nil
}