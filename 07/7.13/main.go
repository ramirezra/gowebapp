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

func main() {
	post := Post{
		ID:      1,
		Content: "Hello World!",
		Author: Author{
			ID:   2,
			Name: "Robinson Ramirez",
		},
		Comments: []Comment{
			Comment{
				ID:      3,
				Content: "Have a great day!",
				Author:  "Adam",
			},
			Comment{
				ID:      4,
				Content: "How are you todya?",
				Author:  "Betty",
			},
		},
	}

	jsonFile, err := os.Create("post.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "\t")
	// 	encoder.Indent("", "\t")

	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
	// output, err := json.MarshalIndent(&post, "", "\t\t")
	// if err != nil {
	// 	fmt.Println("Error marshalling to JSON:", err)
	// 	return
	// }
	// err = ioutil.WriteFile("post.json", output, 0644)
	// if err != nil {
	// 	fmt.Println("Error writing JSON to file:", err)
	// 	return
	// }
}
