// 声明包
package main

// 引入fmt
import "fmt"

// 只能出现在函数体中
//t1, t2 := 1, 2

// 因式分解 通常定义全局变量
var (
	t3 int
	t4 bool
)

func main() {
	// 定义单个变量
	var test string = "test"

	// 定义多个变量
	var test2, test3 int = 3, 4

	// 打印
	fmt.Println(test, test2, test3)

	fmt.Println("---------------------")

	// 如果定义了变量没给初始值 默认是0 false
	var test4 int
	fmt.Println(test4)
	// 不给类型 根据数据类型
	var test5 = "tttttt"
	fmt.Println(test5)
	// 默认false
	var test6 bool
	fmt.Println(test6)

	fmt.Println("下面都是nil")
	var a *int
	var b []int
	var c map[string] int
	var d chan int
	var e func(string) int
	var f error
	fmt.Println(a, b, c, d, e, f)

	// 多变量声明
	fmt.Println("下面多变量声明")
	// 声明类型
	var a1, a2, a3 int
	a1, a2, a3 = 1, 2, 3
	fmt.Println("a", a1, a2, a3)
	// 推测类型
	var b1, b2, b3 = 1, "a", true
	fmt.Println("b", b1, b2, b3)
	// 语法糖
	c1, c2, c3 := "c", 3, false
	fmt.Println("c", c1, c2, c3)
	// 因式分解
	t3, t4 = 111, false
	fmt.Println("t", t3, t4)

	// 不适用也会报错
	//var aaaa int
}
