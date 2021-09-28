package utils

func ToSay(f func()) {
	f()
}

func ChangeMap(m map[int]string) map[int]string {
	m[1] = "ggg"
	return m
}
