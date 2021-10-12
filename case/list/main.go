package main

import (
	"container/list"
	"fmt"
)

func main() {
	/**
	list 的内部实现是双链表
	*/
	//初始化方式1
	L1 := list.New() //得到指针，是个引用
	//初始化方式2
	var L2 list.List //得到不是引用
	fmt.Printf("%T\n", L1)
	fmt.Printf("%T\n", L2)
	fmt.Printf("%p\n", L1)
	fmt.Printf("%p\n", &L2)

	//列表的元素可以是任意值
	//插入
	fmt.Println("------------------------插入-----------------------")
	fmt.Println("----从首部插入----")
	e1 := L1.PushFront(1) //e1就是插入的节点的引用，插入后会返回
	fmt.Println(L1)
	fmt.Println("----从尾部插入----")
	e2 := L1.PushBack("Sad")

	fmt.Println("------------------------遍历-----------------------")
	// fmt.Printf("%v", L1)//这样是输不出list的
	//顺序输出
	PrintList(L1)
	PrintListReverse(L1)

	//在指定节点插入,InsertBefore，InsertAfter前后关系是按正序来的呦
	fmt.Println("------------------------指定节点位置插入-----------------------")
	L1.InsertBefore("Y", e1)
	L1.InsertAfter("X", e2)
	PrintList(L1)

	//更多的用法用到了在深入学习
	// TODO
}

func PrintList(L1 *list.List) {
	fmt.Println("----正序输出----")
	for e := L1.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func PrintListReverse(L1 *list.List) {
	fmt.Println("----倒序输出----")
	for e := L1.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

}
