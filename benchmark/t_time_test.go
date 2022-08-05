package main

import (
	"fmt"
	"testing"
	"time"
)


func Test_Time(t *testing.T) {
	timestamp := 1653614776
	str := unixToStr(int64(timestamp),  "2006/01/02 15:04:05")
	fmt.Println("time:", str)
}
func unixToStr(timeUnix int64, layout string) string {
	timeStr := time.Unix(timeUnix, 0).Format(layout)
	return timeStr
}
// func strToUnix(timeUnix int64, layout string) string {
// 	timeStr := timeUnix.Format("2006-01-02 15:04:05")
// 	return timeStr
// }
func Test_Single(t *testing.T){
	res:=time.Now().Format("200601")
	fmt.Println("res:", res)
}
func Test_Prev(t *testing.T){
	current :=time.Now().Format("200601") 
	prev := time.Now().AddDate(0, -11, 0)             //年，月，日   获取一个月前的时间
	resTime := prev.Format("200601") //获取的时间的格式
	fmt.Println("current:", current)
	fmt.Println("resTime:", resTime)
}
