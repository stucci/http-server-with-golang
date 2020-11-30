package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Secure Hello World without crypt/tls!")
}

func main() {
	muxSSL := http.NewServeMux()
	muxSSL.HandleFunc("/", handler)

	ServerHTTP := &http.Server{
		Addr:           ":443",
		Handler:        muxSSL,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := ServerHTTP.ListenAndServeTLS("./cert.pem", "./key.pem")
	if err != nil {
		fmt.Println(err)
	}

	// note: access to https://localhost:443
}
