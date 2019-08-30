package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			break
		}
	}

	// body := `
	// <!DOCTYPE html>
	// <html lang="en">
	// <head>
	// 	<meta charset="UTF-8">
	// 	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	// 	<meta http-equiv="X-UA-Compatible" content="ie=edge">
	// 	<title>Hello</title>
	// </head>
	// <body>
	// 	<h1>Hello, World!</h1>
	// </body>
	// </html>
	// `

	// io.WriteString(conn, "HTTP/1.1 GET")
	// io.WriteString(conn, "Content-Type: text/html")
	// fmt.Fprintf(conn, "Content-Length: %d\n", len(body))
	// fmt.Fprintf(conn, "\r\n")
	// fmt.Fprintf(conn, body)
	io.WriteString(conn, "HELLO WORLD")
}
