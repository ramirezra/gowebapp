package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Text interface {
	Fetch(id int) (err error)
	Create() (err error)
	Update() (err error)
	Delete() (err error)
}

// Post exported
type Post struct {
	DB      *sql.DB
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// Fetch a single post
func (post *Post) Fetch(id int) (err error) {
	err = post.DB.QueryRow("SELECT id, content, author FROM posts where id = $1", id).Scan(&post.ID, &post.Content, &post.Author)
	return
}

// Retrieve a single post
func Retrieve(id int) (post Post, err error) {
	post = Post{}
	err = post.DB.QueryRow("SELECT id, content, author FROM posts where id = $1", id).Scan(&post.ID, &post.Content, &post.Author)
	return
}

// Create a new post
func (post *Post) Create() (err error) {
	statement := "INSERT INTO posts (content, author) VALUES ($1,$2) RETURNING id"
	stmt, err := post.DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.ID)
	return
}

// Update a post
func (post *Post) Update() (err error) {
	_, err = post.DB.Exec("UPDATE posts SET content = $2, author =$3 WHERE id =$1", post.ID, post.Content, post.Author)
	return
}

// Delete a post
func (post *Post) Delete() (err error) {
	_, err = post.DB.Exec("DELETE FROM posts WHERE id = $1", post.ID)
	return
}
