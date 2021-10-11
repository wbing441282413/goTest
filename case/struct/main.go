package main

import (
	"fmt"

	"github.com/wbing441282413/gotest/case/struct/struct_entity"
)

func main() {
	stu := new(struct_entity.Student) //这里还是可以引用entity包下的东西的
	stu.Name = "王兵"

	fmt.Println(stu.Name)

	struct_entity.Say()

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

	//未初始化的结构体
	var xx struct_entity.Singer
	fmt.Println("没初始化的结构体：", xx.Song.SongName, "count:", xx.Count)
}

type Stu struct {
	id   int
	name string
}
