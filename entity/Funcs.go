package entity

//有名函数定义，函数名是add
//add 类型是函数字面量类型 func(int, int) int
func Adds(a, b int) int {
	return a + b
}

//函数声明语句，用于 Go 代码调用汇编代码
//func adds(int, int) int
//add 函数的签名，实际上就是 add 的字面量类型
//func (int, int) int
//匿名函数不能独立存在，常作为函数参数、返回值，或者赋值给某个变量
//匿名函数可以直接显式初始化2
//匿名函数的类型也是函数字面量类型 func (int, int) int
//func (a,b int) int {
//	return a+b
//}
//新定义函数类型ADD
//ADD 底层类型是函数字面量类型 func (int, int) int
type ADD func(int, int) int
