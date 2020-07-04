package main
import "fmt"

// 循环
func main() {
	// 循环1
	sum := 0
	for i := 0; i<= 10; i++ {
		sum += i
	}
	fmt.Println(sum, "for 1")

	// 循环2
	sum1 := 1
	for sum1 < 10 { //类似while
		sum1 += sum1
	}
	fmt.Println(sum1, "for 2")

	// 循环3
	sum2 := 1
	for ; sum2 < 10; {
		sum2 += sum2
	}
	fmt.Println(sum2, "for 3")

	// 循环4 无限循环
	//sum3 := 1
	//for {
	//	sum3 ++
	//}


	// 循环5 遍历
	arr := []string{"apple", "goggle"}
	for item, value := range arr {
		fmt.Println(item, value)
	}

	arr1 := []int{1, 2, 3, 4}
	for item, value := range arr1 {
		fmt.Println(item, value)
	}
}