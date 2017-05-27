package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

// Driver exported
var Driver = "postgres"

// Account exported
var Account = "gwp"

// Pword exported
var Pword = "gwp"

// Datastore exported
var Datastore = "localhost/gwp"

// DB exported
var DB *sql.DB

func init() {
	var err error

	DB, err = sql.Open(Driver, Driver+"://"+Account+":"+Pword+"@"+Datastore+"?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
