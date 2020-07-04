package main

import "fmt"

/**
  练习Go的接口
 */

type Phone interface {
	call()
}

type IPhone struct {

}
func (iphone IPhone) call() {
	fmt.Println("i am iphone")
}

// 可以这么理解
// 接口是面向方法
// 每个方法可以对不同的类实现不同的处理方法体
type Nokia struct {

}
func (nokia Nokia) call() {
	fmt.Println("i am nokia")
}

func main() {
	var phone Phone

	// 创建的时候指定实现类
	phone = new(Nokia)
	// Nokia的方法处理对应的是Nokia的call
	phone.call()

	phone = new(IPhone)
	phone.call()
}