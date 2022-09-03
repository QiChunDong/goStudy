package stu

type alterCounter int

// 将工厂函数命名为New是Go语言的一个习惯
// main中可以使用New  因为大写而字母开头 是为公开
// main不能直接使用alterCounter 因为小写开头 视为未公开
func New(value int) alterCounter {
	return alterCounter(value)
}

// 结构体也是一样 未公开的字段不能直接访问
type PubStruct struct {
	Name  string
	email string
}

// 内部类 不公开
type pubStruct1 struct {
	Name  string
	Email string
}

// 公开的结构体中 引入了不公开内部类，但是内部类的字段都是公开的，这样字段会提升到
// 外部类中，外部类对象可以直接定义内部类的公开字段
type PubStruct2 struct {
	pubStruct1
	Level int
}
