package main

import "fmt"

// 下面的是通道 channel
//ch <- v    // 把 v 发送到通道 ch
//v := <-ch  // 从 ch 接收数据
// 并把值赋给 v
// 定义
//ch := make(chan int)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

// 通道有缓冲区 可以异步执行
// 缓冲区有大小限制 make的时候第二个参数指定缓冲区大小
// 缓冲区满 则源头不能再往缓冲区给值
func testBuffer() {
	c := make(chan int, 2)

	c <- 2
	c <- 5

	fmt.Println(<- c)
	fmt.Println(<- c)
}

// 通道可以遍历
// 通道也可以手动关闭
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i ++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
}
func testFi() {
	c := make(chan int, 10)
	// 并发给值
	go fibonacci(cap(c), c)
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	s := []int{ 7, 3, 3, 5}
	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <- c, <- c
	fmt.Println(x, y, x + y)

	fmt.Println("--buffer----")
	// 测试缓冲区
	testBuffer()
	fmt.Println("--close----")
	// 测试缓冲区容量和关闭和遍历
	testFi()
}