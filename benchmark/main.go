package main

import (
	"fmt"
	"time"
)

var WeekDayMap = map[string]int8{
	"Monday":    1,
	"Tuesday":   2,
	"Wednesday": 3,
	"Thursday":  4,
	"Friday":    5,
	"Saturday":  6,
	"Sunday":    7,
}
func main() {
	// current :=time.Now().Format("2006-01-02") 
	prev := time.Now().AddDate(0, -2, 0) 	
	resTime := prev.Format("2006-01-02 00:00:00") //获取的时间的格式
	fmt.Println("res:", resTime)
	stamp, _ := time.ParseInLocation("2006-01-02 15:04:05", resTime, time.Local)
	fmt.Println("stamp:", stamp.Unix())
	week := stamp.Weekday().String()
	fmt.Println("weeK:", week)

	stampUnix := stamp.AddDate(0, 2, -1) 
	fmt.Println("stampUnix:", stampUnix)

	created := int64(1664553600)

	t:=time.Unix(created,0) 
	fmt.Println("t.unix:", t.Unix())
	
	resJia := t.Add(time.Hour * 2)
	fmt.Println("resJia:", resJia.Unix())
}
