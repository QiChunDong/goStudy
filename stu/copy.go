package stu

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func TestCopy() {
	var b bytes.Buffer

	// 字符串写入buffer
	b.Write([]byte("hello"))
	// 使用Fprintf将字符串拼接到buffer
	fmt.Fprintf(&b, " World!")
	// buffer 写入stdout
	io.Copy(os.Stdout, &b)
}
