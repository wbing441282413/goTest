package entity

import (
	"log"
	"net/http"
)

func SendToWeb() {
	http.HandleFunc("/index", index) // index 为向 url发送请求时，调用的函数
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
} //url=localhost:8000/index

func index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w,"你好首页")
	w.Write([]byte("你好"))
}





