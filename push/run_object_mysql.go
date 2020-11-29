package main2

//var (
//	in       int
//)
//
//const (
//	loop = 90
//	step = 1000
//	channel = 1000
//	workers = 4
//)
//
//func pushMulti(list []models.Object) {
//	s:= ""
//	n := ""
//	q, err := models.Mysql.Query("SELECT MAX(ObjectID) as id FROM objects")
//	id := 0
//	q.Next()
//	q.Scan(&id)
//	q.Close()
//	for j := 0; j < len(list); j++ {
//		if j == len(list) - 1 {
//			id++
//			s += "("+ list[j].ObjectName +"," + fmt.Sprint(list[j].Score) + ","+ fmt.Sprint(list[j].UserID) +")"
//			n += "("+ fmt.Sprint(list[j].UserID) +"," + fmt.Sprint(id) +")"
//		} else {
//			id++
//			s += "("+ list[j].ObjectName +"," + fmt.Sprint(list[j].Score) + ","+ fmt.Sprint(list[j].UserID) +"),"
//			n += "("+ fmt.Sprint(list[j].UserID) +"," + fmt.Sprint(id) +"),"
//		}
//	}
//	i, err := models.Mysql.Query("INSERT INTO objects(ObjectName, Score, UserID) VALUES " + s)
//	if err != nil {
//		log.Fatal(err)
//	}
//	j, err := models.Mysql.Query("INSERT INTO `objects-user`(UserID, ObjectID) VALUES " + n)
//	if err != nil {
//		log.Fatal(err)
//	}
//	i.Close()
//	j.Close()
//}
//
//func main() {
//	in = 0
//	startTime := time.Now().Unix()
//	models.InitBigset()
//	for k := 0; k < loop; k++ {
//		//wg.Add(channel)
//		listUser := loadListUser(k * channel, channel)
//		for i := 0; i < channel; i++ {
//			n := fmt.Sprint(k*channel + i)
//			list := []models.Object{}
//			for j := 0; j < step; j++ {
//				s := rand.Intn(99)
//				u := models.Object{
//					ObjectName: n,
//					Score: int32(s),
//					UserID: listUser[i].UserID,
//				}
//				list = append(list, u)
//			}
//			pushMulti(list)
//			in++
//			log.Println("pushed", in*step)
//		}
//		//wg.Wait()
//	}
//	fmt.Println(time.Now().Unix() - startTime)
//}
//
//func loadListUser(from int, count int) []models.User {
//	listIt, _ := models.BigsetIf.BsGetSlice2(models.USER, int32(from), int32(count))
//	list := []models.User{}
//	for _, it := range listIt {
//		u := models.User{}
//		json.Unmarshal(it.GetValue(), &u)
//		list = append(list, u)
//	}
//	return list
//}