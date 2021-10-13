package timeutils

import (
	"fmt"
	"time"
)

func TimeSomething() {
	curretTime := time.Now()                                  //当前时间
	curretTimeDay := time.Now().Day()                         //当前时间天
	timeUnix := time.Now().Unix()                             //时间戳，单位s
	formatedTime := time.Now().Format(time.RFC3339)           //时间转字符串
	formatedTime2 := time.Now().Format("2006-01-02 15:04:05") //固定写死的这个时间，改了时间就错了，时间转字符串
	formatedTime3 := time.Now().Format("2006-01-02 15:00:00") //固定写死的这个时间，改了时间就错了，时间转字符串
	fmt.Println("当前时间", curretTime)
	fmt.Println("当前时间天", curretTimeDay)
	fmt.Println("时间戳s", timeUnix)
	fmt.Println("格式化时间", formatedTime)
	fmt.Println("格式化时间", formatedTime2)
	fmt.Println("格式化时间", formatedTime3)

	//字符串转时间
	timeStr := "2021-09-24 16:23:18"
	StrtoTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err == nil {
		fmt.Println("字符串转换成时间为：	", StrtoTime)
	}
	//StrtoTime2, err2 := time.Parse(time.RFC3339,timeStr)
	//fmt.Println(err2,StrtoTime2)

}
