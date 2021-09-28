package entity

import "fmt"

type Usber interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (Phone) Start() { //结构值接收者
	fmt.Println("start")
}

func (Phone) Stop() {
	fmt.Println("start")
}
