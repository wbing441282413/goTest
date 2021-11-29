package main

import (
	"fmt"
)

func main() {
	foo()
	fmt.Println("main结束运行")
}
func foo() {
	defer func() { // 如果f没有捕获f自己的panic，则会捕获f中的panic， 因为这里的recover捕获的是foo中的panic而f也是在foo这调用的
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	f()

	fmt.Println("foooooo") // 如果f捕获了f自己的panic则会输出
}

func f() {
	// defer func() {
	//     if e := recover(); e != nil {
	//         fmt.Println(e)
	//     }
	// }()

	fmt.Println("fff")
	panic("panic-01 f : f") // 如果这没有处理异常，则会在调用者函数中被捕获
}
