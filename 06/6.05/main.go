package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB
var Driver = "postgres"
var Account = "gwp"
var Pword = "gwp"
var Datastore = "gwp"

// Connect ot database
func init() {
	var err error
	// Db, err = sql.Open(Driver, Driver+"://"+Account+":"+Pword+"@"+Datastore+"?sslmode=disable")
	Db, err = sql.Open("postgres", "user=gwp dbname=gqp password=gwp sslmod=disable")
	if err != nil {
		panic(err)
	}
	Db.Ping()
	fmt.Println("You are connected to the database")
}

// Posts exported
func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("SELECT id, content, author FROM posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// GetPost gets a single post
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("SELECT id, content, author FROM posts WHERE id=$1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// Create a post
func (post *Post) Create() (err error) {
	statement := "INSERT INTO posts (content, author) VALUES ($1, $2)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

// Update a post
func (post *Post) Update() (err error) {
	_, err = Db.Exec("UPDATE posts SET content = $2, author =$3 where id = $1", post.Id, post.Content, post.Author)
	return
}

// Delete a post
func (post *Post) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM posts WHERE id = $1", post.Id)
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Robinson Ramirez"}

	fmt.Println(post)
	post.Create()
	fmt.Println(post)

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()

	// Get all posts.
	posts, _ := Posts(10)
	fmt.Println(posts)
	// readPost.Delete()

}
