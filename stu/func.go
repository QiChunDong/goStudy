package stu

import "fmt"

// 用户类型
type user struct {
	name  string
	email string
}

// 向用户使用者发送通知
func (u user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// 向用户指针和对象都可以修改email
func (u *user) changeEmail(email string) {
	u.email = email
}

func TestFunc() {
	// user类型的值可以调用使用者声明的发给发消息
	bill := user{"Bill", "bill@email.com"}
	bill.notify()
	bill.changeEmail("bill1111@email.com")
	bill.notify()

	// user的指针  可以调用change的
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()
	lisa.changeEmail("lisa1111@email.com")
	lisa.notify()
}
