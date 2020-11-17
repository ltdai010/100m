package main

import (
	"fmt"
	"log"
	"push-data/models"
	"sync"
	"time"
)

var (
	chanObject chan []models.User
	wg       sync.WaitGroup
	in       int
)

const (
	loop = 1
	step = 10000
	channel = 100
	workers = 4
)

func pushMulti(list <- chan []models.User, wg *sync.WaitGroup) {
	for i := range list {
		err := models.AddMultipleUser(i)
		if err != nil {
			log.Fatal(err)
		}
		in++
		log.Println("pushed", in*step)
		wg.Done()
	}
}

func main() {
	in = 0
	startTime := time.Now().Unix()
	chanObject = make(chan []models.User, channel+10)
	models.InitBigset()
	for i := 0; i < workers; i++ {
		go pushMulti(chanObject, &wg)
	}
	for k := 0; k < loop; k++ {
		count, err := models.Kvcountersv.GetStepValue(models.USER_COUNTER, step*channel)
		wg.Add(channel)
		for i := 0; i < channel; i++ {
			list := []models.User{}
			n := fmt.Sprint(i)
			fmt.Println("start from", count)
			if err != nil {
				log.Fatal(err)
			}
			for j := 0; j < step; j++ {
				u := models.User{
					UserID: int32(count),
					Username: n,
				}
				count++
				list = append(list, u)
			}
			chanObject <- list
			log.Println("creating", ((i + 1)*step))
		}
		wg.Wait()
	}
	fmt.Println(time.Now().Unix() - startTime)
}
