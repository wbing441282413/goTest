package entity

import (
	"fmt"
	"math"
)

/**
可以自定义错误来返回
*/

type myError struct {
	Num float64
	//problem string
}

//error接口长这样
//type error interface {
//	Error() string
//}
func (e myError) Error() string { //实现了error接口
	return fmt.Sprintf("%f是负数，不允许开根号", e.Num) //怎么调用到这的呢？是在主程序中调用输出error的时候调用的，print方法会根据是否为error接口来输出对应的 内容，debug可以看到
}
func Sqrt(f float64) (float64, error) { //myErro实现了error接口才能怎么用
	if f < 0 {
		return -1, myError{Num: f}
	}
	return math.Sqrt(f), nil
}
