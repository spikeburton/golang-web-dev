package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("times-visited")
	if err != nil {
		fmt.Println("First visit - initializing count")
		c = &http.Cookie{
			Name:  "times-visited",
			Value: "0",
			Path:  "/",
		}
	}

	i, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}
	c.Value = strconv.Itoa(i + 1)
	http.SetCookie(w, c)
	fmt.Fprintln(w, "Times visited: ", c.Value)

}
