package main

import (
	"fmt"
	pract "gostudy/pract"
	"gostudy/stu"
	t "gostudy/tools"
	"os"
)

// main方法
func main() {
	fmt.Println("hey world!!!")

	fmt.Println("-----------------BMI  MY TEST CODE")
	args := os.Args
	var W float64
	var H float64
	fmt.Sscanf(args[1], "%f", &W)
	fmt.Sscanf(args[2], "%f", &H)
	t.Printfln("身高，%.2f", H)
	t.Printfln("体重，%.2f", W)
	pract.ClacBMI(W, H)

	fmt.Println("-----------------int64")
	stu.TestStruct()

	fmt.Println("-----------------func")
	stu.TestFunc()

}
