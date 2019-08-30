package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

var tpl *template.Template

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		URL         *url.URL
		Header      http.Header
		Submissions url.Values
	}{
		req.Method,
		req.URL,
		req.Header,
		req.Form,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseFiles("./index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
