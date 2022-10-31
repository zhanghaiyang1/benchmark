package main

import (
	"fmt"
	"testing"
	"net/url"
)
/* 
url截取
*/

func Test_Url(t *testing.T) {

	resUrl := "http://hhuibao.oss-cn-qingdao.aliyuncs.com"
	link, _ := url.Parse(resUrl)
	s := make([]string, 1)
	fmt.Println("link:", link.Host)
	fmt.Println("link.path:", link.Path)
	s = append(s,  link.Path[1:])
	fmt.Println("s:", s)
}