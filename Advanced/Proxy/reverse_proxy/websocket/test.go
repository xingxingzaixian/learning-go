package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	u, err := url.Parse("ws://192.168.10.213:9884")
	if err != nil {
		log.Fatalln(err)
	}

	err = http.ListenAndServe(":2002", NewProxy(u))
	if err != nil {
		log.Fatalln(err)
	}
}
