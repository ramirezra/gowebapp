package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	// var t *template.Template
	if rand.Intn(10) > 5 {
		t := template.Must(template.ParseFiles("layout.gohtml", "red.gohtml"))
	} else {
		t := template.Must(template.ParseFiles("layout.gohtml", "blue.gohtml"))
	}
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
