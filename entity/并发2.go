package entity

import (
	"fmt"
	"time"
)

/**
提现并发的案例，体会一下并发
*/

func Runnning() { //普通函数
	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

//然后去main函数中看下并行
