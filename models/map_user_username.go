package models

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

type UserUsername struct {
	UserID   int32  `json:"user_id"`
	Username string `json:"username"`
}

const USER_USERNAME = "user_username"

func (g *UserUsername) ParentID() string {
	return g.Username
}

func (this *UserUsername) ChildID() int32 {
	return this.UserID
}

func (this *UserUsername) Key() []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, uint32(this.UserID))
	key := append([]byte(this.ParentID()+"-"), bs...)
	return key
}

func (g *UserUsername) GetBsKey() generic.TStringKey {
	return generic.TStringKey(fmt.Sprintf("%s", USER_USERNAME))
}

func (g *UserUsername) GetChildBsKey() generic.TStringKey {
	return generic.TStringKey(fmt.Sprintf("%s", USER))
}

func (this *UserUsername) AddCToP() error {
	p, err := BigsetIf.BsGetItem2(this.GetChildBsKey(), []byte{byte(this.ChildID())})
	if err != nil {
		return err
	}
	if p == nil {
		return errors.New("not exist")
	}
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, uint32(this.UserID))
	key := append([]byte(this.ParentID()+"-"), bs...)

	it := generic.TItem{Key: key, Value: bs}
	_, err = BigsetIf.BsPutItem2(this.GetBsKey(), &it)
	if err != nil {
		return err
	}
	return nil
}

func (this *UserUsername) CInP() (bool, error) {
	key := append([]byte(this.ParentID()+"-"), byte(this.UserID))
	it, err := BigsetIf.BsGetItem2(this.GetBsKey(), key)
	if it == nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (this *UserUsername) UnMarshalArrayTItem(objects []*generic.TItem) ([]User, error) {
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

func (this *UserUsername) GetPaginate(count int32) ([]*User, error) {
	setItems, err := BigsetIf.BsGetSliceFromItem2(this.GetBsKey(), generic.TItemKey(this.GetChildBsKey()+"-"), count)
	if err != nil {
		return nil, err
	}
	mapOb := []*User{}
	for _, i := range setItems {
		it, err := BigsetIf.BsGetItem2(this.GetChildBsKey(), i.GetValue())
		if err != nil {
			return nil, err
		}
		if it == nil {
			continue
		}
		var o User
		err = json.Unmarshal(it.GetValue(), &o)
		if err != nil {
			return nil, err
		}
		mapOb = append(mapOb, &o)
	}
	return mapOb, err
}
