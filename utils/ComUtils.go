package utils

func Reverse(s []int) { //在原内存空间进行数组和切片进行翻转
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Equal(x, y map[string]int) bool { //比较map是否相等
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv { //因为找不到key时map也会返回0，所以不仅仅是比值，还要看一下key存不存在
			return false
		}
	}
	return true
}
