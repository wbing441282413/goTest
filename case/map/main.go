package main

import (
	"github.com/wbing441282413/gotest/case/map/map01"
)

//好像必须写完整，不然会报错
func main() {
	//待查询的数据：
	//数据内容不重复

	lis := []map01.Profile{
		{Name: "张三", Age: 23, Address: "cq"},
		{Name: "李四", Age: 25},
		{Name: "王五"},
	}
	//传统查询
	map01.FindData(lis, "张三", 23)

	//利用map的多键索引查询（组合键查询）
	map01.BuildIndex(lis)     //构建基于查询的组合键（name、age)
	map01.QueryData("张三", 23) //依据name、age进行查询（多条件）

	/**
	想在这里用mod引入另一个模块struct的student，好像不太容易，会出现ambiguous import的问题，因为
	如果要把项目分成多个模块，那每个模块都应该有自己的单独的tag，并且在引入的时候，require的时候要把tag指明
	这样才不会出现ambiguous import的问题，可以看下https://zhuanlan.zhihu.com/p/134184461，有说到
	*/
	// stu := new(student.Student)
	// stu.Name = "王兵"
	// fmt.Println(stu.Name)

}
