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