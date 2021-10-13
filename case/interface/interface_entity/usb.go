package interface_entity

import "fmt"

type Usber interface { //定义了2个接口
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (Phone) Start() { //结构值接收者，接口是结构体而不是结构体指针的时候，那在赋值给接口的时候，结构值和结构指针都可以赋值给接口
	fmt.Println("start")
}

//如果结构体接收者是指针，那赋值给接口的时候，必须是指针才能赋给接口

func (Phone) Stop() {
	fmt.Println("start")
}
