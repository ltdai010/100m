package models

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

type ObjectObjectName struct {
	ObjectID   int32  `json:"object_id"`
	ObjectName string `json:"object_name"`
}

const OBJECT_OBJECT_NAME = "object_name"


func (this *ObjectObjectName) ChildID() int32 {
	return this.ObjectID
}

func (this *ObjectObjectName) Key() []byte {
	cp := make([]byte, 4)
	binary.BigEndian.PutUint32(cp, uint32(this.ChildID()))
	key := append([]byte(this.ObjectName+"-"), cp...)
	return key
}

func (g *ObjectObjectName) GetBsKey() generic.TStringKey {
	return generic.TStringKey(fmt.Sprintf("%s", OBJECT_OBJECT_NAME))
}

func (g *ObjectObjectName) GetChildBsKey() generic.TStringKey {
	return generic.TStringKey(fmt.Sprintf("%s", OBJECT))
}

func (this *ObjectObjectName) AddCToP() error {
	cp := make([]byte, 4)
	binary.BigEndian.PutUint32(cp, uint32(this.ChildID()))
	p, err := BigsetIf.BsGetItem2(this.GetChildBsKey(), cp)
	if err != nil {
		return err
	}
	if p == nil {
		return errors.New("not exist ")
	}
	key := append([]byte(this.ObjectName+"-"), cp...)

	it := generic.TItem{Key: key, Value: cp}
	_, err = BigsetIf.BsPutItem2(this.GetBsKey(), &it)
	if err != nil {
		return err
	}
	return nil
}

func (this *ObjectObjectName) CInP() (bool, error) {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, uint32(this.ObjectID))
	key := append([]byte(this.ObjectName+"-"), bs...)
	it, err := BigsetIf.BsGetItem2(this.GetBsKey(), key)
	if it == nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (this *ObjectObjectName) UnMarshalArrayTItem(objects []*generic.TItem) ([]User, error) {
	objs := make([]User, 0)
	for _, object := range objects {
		obj := User{}
		err := json.Unmarshal(object.GetValue(), &obj)

		if err != nil {
			return make([]User, 0), err
		}

		objs = append(objs, obj)
	}

	return objs, nil
}

func (this *ObjectObjectName) GetPaginate(count int32) ([]*Object, error) {
	setItems, err := BigsetIf.BsGetSliceFromItem2(this.GetBsKey(), generic.TItemKey(this.GetChildBsKey()+"-"), count)
	if err != nil {
		return nil, err
	}
	mapOb := []*Object{}
	for _, i := range setItems {
		it, err := BigsetIf.BsGetItem2(this.GetChildBsKey(), i.GetValue())
		if err != nil {
			return nil, err
		}
		if it == nil {
			continue
		}
		var o Object
		err = json.Unmarshal(it.GetValue(), &o)
		if err != nil {
			return nil, err
		}
		mapOb = append(mapOb, &o)
	}
	return mapOb, err
}