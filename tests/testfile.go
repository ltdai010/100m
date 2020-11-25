package test

import (
	"encoding/binary"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"log"
	"push-data/models"
	"time"
)

func main() {
	models.InitBigset()
	//i, _ := models.BigsetIf.GetTotalCount2(models.OBJECT_USER)
	//log.Println(i)
	//bs1 := make([]byte, 4)
	//binary.BigEndian.PutUint32(bs1, 1000000)
	//bs2 := make([]byte, 4)
	//binary.BigEndian.PutUint32(bs2, 1000000)
	//s := time.Now().Unix()
	//fmt.Println(time.Now().Unix() - s)
	//RemoveAll(models.OBJECT_OBJECT_NAME)
	startTime := time.Now().UnixNano()
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, 1090000)
	key := append(bs, []byte("-")...)
	list, err := models.BigsetIf.BsGetSliceFromItem2(models.OBJECT, key, 10)
	log.Println(len(list))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("executed time", time.Now().UnixNano() - startTime)
}

func RemoveAll(key string) {
	leng, _ := models.BigsetIf.GetTotalCount2(generic.TStringKey(key))
	for i := 0; i < 100; i++ {
		listit, _ := models.BigsetIf.BsGetSlice2(generic.TStringKey(key), 0, int32(leng / 100))
		for _, it := range listit {
			models.BigsetIf.BsRemoveItem2(generic.TStringKey(key), it.GetKey())
		}
		log.Println(int64(i) * leng / 100)
	}
}
