package stutest

import (
	"fmt"
	"strconv"
	"testing"
)

// 必须要以Benchmark开头
// b *testing.B作为唯一参数
// 文件也必须以_test.go结尾
func BenchmarkSpringf(b *testing.B) {
	number := 10

	// 
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}

}

// strconv.FormatInt
func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}

}

// strconv.Itoa
func BenchmarkItoa(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}

}

// 测试命令  go test -benchmem -run=^$ -bench .
// goos: windows
// goarch: amd64
// pkg: gostudy/stutest
// cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
// BenchmarkSpringf-8      16697068                68.57 ns/op            2 B/op          1 allocs/op
// BenchmarkFormat-8       472447268                2.543 ns/op           0 B/op          0 allocs/op
// BenchmarkItoa-8         446101960                2.679 ns/op           0 B/op          0 allocs/op
// PASS
// ok      gostudy/stutest 4.353s
// 每op消耗68.57纳秒

// 测试命令  go test -benchmem -run=^$ -bench . -benchmem
// 每op消耗2b