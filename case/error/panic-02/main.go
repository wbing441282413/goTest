package main

import (
	"fmt"
)

func main() {

	foo()
	fmt.Println("main结束运行")
}
func foo() {
	defer func() {
		if e := recover(); e != nil { // 捕获的是panic A，  因为这里捕获了所以main这可以在此函数的调用处进行向下执行，所以会输出main结束运行
			fmt.Println(e)
		}
	}()

	fmt.Println("A")
	panic("panic-01 A : foo ") // panic是运行时异常 ，直接转到执行defer函数
	fmt.Println("AA")          // 不会执行
	return
}
