package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	fmt.Println(err)
}
