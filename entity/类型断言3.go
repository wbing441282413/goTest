package entity

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
func (b *Bird) Fly() {
	fmt.Println("鸟走路")
}

func (b *Bird) Walk() {
	fmt.Println("鸟飞")
}
