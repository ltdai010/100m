package main

//var (
//	chanObject chan []models.Object
//	wg         sync.WaitGroup
//	in         int
//)
//
//const (
//	loop = 20
//	step = 1000
//	channel = 1000
//	workers = 4
//)
//
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
//	startTime := time.Now().Unix()
//	chanObject = make(chan []models.Object, channel)
//	models.InitBigset()
//	for i := 0; i < workers; i++ {
//		go pushMulti(chanObject, &wg)
//	}
//	for k := 0; k < loop; k++ {
//		count, err := models.Kvcountersv.GetStepValue(models.OBJECT_COUNTER, step*channel)
//		wg.Add(channel)
//		users := []*models.User{}
//		listit, err := models.BigsetIf.BsGetSlice2(models.USER, int32(k*channel), channel)
//		if err != nil {
//			log.Fatal(err)
//			return
//		}
//		for _, it := range listit {
//			u := &models.User{}
//			json.Unmarshal(it.GetValue(), u)
//			users = append(users, u)
//		}
//		for i := 0; i < channel; i++ {
//			list := []models.Object{}
//			n := fmt.Sprint(i)
//			fmt.Println("start from", count)
//			if err != nil {
//				log.Fatal(err)
//			}
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
//			chanObject <- list
//			log.Println("creating", ((i + 1)*step))
//		}
//		wg.Wait()
//	}
//	fmt.Println(time.Now().Unix() - startTime)
//}
