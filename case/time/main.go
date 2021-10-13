package main

import (
	"fmt"
	"time"

	"github.com/wbing441282413/gotest/case/time/timeutils"
)

func main() {
	//时间相关
	var currentTime time.Time = time.Now()
	fmt.Println(currentTime)

	timeutils.TimeSomething()

}
