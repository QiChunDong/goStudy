package tools

import (
	"fmt"
)

func Printfln(format string, a ...interface{}) {
	fmt.Printf(format + "\n", a...)
}