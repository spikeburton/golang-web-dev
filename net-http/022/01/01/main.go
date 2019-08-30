package main

import (
	"io"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "INDEX")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog dog dog")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Spike Burton")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
