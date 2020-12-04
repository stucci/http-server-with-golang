package main

// [Go言語で複数ドメインにも対応可能なHTTPSサーバーの作り方](https://louliz.com/ja/programming/go/create-https-server-with-golang)

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Secure Hello World!")
}

func main() {
	muxSSL := http.NewServeMux()
	muxSSL.HandleFunc("/", handler)

	tlsConfig := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		MaxVersion:               tls.VersionTLS13,
		PreferServerCipherSuites: false,
		CipherSuites: []uint16{
			tls.TLS_AES_256_GCM_SHA384,
		},
	}
	tlsConfig.Certificates = make([]tls.Certificate, 1)
	fmt.Print(tlsConfig.CipherSuites)

	var err error
	/** to create `certificate` and `private key`
	go run "C:\Go\src\crypto\tls\generate_cert.go" --rsa-bits=2048 --host=localhost
	*/
	tlsConfig.Certificates[0], err = tls.LoadX509KeyPair("./cert.pem", "./key.pem")
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig.BuildNameToCertificate()

	ServerSSL := &http.Server{
		Handler:        muxSSL,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	listener, err := tls.Listen("tcp", ":443", tlsConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = ServerSSL.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
	// note: access to https://localhost:443
}
