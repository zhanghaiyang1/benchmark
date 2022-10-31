package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)


func Test_RandNum(t *testing.T) {
	fmt.Println("start--")
	num := 4
	fmt.Println("rand:", r.Intn(num))
}
var r *rand.Rand

func init() {
	fmt.Println("init---")
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

func RandNumber(number int) int {
	return r.Intn(number)
}
