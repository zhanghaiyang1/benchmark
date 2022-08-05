package main

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func Test_No(t *testing.T) {
	no := 1123
	res := no / 100
	fmt.Println("res:", res)

	res2 := decimal.NewFromInt(int64(no)).Div(decimal.NewFromInt(100)).IntPart()
	fmt.Println("res2:", res2)
}