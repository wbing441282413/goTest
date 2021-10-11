package main

import (
	"fmt"

	"github.com/wbing441282413/gotest/case/interface/interface_entity"
)

func main() {

	song := interface_entity.Song{"哈哈哈"}
	singer := interface_entity.Singer{3, song}
	fmt.Println(song.GetSongName())
	singer.Intro()

	var sing interface_entity.Sing
	sing = singer
	sing.Sing()

	//不论是定义在哪里的接口，只要绑定了就能用
	var tt interface_entity.Tt
	tt = singer
	tt.Tt()
	tt = song
	tt.Tt()
}
