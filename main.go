package main

import (
	"fmt"
	"io"
	"os"
	"log"
	"strings"
)

func main() {
	name := "Robinson Ramirez"

	str := fmt.Sprint(`
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset-"UTF-8">
<title>GOLANG WEb Application</title>
<body>
<h1>` + name + `
</h1>
</body>
</html>`)

	newfile, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer newfile.Close()

	io.Copy(newfile, strings.NewReader(str))
}