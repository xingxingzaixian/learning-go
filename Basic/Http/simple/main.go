package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request) {
	upath := "error handler"
	w.WriteHeader(500)
	io.WriteString(w, upath)
}

func TimeoutHandler(w http.ResponseWriter, req *http.Request) {
	time.Sleep(6 * time.Second)
	upath := "timeout handler"
	w.WriteHeader(200)
	io.WriteString(w, upath)
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		upath := fmt.Sprintf("%s\n", request.URL.Path)
		realIP := fmt.Sprintf("RemoteAddr=%s,X-Forwarded-For=%v,X-Real-Ip=%v\n", request.RemoteAddr, request.Header.Get("X-Forwarded-For"), request.Header.Get("X-Real-Ip"))
		header := fmt.Sprintf("headers =%v\n", request.Header)
		io.WriteString(writer, upath)
		io.WriteString(writer, realIP)
		io.WriteString(writer, header)
	})
	http.HandleFunc("/error", ErrorHandler)
	http.HandleFunc("/timeout", TimeoutHandler)
	fmt.Println("Start server on port 2003")
	http.ListenAndServe(":2003", nil)
}
