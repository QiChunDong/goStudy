package main

import "fmt"

// 控制语句
func main() {

	// switch
	var value int = 1
	var desValue string = "test"
	switch value {
		case 2: desValue = "test2"
		case 3: desValue = "test3"
		case 4: desValue = "test4"
		case 5: desValue = "test5"
	default:
		desValue = "default"
	}

	switch {
		case desValue == "test2": fmt.Println(desValue, 22222222222)
		case desValue == "test3": fmt.Println(desValue, 33333333333)
		case desValue == "test4": fmt.Println(desValue, 44444444444)
		case desValue == "test5": fmt.Println(desValue, 55555555555)
	default:
		fmt.Println(desValue, "default")
	}
}