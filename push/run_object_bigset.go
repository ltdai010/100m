package main

//import (
//	"encoding/json"
//	"fmt"
//	"log"
//	"math/rand"
//	"push/models"
//	"sync"
//	"time"
//)
//
//var (
//	chanObject chan []models.Object
//	wg         sync.WaitGroup
//	in         int
//)
//
//const (
//	loop = 5
//	step = 1000
//	channel = 1000
//	workers = 4
//)
//
////worker
//func pushMulti(list <- chan []models.Object, wg *sync.WaitGroup) {
//	for i := range list {
//		err := models.AddMultipleObject(i)
//		if err != nil {
//			log.Fatal(err)
//		}
//		in++
//		log.Println("pushed", in*step)
//		wg.Done()
//	}
//}
//
//func main() {
//	in = 0
//	chanObject = make(chan []models.Object, channel)
//	//init database
//	models.InitBigset()
//	//run worker
//	for i := 0; i < workers; i++ {
//		go pushMulti(chanObject, &wg)
//	}
//	//check start time
//	startTime := time.Now().Unix()
//	//start loop
//	for k := 0; k < loop; k++ {
//		//get recentID
//		count, err := models.Kvcountersv.GetStepValue(models.OBJECT_COUNTER, step*channel)
//		wg.Add(channel)
//		users := []*models.User{}
//		//get user to make object of
//		listit, err := models.BigsetIf.BsGetSlice2(models.USER, int32(count/step), channel)
//		log.Println("get user from", count/step)
//		if err != nil {
//			log.Fatal(err)
//			return
//		}
//		//make user list
//		for _, it := range listit {
//			u := &models.User{}
//			json.Unmarshal(it.GetValue(), u)
//			users = append(users, u)
//		}
//		//start channel
//		for i := 0; i < channel; i++ {
//			list := []models.Object{}
//			n := fmt.Sprint(i)
//			fmt.Println("start from", count)
//			if err != nil {
//				log.Fatal(err)
//			}
//			//make object
//			for j := 0; j < step; j++ {
//				r := rand.Intn(99)
//				u := models.Object{
//					ObjectId:   int32(count),
//					ObjectName: n,
//					Score:      int32(r),
//					UserID:     users[i].UserID,
//				}
//				count++
//				list = append(list, u)
//			}
//			//push to channel
//			chanObject <- list
//			log.Println("creating", ((i + 1)*step))
//		}
//		//wait for all channel are empty
//		wg.Wait()
//	}
//	//calculate executed time
//	fmt.Println(time.Now().Unix() - startTime)
//}
