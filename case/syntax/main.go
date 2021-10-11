package main

import "fmt"

func main() {
	//if语句
	if 1 == 1 && 1 == 1 {
		fmt.Println("ccccccc")
	}

	//switch语句
	var score int = 90
	switch {
	case score >= 90:
		fmt.Println("成绩优良")
		fallthrough
	case score > 70 && score < 90:
		fmt.Println("成绩合格") //如果在case后面使用fallthrough 关键字，则当前case执行完成后会继续执行下一个case，
	case score < 60:
		fmt.Println("成绩不及格")
	}
}
