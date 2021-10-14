package entity

type Movie struct {
	Title  string
	Year   int  `json:"released"`        //后面这一串是结构体成员Tag
	Color  bool `json:"color,omitempty"` //后面这一串是结构体成员Tag,第一个参数是JSON的名字（这里是color），第二个参数omitempty选项，表示当Go语言结构体成员为空或零值时不生成该JSON对象
	Actors []string
}
