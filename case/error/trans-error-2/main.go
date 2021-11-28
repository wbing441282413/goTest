package main

import (
	"errors"
	"fmt"
	"reflect"
)

/**
在trans-error中，通过成员变量来保存原来的error，其实可以通过链式error来实现
*/
func main() {
	// wrapError
	/**
	其实wrapError也是通过成员变量保存原error
	type wrapError struct{
		msg string
		err error
	}
	*/
	//创建wrapError， 通过Errorf来创建, Errorf格式动词是%w删除的就是wrapError
	err := errors.New("xxxx")
	wrapErr := fmt.Errorf("wrapppp ,%w", err) // 返回的error是wrapError

	//if _, ok := wrapErr.(wrapError); ok { //wrapError是私有的，所以这里应该这么断言
	if _, ok := wrapErr.(interface{ Unwrap() error }); ok { //wrapError是私有的，所以这里应该这么断言
		fmt.Println("wrapErr是一个WrapError")
		//v := reflect.ValueOf(wrapErr)
		//fmt.Printf("%v",v)
	}

	fmt.Println("------------------------------wrapError的成员----------------------------------------------------")
	e := &MyError{content: "太帅的错"}
	e2 := fmt.Errorf("太优秀的错，%w", e) // 可以打断点在这看，wrapError的成员err是这里的e, 成员msg是新的，也就是这里拼接的内容
	e3 := fmt.Errorf("太善良的错，%w", e2)
	v := reflect.ValueOf(e3)
	t := reflect.TypeOf(e3)
	fmt.Printf("%v \n", v)
	fmt.Printf("%v \n", t)
	fmt.Printf("%T \n", e3)
	fmt.Println(e3.Error()) // TODO 不知道为啥不能调用Unwrap,书上说的不明白
	fmt.Printf("%T \n", v.Elem().Interface())
	fmt.Printf("%T \n", v.Interface())

	//要调用Unwrap,要用Unwrap方法， 自定义的error也实现Unwrap方法
	fmt.Println("------------------------------Unwrap的用法----------------------------------------------------")
	fmt.Printf("e的成员，原error是：%T \n", errors.Unwrap(e))
	fmt.Printf("e2的成员，原error是：%T \n", errors.Unwrap(e2))
	fmt.Printf("e3的成员，原error是：%T \n", errors.Unwrap(e3))

	fmt.Println("------------------------------判断链中是否有某error----------------------------------------------------")
	fmt.Println(errors.Is(e3, e)) // 如果传入的e3实现了Is方法，则会先访问的自定义的Is方法

	var m *MyError
	fmt.Println(errors.As(e3, &m)) // As方法更棒，传入的是interface， 而且如果匹配到了还回写入到m中 ,这里传入的是指针的地址
	fmt.Println(m)
}

type MyError struct {
	content string
	error   error
}

func (e *MyError) Error() string {
	return e.content
}

func (e *MyError) Unwrap() error {
	return e.error
}
