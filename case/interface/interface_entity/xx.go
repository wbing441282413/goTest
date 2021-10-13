package interface_entity

import "fmt"

type Service interface { //接口有2个方法
	Start()
	Stop()
}

type Aa struct { //Aa中内嵌了Bb
	Bb
}
type Bb struct{}

func (Aa) Start() {
	fmt.Println("我是Aa。我亲自实现了start方法")
}

func (Bb) Stop() {
	fmt.Println("我是Bb。我亲自实现了stop方法")
}

type Habe interface {
	Show()
}

func (Bb) Show() { //Bb实现了接口Habe,Aa没有实现但是Aa内嵌了Bb,Bb实现了就好比Aa实现了Habe
	fmt.Println("我还有另外一个接口Habe的实现方法，看下Aa可以调用吗?验证了Aa是可以调用的")
}
