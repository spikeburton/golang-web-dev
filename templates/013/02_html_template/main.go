package main

import (
	"html/template"
	"log"
	"os"
)

type page struct {
	Title, Header, Body string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p := page{
		"Site Title",
		"WOW A Header",
		"<script>alert('YO!');</script>",
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
}
