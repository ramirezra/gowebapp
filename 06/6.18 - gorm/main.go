package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

// Post defined
type Post struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
}

// Comment defined
type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
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

// Create POST - For CRUD app
func (post *Post) Create() {
	statement := "INSERT INTO posts (content, author) values ($1, $2) RETURNING id"
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

// Create Comment
func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post not found")
		return
	}
	err = DB.QueryRow("INSERT INTO comments (content, author, post_id) VALUES($1,$2,$3) RETURNING ID", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	return
}

// GetPost
func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	err = DB.QueryRow("SELECT id, content, author FROM posts where id=$1", id).Scan(&post.Id, &post.Content, &post.Author)

	rows, err := DB.Query("SELECT id, content, author FROM comments")
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}

// UpdatePost
func (post *Post) UpdatePost() (err error) {
	_, err = DB.Exec("UPDATE posts SET content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

// DeletePost
func (post *Post) DeletePost() (err error) {
	_, err = DB.Exec("DELETE FROM gwp WHERE id = $1", post.Id)
	return
}

// GetAllPosts
func GetAllPosts(limit int) (posts []Post, err error) {
	rows, err := DB.Query("SELECT id, content, author FROM gwp limit $1", limit)
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

func main() {
	// Define Post
	post := Post{Content: "Hello World!", Author: "Robinson Ramirez"}
	fmt.Println(post)

	// Create Post
	post.Create()
	fmt.Println(post)

	// Create Comment
	comment := Comment{Content: "Good Post!", Author: "Joe", Post: &post}
	comment.Create()

	// Get Post
	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0])

	// // Modify Post
	// readPost.Content = "Hola Mundo!"
	// readPost.Author = "Pedro Almodovar"
	// readPost.UpdatePost()
	// fmt.Println(readPost)
	// // Delete Post
	// readPost.DeletePost()
	//
	// http.HandleFunc("/", index)
	// http.ListenAndServe(":8080", nil)
}

// Handlers for web app
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
