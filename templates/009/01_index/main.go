package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
var fm = template.FuncMap{}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	xs := []string{"zero", "one", "two", "three", "four", "five"}
	data := struct {
		Words []string
		Lname string
	}{
		xs,
		"Burton",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
