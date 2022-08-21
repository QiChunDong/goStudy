package stu

import "fmt"

type Duration int64

func TestStruct() {
	var i int64 = int64(1000)
	dur := Duration(i)
	fmt.Println(dur)
}
