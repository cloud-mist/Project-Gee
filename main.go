package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 两个路由，根据不同的Path，调用不同的处理函数
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":9999", nil)) // 启动web服务器
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

type Engine struct{}

// 实现Handler接口
func (engine Engine) ServeHttp(w http.ResponseWriter, req *http.Request) {
	// request 包含了http请求的所有信息，
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
