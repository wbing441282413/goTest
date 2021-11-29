package main

import (
	"fmt"
)

func main() {
	foo()
	fmt.Println("main结束运行") // 如果子函数的panic都捕获了，就可以执行这句
}

/**
  panic的捕获顺序，defer的执行顺序
*/
func foo() {
	defer func() { // 5
		if e := recover(); e != nil { // 捕获的是panic A
			fmt.Println(e)
		}
	}()

	defer func() { // 2
		defer func() { // 捕获的是panic B // 4 (捕获的是3处的panic)
			if e := recover(); e != nil {
				fmt.Println("inner --", e)
			}
		}()

		fmt.Println("B")
		panic("panic-01 B : inner defer") // 3
		fmt.Println("BB")
	}()

	fmt.Println("A")
	panic("panic-01 A : foo ") // 1
	fmt.Println("AA")
	return
}
