package main

import (
	"fmt"
	"net/http"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是一个开始")
}

func main() {
	http.HandleFunc("/", myWeb) //注册路由函数

	fmt.Println("服务器即将开启")

	err := http.ListenAndServe(":8080", nil) //开始监听，处理请求，返回响应
	if err != nil {
		fmt.Println("服务器开启错误: ", err)
	}
}
