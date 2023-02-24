package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

var (
	proxyAddr = "http://127.0.0.1:2003"
	port      = "2002"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//step 1 解析代理地址，并更改请求体的协议和主机
	proxy, err := url.Parse(proxyAddr)
	r.URL.Scheme = proxy.Scheme
	r.URL.Host = proxy.Host

	//step 2 请求真实服务器
	transport := http.DefaultTransport
	resp, err := transport.RoundTrip(r)
	if err != nil {
		log.Print(err)
		return
	}

	//step 3 把真实请求内容返回给上游
	for k, vv := range resp.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}

	defer resp.Body.Close()
	bufio.NewReader(resp.Body).WriteTo(w)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Start serving on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
