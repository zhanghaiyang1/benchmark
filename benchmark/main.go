package main

import (
	"fmt"
	"time"
)

func main() {
	aType := int32(0)
	go func() {
		aType = 1
		fmt.Printf("1执行时间：%d", time.Now().Unix())
	}()
	time.Sleep(time.Second * 5)
	go func() {
		fmt.Printf("2执行时间：%d", time.Now().Unix())
	}()
	fmt.Println("atype:", aType)
}
