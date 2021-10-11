package main

import (
	"fmt"

	"github.com/wbing441282413/gotest/case/struct/strcut_entity"
)

func main() {
	stu := new(strcut_entity.Student) //这里还是可以引用entity包下的东西的
	stu.Name = "王兵"

	fmt.Println(stu.Name)

	strcut_entity.Say()

	// fmt.Println("student:	", student.GetName(*student))

	// student2 := student.Student{1, "sdsd"}
	// student.Swap(&student2, student) //与结构体相关的不能直接调用
	// fmt.Println("student:	", student.GetName(*student))

	//var s Stu
	s := new(Stu)
	pp := Stu{1, "wb"}
	var t Stu
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", t)
	fmt.Printf("%T\n", pp)
	var pe = &Stu{}
	fmt.Printf("%T\n", pe)

}

type Stu struct {
	id   int
	name string
}
