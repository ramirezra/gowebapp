package main

import (
	"database/sql"
	"fmt"

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

// Retrieve a single post
func Retrieve(id int) (post Post, err error) {
	post = Post{}
	err = DB.QueryRow("SELECT id, content, author FROM posts where id = $1", id).Scan(&post.ID, &post.Content, &post.Author)
	return
}

// Create a new post
func (post *Post) Create() (err error) {
	statement := "INSERT INTO posts (content, author) VALUES ($1,$2) RETURNING id"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.ID)
	return
}

// Update a post
func (post *Post) Update() (err error) {
	_, err = DB.Exec("UPDATE posts SET content = $2, author =$3 WHERE id =$1", post.ID, post.Content, post.Author)
	return
}

// Delete a post
func (post *Post) Delete() (err error) {
	_, err = DB.Exec("DELETE FROM posts WHERE id = $1", post.ID)
	return
}
