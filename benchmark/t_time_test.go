package main

import (
	"fmt"
	"testing"
	"time"
)

/*
时间转换
*/

func Test_Unix2Str(t *testing.T) {
	timestamp := 1653613076
	str := unixToStr(int64(timestamp),  "2006/01/02 15:04:05")
	fmt.Println("time:", str)
	h := str[11:13]
	fmt.Println("H:", h)
	
	
}


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


func Str2Unix(){
	prev := time.Now().AddDate(0, -2, 0) 	
	resTime := prev.Format("2006-01-02 00:00:00") //获取的时间的格式
	fmt.Println("res:", resTime)
	stamp, _ := time.ParseInLocation("2006-01-02 15:04:05", resTime, time.Local)
	fmt.Println("stamp:", stamp.Unix())
}

func unixToStr(timeUnix int64, layout string) string {
	timeStr := time.Unix(timeUnix, 0).Format(layout)
	return timeStr
}