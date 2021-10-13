package map01

/**
map多键索引
*/

import (
	"fmt"
)

type queryKey struct {
	Name string
	Age  int
}
type Profile struct {
	Name    string
	Age     int
	Address string
}

var mapper = make(map[queryKey]Profile)

func BuildIndex(list []Profile) {
	for _, profile := range list {
		key := queryKey{ //查询的组合键为Name、Age构建的结构体
			Name: profile.Name,
			Age:  profile.Age,
		}
		mapper[key] = profile
	}
}
func QueryData(name string, age int) {
	key := queryKey{Name: name, Age: age}
	result, ok := mapper[key]
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("没有找到对应的数据")
	}

}
func FindData(list []Profile, name string, age int) {
	for _, data := range list {
		if data.Name == name && data.Age == age {
			fmt.Println(data)
			return
		}
	}
	fmt.Println("没有找到对应的数据")
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

func ChangeMap(m map[int]string) map[int]string {
	m[1] = "ggg"
	return m
}
