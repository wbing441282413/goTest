package main

import (
	"errors"
	"fmt"
)

/**
type error struct{
	Error（） string
}
实现error接口的都可以作为error
*/
func main() {
	fmt.Println("----------------------------error 的 创建--------------------------------------")
	fmt.Println("======new创建====")
	err := errors.New("xxxx") // 实际是返回了一个*errorString实例，因为*errorString实现了接口error
	fmt.Println(err)

	fmt.Println("======Errorf创建====")
	err = fmt.Errorf("xxxxx%d", 1) //适于需要格式化输出的error，性能不好，一般不建议用
	fmt.Println(err)

	fmt.Println("----------------------------error 的 断言--------------------------------------")
	e := &MyError{content: "太帅了是我的错"}
	fmt.Println(assertError(e))
	fmt.Println(assertError(err))

}

type MyError struct { // 自定义异常
	content string
}

func (e *MyError) Error() string {
	return e.content
}

func assertError(err error) bool {
	if v, ok := err.(*MyError); ok {
		fmt.Println("是MyError异常: ", v)
		return true
	}
	return false
}
