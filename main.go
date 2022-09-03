package main

import (
	"fmt"
	"gostudy/stu"
	"os"
)

// init main之前执行
func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./example2 <url>")
		os.Exit(-1)
	}
}

// main方法
func main() {
	fmt.Println("hey world!!!")

	// fmt.Println("-----------------BMI  MY TEST CODE")
	// args := os.Args
	// var W float64
	// var H float64
	// fmt.Sscanf(args[1], "%f", &W)
	// fmt.Sscanf(args[2], "%f", &H)
	// t.Printfln("身高，%.2f", H)
	// t.Printfln("体重，%.2f", W)
	// pract.ClacBMI(W, H)

	// fmt.Println("-----------------int64")
	// stu.TestStruct()

	// fmt.Println("-----------------func")
	// stu.TestFunc()

	// fmt.Println("-----------------curl")
	// 测试url curl 发送请求
	// r, err := http.Get(os.Args[1])
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// 从body取值
	// io.Copy(os.Stdout, r.Body)
	// if err := r.Body.Close(); err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("-----------------copy")
	// stu.TestCopy()

	// fmt.Println("-----------------测试公开性")
	// 未公开不能直接使用
	// val := stu.alterCounter(1)
	val1 := stu.New(1)
	fmt.Printf("val1 is %d\n", val1)
}
