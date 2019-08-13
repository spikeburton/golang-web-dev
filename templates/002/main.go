package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Document</title>
	</head>
	<body>
		<h1>` + name + `</h1>
	</body>
	</html>
	`

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file")
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))
}
