package entity

import "fmt"

type Song struct {
	SongName string
}

type Tt interface {
	Tt()
}

func (song Song) Tt() {
	fmt.Println("我是Song,Tt接口在我这边，我绑定了他")
}

func (song Song) GetSongName() string {
	return song.SongName
}
