package main

import (
	"io"
	"net/http"
)

func main() {
	//http://localhost:8000/go
	http.HandleFunc("/go", healthz)
	http.HandleFunc("/philosophy", philosophy)
	// addr：监听的地址
	// handler：回调函数
	http.ListenAndServe("localhost:8000", nil)
}
func philosophy(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "The first philosophy is metaphysics")
}

// handler函数
func healthz(w http.ResponseWriter, r *http.Request) {
	// 回复
	io.WriteString(w, "200")
	//w.Write([]byte("200"))
}
