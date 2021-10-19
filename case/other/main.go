package main

import (
	"fmt"

	"github.com/wbing441282413/goTest/case/other/entity"
)

var global *int

func A() int {
	var x = 1
	global = &x
	return x
}

func B() int {
	y := new(int)
	*y = 3
	return *y
}
func main() {

	//逃逸分析1
	stu := new(entity.Student)
	stu.Name = "tom" //可以看到"tom"逃逸了
	//优化方案是将类型设置为固定类型
	stu2 := new(entity.Student2)
	stu2.Name = "wbing" //可以看到没逃逸

	//逃逸分析2
	fmt.Println()

	//逃逸分析3
	A() //x逃逸了
	B()

	//逃逸分析4
	x := foo()
	fmt.Println(*x)

	//逃逸分析5
	getStu()
	getStu2()

}
func getStu() *entity.Student2 {
	a := new(entity.Student2) //new(entity.Student2) escapes to heap,指针	发生了逃逸
	a.Name = "aa"
	return a
	/*
		传递指针可以减少底层值的拷贝，可以提高效率，但是如果拷贝的数据量小，
		由于指针传递会产生逃逸，可能会使用堆，也可能会增加 GC 的负担，所以传递指针不一定是高效
	*/
}

func getStu2() entity.Student2 {
	b := entity.Student2{}
	b.Name = "bb"
	return b
}

func foo() *int {
	t := 3
	return &t
}

/*
	使用 go run -gcflags '-m -l' main.go

	逃逸分析是发生在编译阶段

	为什么需要了解逃逸分析？

	通过逃逸分析我们能够知道变量是分配到堆上还是栈上，
	如果分配到栈上，内存的分配和释放都是由编译器进行管理，分配和释放的速度非常快，
	如果分配到堆上，堆不像栈那样可以自动清理，它会引起频繁地进行垃圾回收（GC），
	而垃圾回收会占用比较大的系统开销。

什么是逃逸分析？
在编译程序优化理论中，逃逸分析是一种确定指针动态范围的方法，简单来说就是分析在程序的哪些地方可以访问到该指针。

简单的说，它是在对变量放到堆上还是栈上进行分析，该分析在编译阶段完成。
如果一个变量超过了函数调用的生命周期，也就是这个变量在函数外部存在引用，编译器会把这个变量分配到堆上，这时我们就说这个变量发生逃逸了。

逃逸的变量需要额外分配内存，同时对性能的优化可能会产生细微的影响。

*/
