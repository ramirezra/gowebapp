package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Post exported
type Post struct {
	XMLName xml.Name `xml:"post"`
	ID      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	XML     string   `xml:",innerxml"`
}

// Author exported
type Author struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}
	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}
