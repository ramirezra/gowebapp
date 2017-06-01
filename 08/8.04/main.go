package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Post exported
type Post struct {
	ID       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

// Author exported
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Comment exported
type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func decode(filename string) (post Post, err error) {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error openin JSON file:", err)
		return
	}
	defer jsonFile.Close()
	decoder := json.NewDecoder(jsonFile)

	err = decoder.Decode(&post)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	return
}

func main() {
	post, err := decode("post.json")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(post)
}
