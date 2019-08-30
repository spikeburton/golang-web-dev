package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HOTDOG")
}

func main() {
	var d hotdog
	d = 5
	http.ListenAndServe(":8080", d)
}
