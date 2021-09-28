package entity

import (
	"fmt"
	"time"
)

func Caculate() { //TODO 可能不是这样用协程的，还要在看看
	go doCaculate()
	doCaculate()
}

func doCaculate() {
	start := time.Now().UnixMilli()
	a := 0
	for i := 0; i < 200000; i++ {
		for j := 0; j < 10000; j++ {
			a += j
		}
	}
	fmt.Println(a)
	end := time.Now().UnixMilli()
	fmt.Println("用时：	", end-start)
}
