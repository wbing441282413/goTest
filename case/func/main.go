package main

import (
	"fmt"

	"github.com/wbing441282413/gotest/case/func/utilfunc"
)

func main() {

	fmt.Println("---------------------------匿名函数作形参------------------------------")

	//类似于Java中的匿名函数
	utilfunc.ToSay(func() {
		fmt.Println("我在帮别人说话")
	}) //ToSay函数的参数是一个函数，这里我们声明了一个匿名函数作为参数

	fmt.Println("---------------------------------------------------------")

}

/**

关于函数签名和函数字面量类型：
函数签名”就是“有名函数”或“匿名函数”的字面量类型
对于 type NewFuncType FuncLiteral来说：
NewFuncType 为新定义的函数命名类型
FuncLiteral 为函数字面量类型，也就是常说的函数签名
函数字面量类型，函数字面量类型的语法表达格式是 func (InputTypeList)OutputTypeList
//函数声明＝函数名＋函数签名
//函数签名
func (InputTypeList)OutputTypeList
//函数声明
func FuncName (InputTypeList)OutputTypeList


举例子：
函数声明：
func add(int,int)int
函数签名,也叫作函数字面量类型：也即是在定义一个函数的时候，这个函数的类型就是他的函数签名
func (int,int) int
函数类型，也叫作函数命名类型(不过注意是类型不是实例呦)：
type ADD func(int int) int
函数的定义：
 func add(a, b int) int {
     return a+b
 }

定义的函数可以赋值给函数类型：前提是函数签名要对应一样
var d ADD
d = add

*/
