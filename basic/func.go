package main
import "fmt"

func main() {
	var v1 = 100
	var v2 = 100
	v3 := test(v1, v2)

	fmt.Println(v1, v2, v3, "res")

	v4, v5 := swap(v1, v2)
	fmt.Println(v4, v5, "res3")

	test1(&v1, &v2)
	fmt.Println(v1, v2, "res1")

	// 闭包
	testB1 := testB(11)
	fmt.Println(testB1(21342), "resB")

	// 测试方法打印 类型的属性
	my := MyType{
		name: "test",
		age:  123,
	}
	var you MyType
	you.name = "clsd"
	you.age = 1323213
	say(my)
	say(you)
}
func swap(x, y int) (int, int) {
	return y, x
}

func test(x, y int) int {
	x += 100
	y += 100
	return x + y
}

func test1(x, y *int) {
	tttt := 1001
	aaaa := 1001
	*x = tttt + 100
	*y = aaaa
}

// 闭包
// 返回值是函数
// 如果函数中带参数 那么声明的返回值类型要在func()里加入参数类型
func testB(a int) func(int) string {
	b := 100
	return func(c int) string {
		a += b
		a += c
		return "ckjsck" + "111"
	}
}


// 方法
type MyType struct {
	name string
	age int
}

func say(v MyType) {
	fmt.Println(v.name, v.age)
}