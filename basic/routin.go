package main

import "fmt"
import "time"

// 练习并发
// 通过go关键字的方法前缀来开启线程
// 线程由go运行时环境控制

func sayHey(param string) {
	fmt.Println(param)
}

func say10(param string) {
	for i := 0; i < 10; i ++ {
		time.Sleep(100 * time.Millisecond)
		sayHey(param)
	}
}


func main() {
	say10("1")
	go say10("world")
	say10("2")
}