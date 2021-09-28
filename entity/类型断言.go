package entity

import "fmt"

func Yuu(s string) {
	fmt.Println("----", s)
}

type YYY func(string)

func HJ(it interface{}) {
	//YYY(it)("hai")//这样是会报错的，调用一个函数的时候，如果传入的参数是个函数，且定义函数的时候函数的参数类型被指定为空接口，是不能直接调用的
	if v, ok := it.(YYY); ok { //必须要先断言空接口的类型才可以
		v("hai")
	} else {
		fmt.Println("不是YYY类型断言调用失败")
	}
}

//func main(){
//	HJ(YYY(Yuu))//传入函数作为参数
//}

func XXC(int) {}
