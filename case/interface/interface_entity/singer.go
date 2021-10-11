package interface_entity

import "fmt"

type Singer struct {
	Count int
	Song  Song
}

func (s Singer) Sing() {
	fmt.Println(s.Song.SongName)
}

func (s Singer) Tt() {
	fmt.Println("tt ")
}

type Sing interface { //接口
	Sing()
}

func (singer Singer) Intro() {
	fmt.Println(singer.Count, singer.Song.SongName)
}
