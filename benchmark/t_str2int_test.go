package main

import (
	"strconv"
	"testing"
	"fmt"
)

func Test_str2int(t *testing.T){
	
	h := "09"
	hInt, _ := strconv.Atoi(h)
	fmt.Println(hInt)

	i, _ := strconv.ParseInt("12345", 10, 64)
	fmt.Println(i)
}