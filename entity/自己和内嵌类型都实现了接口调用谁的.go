package entity

import "fmt"

type YY struct {
	XX
}
type XX struct{}

type GS interface {
	HH()
}

//都实现了看看YY的实例会调用谁的
func (y YY) HH() {
	fmt.Println("我是YY实现的")
}

func (X XX) HH() {
	fmt.Println("我是xx实现的")
}
