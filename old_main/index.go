package main

// [Go言語でシンプルで簡単なHTTPサーバーの作り方](https://louliz.com/ja/programming/go/create-a-http-server-with-golang)

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func handlerCat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "にゃーん")
}

func main() {
	muxHTTP := http.NewServeMux()
	muxHTTP.HandleFunc("/", handler)
	muxHTTP.HandleFunc("/cat", handlerCat)

	ServerHTTP := &http.Server{
		Addr:           ":8080",
		Handler:        muxHTTP,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := ServerHTTP.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
