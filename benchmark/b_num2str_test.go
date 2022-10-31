package main

import (
	"fmt"
	"strconv"
	"testing"
)

//Sprintf
func BenchmarkSprintfNum2Str(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", num)
	}
}

//Format
func BenchmarkFormatNum2Str(b *testing.B) {
	num := int64(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(num, 10)
	}
}

//Itoa
func BenchmarkItoaNum2Str(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(num)
	}
}
/*  

1，运行 
go test -bench='Num2Str$'	//只运行以Num2Str结尾的benchmark用例

-benchmem表示查看内存分配情况

2，结果
goos: darwin
goarch: arm64
pkg: benchmark
BenchmarkSprintfNum2Str-8       28696822                41.83 ns/op            2 B/op          1 allocs/op
BenchmarkFormatNum2Str-8        584458755                2.036 ns/op           0 B/op          0 allocs/op
BenchmarkItoaNum2Str-8          586865101                2.037 ns/op           0 B/op          0 allocs/op
PASS
ok      benchmark       5.553s

3，说明
BenchmarkSprintfNum2Str-8	BenchmarkSprintfNum2Str 是测试的函数名 -8 表示GOMAXPROCS（线程数）的值为12
28696822	表示一共执行了28696822次，即b.N的值
41.83 ns/op	表示平均每次操作花费了41.83纳秒
2 B/op	表示每次操作申请了2Byte的内存申请
1 allocs/op 表示每次操作申请了1次内存

4，对比
可以发现，BenchmarkSprintf方法耗时最长，BenchmarkFormat最快，BenchmarkItoa也很快。差别在于fmt.Sprintf()执行过程中进行了一次内存分配1 allocs/op。
*/
