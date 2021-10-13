package utils

func Reverse(s []int) { //在原内存空间进行数组和切片进行翻转
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
