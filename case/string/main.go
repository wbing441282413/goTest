package main

import "fmt"

func main() {
	sss, hhh := "sd", "sd"
	if sss == hhh {
		fmt.Println("判断字符串相等可以用==")
	}

	so := []string{"sd", "sd", "ss", "ss", "xx"}

	var sy string
	sy = fmt.Sprintf("%q", so) //转为字符串并返回
	fmt.Println(sy)
}
