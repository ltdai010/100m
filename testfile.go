package main

import (
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"log"
	"push-data/models"
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
	////listit, _, _ := models.BigsetIf.BsRangeQueryByPage2(models.OBJECT, bs1, bs2, 0, 10)
	//listit, _ := models.BigsetIf.BsGetSliceFromItem2(models.OBJECT_USER, bs2, 10)
	//for _, it := range listit {
	//	log.Println(binary.LittleEndian.Uint32(it.GetKey()))
	//}
	//fmt.Println(time.Now().Unix() - s)
	//RemoveAll(models.OBJECT_OBJECT_NAME)
	i, err := models.BigsetIf.GetTotalCount2(models.OBJECT)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(i)
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
