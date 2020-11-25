package main

//var (
//	chanObject chan []models.User
//	wg       sync.WaitGroup
//	in       int
//)
//
//const (
//	loop = 10
//	step = 10000
//	channel = 100
//	workers = 4
//)
//
//func pushMulti(list []models.User) {
//	s:= ""
//	for j := 0; j < len(list); j++ {
//		if j == len(list) - 1 {
//			s += "("+ list[j].Username +")"
//		} else {
//			s += "(" + list[j].Username +"),"
//		}
//	}
//	i, err := models.Mysql.Query("INSERT INTO users(Username) VALUES " + s)
//	defer i.Close()
//	if err != nil {
//		log.Fatal(err)
//	}
//	in++
//	log.Println("pushed", in*step)
//}
//
//func main() {
//	in = 0
//	startTime := time.Now().Unix()
//	models.InitBigset()
//	for k := 0; k < loop; k++ {
//		//wg.Add(channel)
//		for i := 0; i < channel; i++ {
//			n := fmt.Sprint(k*channel + i)
//			list := []models.User{}
//			for j := 0; j < step; j++ {
//				u := models.User{
//					Username: n,
//				}
//				list = append(list, u)
//			}
//			pushMulti(list)
//			log.Println("creating", ((i + 1)*step))
//		}
//		//wg.Wait()
//	}
//	fmt.Println(time.Now().Unix() - startTime)
//}
