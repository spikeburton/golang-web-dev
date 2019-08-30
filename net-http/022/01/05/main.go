package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./tpl.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(w, "tpl.gohtml", req.Form)
	if err != nil {
		log.Fatalln(err)
	}
}
