package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Post exported
type Post struct {
	XMLName xml.Name `xml:"post"`
	ID      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

// Author exported
type Author struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		ID:      "1",
		Content: "Hello World!",
		Author: Author{
			ID:   "2",
			Name: "Robinson Ramirez",
		},
	}
	xmlFile, err := os.Create("post.xml")
	if err != nil {
		fmt.Println("Error creating XML file:", err)
		return
	}
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding XML to file:", err)
		return
	}

}
