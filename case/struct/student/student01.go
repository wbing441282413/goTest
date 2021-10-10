package student

import (
	"fmt"
)

func Say() { //首字母大写才是public
	fmt.Println("hai")
}

type Student struct {
	Age  int
	Name string
}

func GetName(stu Student) string { //无需改变值，可以传结构值,值传递
	return stu.Name
}

func Swap(a *Student, b *Student) { //需改变值，要传结构指针,引用传递
	tp := a.Name
	a.Name = b.Name
	b.Name = tp
}

func SwapDig(a int, b int) (int, int) {
	return b, a
}
