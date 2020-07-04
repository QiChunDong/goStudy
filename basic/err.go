package main

import "fmt"

// 练习的是异常处理

// 异常要实现异常的接口
// 这是异常的源码
//type error interface {
//	Error() string
//}

// 定义异常结构体
type DivideError struct {
	dividee int
	divider int
}
// 实现Error接口
func (de *DivideError) Error() string {
	strFmt := `
	不能计算
	被除数： %d
	除数： 0
	`
	return fmt.Sprintf(strFmt, de.dividee)
}

func Ddvide(devidee int, devider int) (result int, errMsg string) {
	if devider == 0 {
		dData := DivideError{
			dividee: devidee,
			divider: devider,
		}
		errMsg = dData.Error()
		return
	} else {
		return devidee / devider, ""
	}
}

func main () {
	res, err := Ddvide(10, 2)

	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}