package main

import (
	"encoding/json"
	"fmt"
	"io"
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

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error openin JSON file:", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		fmt.Println(post)
	}

	// jsonData, err := ioutil.ReadAll(jsonFile)
	// if err != nil {
	// 	fmt.Println("Error reading JSON data:", err)
	// 	return
	// }
	//
	// var post Post
	// json.Unmarshal(jsonData, &post)
	// fmt.Println(post)
}
