package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

// Post defined
type Post struct {
	Id      int
	Content string
	Author  string
}

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

// Create - For CRUD app
func (post *Post) Create() {
	statement := "INSERT INTO gwp (content, author) values ($1, $2) RETURNING id"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

// Get post
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = DB.QueryRow("SELECT id, content, author FROM gwp where id=$1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Robinson Ramirez"}
	fmt.Println(post)
	post.Create()
	fmt.Println(post)

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

// Handlers for web app
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
