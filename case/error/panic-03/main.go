package main

import (
	"fmt"
)

/**
  recover的捕获范围
*/
func main() {
	foo()
	fmt.Println("main结束运行")
}

func foo() {
	defer func() { // x
		defer func() { // y
			if e := recover(); e != nil { // recover是不能跨层调用的，这一层没有panic，那这层的recover就没用了，不会留着去捕获外层的panic
				fmt.Println(e)
			} // y这defer只能用于捕获x函数里的panic，没捕获到并不会留着去捕获foo的panic
			fmt.Println("c")
		}()
		fmt.Println("B")
	}()

	fmt.Println("A")
	panic("panic-01 A ： foo")
	fmt.Println("AA")
}
