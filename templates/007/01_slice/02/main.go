package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	stuff := []string{"one", "two", "three"}
	err := tpl.Execute(os.Stdout, stuff)
	if err != nil {
		log.Fatalln(err)
	}
}
