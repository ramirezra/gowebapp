package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/hello/:name", hello)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}
	server.ListenAndServe()
}
