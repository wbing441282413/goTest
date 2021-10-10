package main

import (
	"fmt"

	"github.com/wbing441282413/gotest/case/struct/student"
)

func main() {
	stu := new(student.Student) //这里还是可以引用entity包下的东西的
	stu.Name = "王兵"

	fmt.Println(stu.Name)

	student.Say()

	// fmt.Println("student:	", student.GetName(*student))

	// student2 := student.Student{1, "sdsd"}
	// student.Swap(&student2, student) //与结构体相关的不能直接调用
	// fmt.Println("student:	", student.GetName(*student))

}
