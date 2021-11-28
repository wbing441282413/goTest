package main

import "fmt"

/**
如何把异常向上抛
*/

type MyError struct {
	content string
	error   error
}

type MySubError struct {
	content string
	err     error // 原来的 error 在向上抛出异常的过程中，抛出链上的异常是有关系的，需要保存这种上下文关系
}

func (e *MyError) Error() string {
	return e.content
}

func (e *MySubError) Error() string {
	return e.content
}

func main() {
	//err := getA(1)
	err := getA(2)
	fmt.Println(err)
}

func getWithoutGirlFriendReason(a int) error { //向上抛出异常
	if a == 1 {
		e := &MyError{content: "太帅了是我的错"}
		return e
	}
	return nil
}

func getA(a int) error {
	if a == 1 || a == 2 {
		var er error
		if err := getWithoutGirlFriendReason(a); err != nil {
			er = &MySubError{err: err, content: fmt.Sprintf("--%s \n --太优秀了是我的错", err)}
		} else {
			er = &MySubError{err: nil, content: fmt.Sprintf("--太优秀了是我的错")}
		}
		return er
	}
	return nil
}
