package main

import "fmt"

func main() {

	var age int = 10
	var p *int = &age
	*p++ //对指针操作直接是对值操作
	x := &p
	fmt.Println(x)
	fmt.Println(age)
}
