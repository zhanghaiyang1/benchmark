package main

import (
	"fmt"
	"time"
)



func main(){

	str := "2022-08-05"
	start := str + " 00:00:00"
	
	fmt.Println(start)

	stamp, _ := time.ParseInLocation("2006-01-02 15:04:05", start, time.Local)
	dateTimestamp := stamp.Unix()
	fmt.Println("dateTimestamp:", dateTimestamp)

	getTime := stamp.AddDate(1, 0, -1)  //年，月，日   获取1
	resTime := getTime.Format("2006-01-02") //获取的时间的格式
	fmt.Println("restime:", resTime)

}
