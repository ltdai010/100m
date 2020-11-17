package models

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

type ObjectUser struct {
	ObjectID int32 `json:"object_id"`
	UserID   int32 `json:"user_id"`
}

const OBJECT_USER = "object_user"

func (g *ObjectUser) ParentID() int32 {
	return g.UserID
}

func (this *ObjectUser) ChildID() int32 {
	return this.ObjectID
}

func (this *ObjectUser) Key() []byte {
	cb := make([]byte, 4)
	binary.BigEndian.PutUint32(cb, uint32(this.ChildID()))
	pb := make([]byte, 4)
	binary.BigEndian.PutUint32(pb, uint32(this.ParentID()))
	key := append(pb, []byte("-")...)
	key = append(key, cb...)
	return key
}

func (g *ObjectUser) GetBsKey() generic.TStringKey {
	return generic.TStringKey(fmt.Sprintf("%s", OBJECT_USER))
}

func (g *ObjectUser) GetChildBsKey() generic.TStringKey {
	return generic.TStringKey(fmt.Sprintf("%s", OBJECT))
}

func (g *ObjectUser) GetParentBsKey() generic.TStringKey {
	return generic.TStringKey(fmt.Sprintf("%s", USER))
}

func (this *ObjectUser) AddCToP() error {
	cb := make([]byte, 4)
	binary.BigEndian.PutUint32(cb, uint32(this.ChildID()))
	p, err := BigsetIf.BsGetItem2(this.GetChildBsKey(), cb)
	if err != nil {
		return err
	}
	if p == nil {
		return errors.New("not exist")
	}
	pb := make([]byte, 4)
	binary.BigEndian.PutUint32(pb, uint32(this.ParentID()))
	g, err := BigsetIf.BsGetItem2(this.GetParentBsKey(), pb)
	if err != nil {
		return err
	}
	if g == nil {
		return errors.New("not exist")
	}
	key := append(pb, []byte("-")...)
	key = append(key, cb...)
	it := generic.TItem{Key: key, Value: cb}
	_, err = BigsetIf.BsPutItem2(this.GetBsKey(), &it)
	if err != nil {
		return err
	}
	return nil
}

func (this *ObjectUser) UnMarshalArrayTItem(objects []*generic.TItem) ([]Object, error) {
	objs := make([]Object, 0)

	for _, object := range objects {
		obj := Object{}
		err := json.Unmarshal(object.GetValue(), &obj)

		if err != nil {
			return make([]Object, 0), err
		}

		objs = append(objs, obj)
	}

	return objs, nil
}

func (this *ObjectUser) GetPaginate(count int32) ([]*Object, error) {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, uint32(this.UserID))
	key := append(bs, []byte("-")...)
	setItems, err := BigsetIf.BsGetSliceFromItem2(this.GetBsKey(), key, count)
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
