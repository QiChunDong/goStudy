package main

import (
	"fmt"
	"gostudy/stupkg"
	"os"
)

// init main之前执行
func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./example2 <url>")
		os.Exit(-1)
	}
}

// main方法
func main() {
	fmt.Println("hey world!!!")

	// fmt.Println("-----------------BMI  MY TEST CODE")
	// args := os.Args
	// var W float64
	// var H float64
	// fmt.Sscanf(args[1], "%f", &W)
	// fmt.Sscanf(args[2], "%f", &H)
	// t.Printfln("身高，%.2f", H)
	// t.Printfln("体重，%.2f", W)
	// pract.ClacBMI(W, H)

	// fmt.Println("-----------------int64")
	// stu.TestStruct()

	// fmt.Println("-----------------func")
	// stu.TestFunc()

	// fmt.Println("-----------------curl")
	// 测试url curl 发送请求
	// r, err := http.Get(os.Args[1])
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// 从body取值
	// io.Copy(os.Stdout, r.Body)
	// if err := r.Body.Close(); err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("-----------------copy")
	// stu.TestCopy()

	// fmt.Println("-----------------测试公开性")
	// 未公开不能直接使用
	// pub1 := stu.alterCounter(1)
	// pub2 := stu.New(1)
	// fmt.Printf("val1 is %d\n", pub2)
	// pub3 := stu.PubStruct{
	// 	Name:  "111",
	// 	email: "11111", // 编译报错 unknown field 'email' in struct literal of type stu.PubStruct
	// }
	// fmt.Println(pub3)
	// pub4 := stu.PubStruct2{
	// 	Level: 1,
	// }
	// pub4.Name = "11111"
	// pub4.Email = "22222"
	// fmt.Printf("pub4: %v\n", pub4)

	// fmt.Println("-----------------测试协程")
	// 打印字母
	// stu.TestRoutine()
	// 打印素数
	// stu.TestPrime()
	// 测试竞争
	// stu.TestComp()
	// 测试加锁
	// stu.TestAtom()
	// 测试互斥锁mutex
	// stu.TestMutex()

	// fmt.Println("-----------------测试通道")
	// 无缓冲的通道
	// stu.TestNoBufferChanel()
	// 无缓冲的通道 2
	// stu.TestNoBufferChanel2()
	// 有缓冲的通道
	// stu.TestBufferChanel()


	// fmt.Println("-----------------测试并发")
	// stupkg.TestRunner()
	// fmt.Println("-----------------测试pool")
	// stupkg.TestPool()
	// fmt.Println("-----------------测试pool")
	// stupkg.TestPool()
	// fmt.Println("-----------------测试work")
	// stupkg.TestWorker()
	// fmt.Println("-----------------测试Log")
	// stupkg.TestLog()
	// fmt.Println("-----------------测试Unmarshal")
	// stupkg.TestUnmashal()
	// fmt.Println("-----------------测试UnmarshalMap")
	// stupkg.TestUnmashalMap()
	// fmt.Println("-----------------测试TestJsonEncode")
	// stupkg.TestJsonEncode()

	// fmt.Println("-----------------测试 TestIO")
	// stupkg.TestIO()
	fmt.Println("-----------------测试 Curl")
	stupkg.TestCurl("https://www.baidu.com/s?wd=%E7%99%BE%E5%BA%A6%E6%90%9C%E7%B4%A2", "resp.txt")
}
