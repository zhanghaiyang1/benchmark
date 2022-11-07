package main

import (
	"fmt"
	"testing"
)


func Test_Replace(t *testing.T) {
	s := []int{9, 8, 7}
	p := &s
	r := *p
	r[0] = 11
	fmt.Println(s[0])
}
func Test_Split(t *testing.T) {
	str := "13521811024"
	fmt.Println(str[5:])
	fmt.Println(str[:8])
}