package main

import (
	"fmt"
	"go/types"

	"github.com/wbing441282413/gotest/case/interface/interface_entity"
)

func main() {

	fmt.Println("------------------接口的实现--------------------")
	song := interface_entity.Song{SongName: "哈哈哈"}
	singer := interface_entity.Singer{Count: 3, Song: song}
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

	fmt.Println("------------------空接的类型-------------------")
	var yu interface{}
	if yu == nil {
		fmt.Println("空接口等于nil")
	}

	switch yu.(type) {
	case types.Nil:
		fmt.Println("空接口类型是Nil")
	default:
		fmt.Println("unknown type,空接口==nil,但是它不是Nil类型的")
	}

	fmt.Println("----------------------------------接口-----接口的接收者----------------------------")
	var pxp = new(interface_entity.Phone) //new得到的是结构指针，new不会进行默认的初始化工作
	pxp.Name = "小米"
	var us interface_entity.Usber //得到的是结构体
	us = pxp                      //Start接口是值，所以这里结构体也可以赋值给接口
	us.Start()

	fmt.Println("-------------不一定要一个类型全部实现接口的所有方法，通过内嵌实现的我也可以用----------------")
	//不一定要一个类型全部实现接口的所有方法，可以通过内嵌其他按类型，让其他类型实现的也可以为我所用
	AF := new(interface_entity.Aa)
	var sssd interface_entity.Service = AF
	sssd.Start()
	sssd.Stop() //可以看到AF也可以调用Stop,虽然Aa没有实现Stop方法，但它的成员Bb实现了Stop方法，所以它也可以调用
	var xxxd interface_entity.Habe = AF
	xxxd.Show() //可以看到AF也可以调用Show,虽然Aa没有实现Show方法，但它的成员Bb实现了Show方法，所以它也可以调用

	fmt.Println("-------------------------自己和内嵌都实现了接口会调用哪个--优先调自己的-----------------------------------")
	//自己和内嵌都实现了接口会调用哪个
	gh := new(interface_entity.YY)
	var xxa = gh
	xxa.HH()

}
