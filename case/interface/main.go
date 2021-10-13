package main

import (
	"fmt"

	"github.com/wbing441282413/gotest/case/interface/interface_entity"
)

func main() {

	fmt.Println("------------------接口的实现--------------------")
	song := interface_entity.Song{"哈哈哈"}
	singer := interface_entity.Singer{3, song}
	fmt.Println(song.GetSongName())

	singer.Intro() //绑定了结构体的方法必须要用实例来调用

	var sing interface_entity.Sing //声明一个接口
	sing = singer                  //结构体Singer实现了接口Sing，所以可以赋值给Sing接口
	sing.Sing()

	//不论是定义在哪里的接口，只要绑定了就能用
	var tt interface_entity.Tt
	tt = singer
	tt.Tt()
	tt = song
	tt.Tt()

	fmt.Println("------------------空接口------1--------------")
	mpp := make(map[interface{}]interface{}) //可以替代任何类型
	mpp["sd"] = "Sda"
	mpp[1] = "Sd"
	mpp[2] = 1
	fmt.Println(mpp)
	fmt.Printf("%T", mpp["sd"])
	fmt.Printf("%T", mpp[1])
	fmt.Printf("%T \n", mpp[2])

	fmt.Println("------------------空接口-----断言----------------")
	vv, kkk := mpp[2].(int) //空接口断言，不是int则kkk为false
	if kkk {
		fmt.Println(vv)
	}

	fmt.Println("------------------接口-----接口的接收者----------------")
	var pxp = new(interface_entity.Phone) //new得到的是结构指针，new不会进行默认的初始化工作
	pxp.Name = "小米"
	var us interface_entity.Usber //得到的是结构体
	us = pxp                      //Start接口是值，所以这里结构体也可以赋值给接口
	us.Start()

}
