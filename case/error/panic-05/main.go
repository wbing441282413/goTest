package main

import (
	"fmt"
)

/**
  一个原则， 有没有捕获的panic出现在此函数中，则此函数会在此panic处结束执行
*/
func main() {
	foo()
	fmt.Println("main结束运行")
}
func foo() {
	defer func() { // x
		defer func() { // recover的范围是 x函数内，f()自己没捕获panic这里会帮他捕获
			if e := recover(); e != nil {
				fmt.Println(e)
			}
		}()
		f()
		fmt.Println("cc") // 因为f()没有自己捕获f自己的panic，所以这里不会输出
	}()
	fmt.Println("A")
	return
}

func f() {
	fmt.Println("fff")
	panic("panic-01 f : f")
}
