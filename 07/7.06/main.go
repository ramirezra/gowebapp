package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// Post exported
type Post struct {
	XMLName  xml.Name  `xml:"post"`
	ID       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	XML      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

// Author exported
type Author struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

// Comment exported
type Comment struct {
	ID      string `xml:"id,attr"`
	Content string `xml:"comntent"`
	Author  Author `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se)
				fmt.Println(comment)
			}
		}
	}

	// xmlData, err := ioutil.ReadAll(xmlFile)
	// if err != nil {
	// 	fmt.Println("Error reading XML data:", err)
	// 	return
	// }
	// var post Post
	// xml.Unmarshal(xmlData, &post)
	// fmt.Println(post)
}
