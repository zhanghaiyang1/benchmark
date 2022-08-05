package main

import (
	"fmt"
	"testing"
)

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++{
		fib(30)
	}
}
func fib(n int) int{
	if n == 0 || n == 1 {
		return n
	}
	return fib(n - 2) + fib(n -1)
}
/* 
1，go test -bench . 运行全部
2，go test -bench='Fib$' . // 只运行以Fib结尾的benchmark用例

*/

func TestRepeat(t *testing.T) {
    fmt.Println("aaaa")
}