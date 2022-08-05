package main

import (
	"fmt"
	"testing"
	"net/url"
)


func Test_Url(t *testing.T) {
	price := 1010000
	rate := 10
	res := price * rate / 100
	fmt.Println("res:", res)

	resUrl := "http://hhuibao.oss-cn-qingdao.aliyuncs.com"
	link, _ := url.Parse(resUrl)
	s := make([]string, 1)
	fmt.Println("link:", link.Host)
	fmt.Println("link.path:", link.Path)
	s = append(s,  link.Path[1:])
	fmt.Println("s:", s)
}