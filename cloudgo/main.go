package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                 // 解析Form（即URL参数）
	if len(r.Form["name"]) == 0 { // 判断是否存在参数 name
		fmt.Fprintf(w, "Hello, but who are you?\n")
	} else {
		fmt.Fprintf(w, "Hello, %s!\n", r.Form["name"][0])
		// 注意，这里的 name 参数对应的值是一个 string 数组，意味着可以接收多个名为 name 的重复参数
	}
}

func main() {
	http.HandleFunc("/", hello)              // 在"/"路径挂载 hello 函数进行监听
	err := http.ListenAndServe(":3000", nil) // 监听3000端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
