package main

import (
	"fmt"
)

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
	fmt.Println(f6())

}

func f1() int { // 匿名返回值，不会影响原返回值
	i := 0
	defer func() {
		i++
	}()
	return i
}

func f2() int { // 匿名返回值，不会影响原返回值
	i := 0
	defer func(i int) {
		i++
	}(i)
	return i
}

func f3() (re int) { // 具名返回值，是可以影响返回值的
	re = 0
	defer func() {
		re++
	}()
	return re
}

func f4() (re int) { // 具名返回值，没有影响返回值
	re = 0
	defer func(re int) {
		re++
	}(re)
	return re
}
func f5() (re []int) { // 具名返回值，影响返回值
	re = []int{1, 1}

	defer func(re []int) {
		re[1] = 0
	}(re)
	return re
}

func f6() (re []int) { // 具名返回值，影响返回值
	re = []int{1, 1}
	defer func() {
		re[1] = 0
	}()
	return re
}
