package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	proxyAddr := "ws://127.0.0.1:9884"
	url, err := url.Parse(proxyAddr)
	if err != nil {
		log.Println(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	addr := "127.0.0.1:2002"
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
