package main

import "fmt"

func main() {
	//if语句
	if 1 == 1 && 1 == 1 {
		fmt.Println("ccccccc")
	}

	//switch语句
	fmt.Println("------------switch--------1-----")
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

	fmt.Println("------------switch--------2-----")
	var ax = 1
	switch ax {
	case 1, 2:
		fmt.Println("xxxxxx")
	case 3:
		fmt.Println("ccccc")
	}

	fmt.Println("------------range--------1-----")
	ss := "sdsd"
	for _, e := range ss {
		fmt.Printf("%c ", e)
	}

	fmt.Println("------------range--------2-----")
	zx := make(map[string]int)
	zx["w"] = 1
	zx["c"] = 2
	for k, v := range zx {
		fmt.Println(k, v)
	}
}
