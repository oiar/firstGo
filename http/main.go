package main

import (
	"fmt"
	"net/http"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
	var c *http.Cookie = &http.Cookie{
		Name:  "que",
		Value: "zhuzhu",
	}
	r.AddCookie(c)
	// fmt.Fprintln(w, r.UserAgent(), r.Cookies())
	fmt.Fprintf(w, "这是一个开始")
}

func main() {
	http.HandleFunc("/", myWeb) //注册路由函数

	fmt.Println("服务器即将开启： http://127.0.0.1:5001")

	err := http.ListenAndServe(":5001", nil) //开始监听，处理请求，返回响应
	if err != nil {
		fmt.Println("服务器开启错误: ", err)
	}
}
