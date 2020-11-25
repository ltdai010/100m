package models

import (
	"database/sql"
	"github.com/OpenStars/EtcdBackendService/KVCounterService"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

var (
	bigsetOnce sync.Once
	BigsetIf   StringBigsetService.StringBigsetServiceIf
	Kvcountersv KVCounterService.KVCounterServiceIf
	Mysql		*sql.DB
)

func InitBigset() {
	bigsetOnce.Do(func() {
		BigsetIf = StringBigsetService.NewStringBigsetServiceModel("/data/",
			[]string{"127.0.0.1:2379"},
			GoEndpointBackendManager.EndPoint{
				Host:      "localhost",
				Port:      "18990",
				ServiceID: "/data/",
			})
		Kvcountersv = KVCounterService.NewKVCounterServiceModel("/data-counter/",
			[]string{},
			GoEndpointBackendManager.EndPoint{
				Host:      "127.0.0.1",
				Port:      "20001",
				ServiceID: "/data-counter/",
			})
		var err error
		Mysql, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
		if err != nil {
			log.Fatal(err)
		}
	})
}

