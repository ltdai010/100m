package models

import (
	"encoding/binary"
	"encoding/json"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

type User struct {
	UserID   int32  `json:"user_id"`
	Username string `json:"username"`
}

const USER_COUNTER = "user_counter"
const USER = "user"

func AddUser(u User) (int32, error) {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, uint32(u.UserID))
	b, err := json.Marshal(&u)
	it := generic.TItem{
		Key:   bs,
		Value: b,
	}
	_, err = BigsetIf.BsPutItem2(USER, &it)
	usernameUser := &UserUsername{
		UserID:   u.UserID,
		Username: u.Username,
	}
	err = usernameUser.AddCToP()
	return u.UserID, err
}

func AddMultipleUser(list []User) error {
	var err error
	var listIt []*generic.TItem
	var listUn []*generic.TItem
	for _, u := range list {
		b, err := json.Marshal(&u)
		if err != nil {
			return err
		}
		bs := make([]byte, 4)
		binary.BigEndian.PutUint32(bs, uint32(u.UserID))
		it := generic.TItem{
			Key:   bs,
			Value: b,
		}
		listIt = append(listIt, &it)
		un := &UserUsername{
			UserID:   u.UserID,
			Username: u.Username,
		}
		unit := generic.TItem{
			Key:   un.Key(),
			Value: bs,
		}
		listUn = append(listUn, &unit)
	}
	_, err = BigsetIf.BsMultiPut2(USER, listIt)
	if err != nil {
		return err
	}
	_, err = BigsetIf.BsMultiPut2(USER_USERNAME, listUn)
	return err
}

func GetUser(uid string) (u *User, err error) {
	i, err := BigsetIf.BsGetItem2(USER, generic.TItemKey(uid))
	if err != nil {
		return nil, err
	}
	var user User
	err = json.Unmarshal(i.GetValue(), &user)
	return &user, err
}

func GetPaginateUsers(pos int32, count int32) ([]*User, error) {
	listIt, err := BigsetIf.BsGetSlice2(USER, pos, count)
	if err != nil {
		return nil, err
	}
	listU := []*User{}
	for _, i := range listIt {
		u := User{}
		err = json.Unmarshal(i.GetValue(), &u)
		if err != nil {
			return nil, err
		}
		listU = append(listU, &u)
	}
	return listU, nil
}
