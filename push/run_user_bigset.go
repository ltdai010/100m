package main2
//
//import (
//	"fmt"
//	"log"
//	"push-data/models"
//	"sync"
//	"time"
//)
//
//var (
//	chanObject chan []models.User //channel store data
//	wg       sync.WaitGroup //create wait group
//	in       int //count
//)
//
//const (
//	loop = 1 //so lan lap
//	step = 10000 //do dai mang push vao 1 lan
//	channel = 100 //so cong viec
//	workers = 4 //so cong nhan
//)
//
////worker
//func pushMulti(list <- chan []models.User, wg *sync.WaitGroup) {
//	for i := range list {
//		err := models.AddMultipleUser(i)
//		if err != nil {
//			log.Fatal(err)
//		}
//		in++
//		log.Println("pushed", in*step)
//		wg.Done() //done 1 channel
//	}
//}
//
//func main() {
//	in = 0 //counter
//	chanObject = make(chan []models.User, channel+10)
//	//initial database
//	models.InitBigset()
//	//run worker
//	for i := 0; i < workers; i++ {
//		go pushMulti(chanObject, &wg)
//	}
//	//check start time
//	startTime := time.Now().Unix()
//	//start loop
//	for k := 0; k < loop; k++ {
//		count, err := models.Kvcountersv.GetStepValue(models.USER_COUNTER, step*channel)
//		//add channels
//		wg.Add(channel)
//		for i := 0; i < channel; i++ {
//			list := []models.User{}
//			name := fmt.Sprint(i)
//			fmt.Println("start from", count)
//			if err != nil {
//				log.Fatal(err)
//			}
//			//create an array of users
//			for j := 0; j < step; j++ {
//				u := models.User{
//					UserID: int32(count),
//					Username: name,
//				}
//				count++
//				list = append(list, u)
//			}
//			//push to channel
//			chanObject <- list
//			log.Println("creating", ((i + 1)*step))
//		}
//		//wait for all channel to be empty
//		wg.Wait()
//	}
//	//calculate executed time
//	fmt.Println(time.Now().Unix() - startTime)
//}
