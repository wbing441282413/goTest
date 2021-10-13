package utilfunc

func ToSay(f func()) { //参数是一个函数
	f() //函数体的内容是：调用传入的这个函数
}

func ChangeMap(m map[int]string) map[int]string {
	m[1] = "ggg"
	return m
}
