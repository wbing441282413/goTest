package entity

import (
	"fmt"
	"sort"
)

func SortUse() {
	//对 int 、 float64 和 string 数组或是切片排序
	//可以直接使用sort包的sort.Ints() 、 sort.Float64s() 和 sort.Strings() 函数排序
	//默认从小到大,源码可以看到是用的快排
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}
	sort.Ints(intList)
	sort.Float64s(float8List)
	sort.Strings(stringList)
	fmt.Println(intList)
	fmt.Println(float8List)
	fmt.Println(stringList)

}

//TODO 待续
