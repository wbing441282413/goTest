package main

import (
	"fmt"

	"github.com/wbing441282413/gotest/case/func/utilfunc"
)

func main() {
	//类似于Java中的匿名函数
	utilfunc.ToSay(func() {
		fmt.Println("我在帮别人说话")
	}) //ToSay函数的参数是一个函数，这里我们声明了一个匿名函数作为参数
}
