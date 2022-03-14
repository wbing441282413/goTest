package main

import (
	"fmt"
	"reflect"
)

type Stu struct {
	name string
	subj []string
}
type Tec struct {
	age int
}

func isEqualUseless(a interface{}, b interface{}) bool {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常： ", err)
		}
	}()
	return a == b
}

func isEqualUseful(a interface{}, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	s1 := Stu{name: "wb", subj: []string{"math", "music"}}
	s2 := Stu{name: "wb", subj: []string{"math", "music"}}
	t1 := Tec{age: 1}
	t2 := Tec{age: 1}

	fmt.Println("-------------------------比较两个结构体中的变量--------------------------------------------------------")
	fmt.Println("====用==比较结构体=======")
	//fmt.Println("结构体是否相等？ ",s1 == s2) // 结构体中slice，是不能用==比较的成员
	fmt.Printf("是否是同一个结构体？ %v\n", &s2 == &s1)
	fmt.Println("结构体是否相等？ ", t1 == t2) // 结构体中没有不能使用==的成员，可以用==比较结构体
	fmt.Printf("是否是同一个结构体？ %v\n", &t2 == &t1)

	fmt.Println("====用reflect比较结构体=======")
	fmt.Println("结构体是否相等？ ", isEqualUseful(s1, s2))
	fmt.Printf("是否是同一个结构体？ %v\n", &s2 == &s1)
	fmt.Println("结构体是否相等？ ", isEqualUseful(t1, t2))
	fmt.Printf("是否是同一个结构体？ %v\n", &t2 == &t1)

	/**
	  反射其实是对interface的操作，interface是由有个（value， type）对组成的
	  反射就是操作这个，
	  reflect的reflect.Type和reflect.value两个对应的就是interface的（type，value），也叫反射对象
	  转换通过valueof和typeof实现

	*/
	fmt.Println("-------------------------第一定律：反射可以把interface对象转为反射对象--------------------------------------------------------")
	x := 1
	stu := Stu{name: "wb", subj: []string{"a", "b"}}
	// 通过反射把interface中的vale和type取出来变为反射对象（其实这里还是从interface中去取的，会强转为interface传入方法的）
	valuee := reflect.ValueOf(stu) // 其实这里是会把stu转为的interface的传入valueof方法的
	typee := reflect.TypeOf(stu)
	// 得到的是反射对象
	fmt.Printf("%v \n", valuee)
	fmt.Printf("%v \n", typee)

	fmt.Println("-------------------------第二定律：反射可以把反射对象转为interface对象--------------------------------------------------------")
	// 反射对象转interfacecon
	xi := valuee.Interface() // 反射对象才有这方法
	fmt.Printf("%v \n", xi)
	fmt.Printf("%T \n", xi)
	fmt.Println(x == xi)

	fmt.Println("-------------------------第三定律：反射可以修改interface对象的值--------------------------------------------------------")
	//s := 2
	//vc := reflect.ValueOf(s)
	//fmt.Printf("%T \n", vc)
	//vc.SetInt(3) // 会错，panic: reflect: reflect.Value.SetInt using unaddressable value
	//// 因为修改的不是地址会是无效的修改
	//fmt.Println(vc.Interface())

	fmt.Println("======正确的修改是要地址===")
	b := 0
	va := reflect.ValueOf(&b)
	fmt.Printf("%T \n", va)
	va.Elem().SetInt(4) // va.Elem()相当于是取值*p, 反射对象对应的inteface必须是指针类型的才能调用
	fmt.Printf("%T \n", va.Interface())
	fmt.Printf("%T \n", va.Elem().Interface())
	fmt.Println(va.Elem().Interface())

	// defer和recover直接放在main函数中，是不会有recover的效果的
}
