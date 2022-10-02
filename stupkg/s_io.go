package stupkg

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func TestIO() {
	// bytes的Buffer实现了io.Write和io.Reade接口
	// 可以使用Write方法写入数据
	var b bytes.Buffer
	b.Write([]byte("hello"))

	// Fprint需要一个是实现了io.Write和io.Reade接口的对象
	// 所以传入的b是byte.Buffer能满足参数要求
	// 追加内容！
	fmt.Fprint(&b, " world!")
	// writeTo 也是io的方法，可以直接用
	// 可以写出道标准输出
	// 当然也可以写出到自定义的设备和文件
	// multi也是支持滴！！！
	b.WriteTo(os.Stdout)
}


func TestCurl(url string, destFile string) {
	r, err := http.Get(url)
	if err != nil {
		log.Fatal("请求异常：", err)
	}

	// 结果输出的目的地
	file, err := os.Create(destFile)
	if err != nil {
		log.Fatalln("目的文件创建异常：", err)
	}
	defer file.Close()

	// 输出到多个设备
	dest := io.MultiWriter(file, os.Stdout)

	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Println("读取响应异常：", err)
	}
}