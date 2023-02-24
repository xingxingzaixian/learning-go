package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type Pxy struct{}

func (p *Pxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)
	transport := http.DefaultTransport

	// step 1, 复制请求
	outReq := new(http.Request)
	*outReq = *req
	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}

	// step 2, 把请求发送给真实服务器
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	// step 3, 把下游请求内容返回给上游
	for key, value := range res.Header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}

	// 把服务器返回的内容，发送给客户端
	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
	res.Body.Close()
}

func main() {
	fmt.Println("Serve on :8082")
	http.Handle("/", &Pxy{})
	http.ListenAndServe("0.0.0.0:8082", nil)
}
