package stu

import "fmt"

// 用户类型
type user struct {
	name  string
	email string
}

// 向用户使用者发送通知
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// 向用户指针和对象都可以修改email
func (u *user) changeEmail(email string) {
	u.email = email
}

type admin struct {
	name  string
	email string
}

func (a *admin) notify() {
	fmt.Printf("Sending email  to admin %s<%s>\n", a.name, a.email)
}

// 将user类型嵌入adminB 测试内部类型
// adminB 是外部类型  user是内部类型
// 如果adminB定义同名属性  可以覆盖user的
type adminB struct {
	user
	level string
}

// func (a adminB) notify() {
// 	fmt.Printf("Sending email  to adminB %s<%s>\n", a.name, a.email)
// }

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

	sendNotification(lisa)

	sam := admin{"sam", "1111@qq.com"}
	sendNotification(&sam)

	lily := adminB{
		user: user{
			name:  "lily",
			email: "1111112@qq.com",
		},
		level: "super",
	}
	lily.user.notify()
	// adminB不自定义notify  adminB也可以调用从内部类提升上来的notify实现
	// 当然如果自定义  可以覆盖
	lily.notify()
	// 也都可以传入方法sendNotification
	// 内外部的值及指针都是接口多态的一部分
	sendNotification(&lily)
	sendNotification(&lily.user)
}

func sendNotification(u notifier) {
	u.notify()
}
