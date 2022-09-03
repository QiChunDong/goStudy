package stu

type alterCounter int

// 将工厂函数命名为New是Go语言的一个习惯
// main中可以使用New  因为大写而字母开头 是为公开
// main不能直接使用alterCounter 因为小写开头 视为未公开
func New(value int) alterCounter {
	return alterCounter(value)
}
