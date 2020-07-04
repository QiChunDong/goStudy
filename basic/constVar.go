package main

import "fmt"

// 常量的使用
func main() {
	// 不能被修改
	// 只能是基本类型 可以显式指定类型 可以隐式推测类型
	const test int = 123
	// 会报错
	//test = 21424234
	fmt.Println("test", test)

	// iota const时默认0 每次+1
	const (
		a1 = iota
		a2
		a3
		a4
	) // 0123
	fmt.Println(a1, a2, a3, a4)

	const (
		b1 = iota
		b2
		b3 = "test" // 此时虽然赋值了 但是依然存在占位
		b4
		b5 = iota // 所以重新赋值为iota时 不是重新从0开始 也不是接着b2 而是从占位开始 此时时4
		b6
	) // 0 1 test test 4 5
	fmt.Println(b1, b2, b3, b4, b5, b6)

	const (
		c1 = 1<<iota // 这里是1<<0 1
		c2 = 3<<iota // 这里是3<<1 6
		c3 // 这里是3<<2 12
		c4 // 这里是3<<3 24
	)
	fmt.Println(c1, c2, c3, c4)
}
