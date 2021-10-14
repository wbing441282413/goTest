package interface_entity

import "fmt"

type Dog struct {
}
type Bird struct {
}

type Walk interface {
	Walk()
}
type Fly interface {
	Fly()
}

func (d *Dog) Walk() {
	fmt.Println("狗走路")
}
func (b *Bird) Fly() { //鸟走和非两个接口都实现了
	fmt.Println("鸟走路")
}

func (b *Bird) Walk() {
	fmt.Println("鸟飞")
}
