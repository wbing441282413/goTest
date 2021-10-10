package main

import (
	"fmt"
	"github.com/wbing441282413/gotest/case/struct/student"
)

func main() {
	stu := new(student.Student)
	stu.Name = "王兵"

	fmt.Println(stu.Name)

	mp := make(map[string]int)
	mp["A"] = 1
	mp["b"] = 2
	delete(mp, "A")
	for name, age := range mp {
		fmt.Println(name)
		fmt.Println(age)
	}

	student.Say()

	// fmt.Println("student:	", student.GetName(*student))

	// student2 := student.Student{1, "sdsd"}
	// student.Swap(&student2, student) //与结构体相关的不能直接调用
	// fmt.Println("student:	", student.GetName(*student))

}
