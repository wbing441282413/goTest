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
