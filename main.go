package main

import (
	"fmt"
	"os"
	t "gostudy/tools"
	pract "gostudy/pract"
)

// main方法
func main() {
	fmt.Println("hey world!!!")
	args := os.Args
	var W float64
	var H float64
	fmt.Sscanf(args[1], "%f", &W)
	fmt.Sscanf(args[2], "%f", &H)
	t.Printfln("身高，%.2f", H)
	t.Printfln("体重，%.2f", W)
	pract.ClacBMI(W, H)

	
}
