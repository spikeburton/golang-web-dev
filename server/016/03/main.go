package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	reqest(conn)
}

func reqest(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	i := 0
	var m string
	var u string

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			f := strings.Fields(ln)
			m = f[0]
			u = f[1]
			fmt.Println("***METHOD", m)
			fmt.Println("***URI", u)
		}
		if ln == "" {
			break
		}
		i++
	}

	response(conn, m, u)
}

func response(conn net.Conn, m string, u string) {
	// body := fmt.Sprintf(`

	// `, m, u)
	var d = struct {
		Method string
		URI    string
	}{
		m,
		u,
	}
	var b bytes.Buffer
	tpl.Execute(&b, d)
	body := b.String()

	switch m {
	case "GET":
		if u != "/" {
			body = `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<meta http-equiv="X-UA-Compatible" content="ie=edge">
				<title>Document</title>
			</head>
			<body>
				404 NOT FOUND
			</body>
			</html>`
		}

		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprintf(conn, "Content-Type: text/html\r\n")
		fmt.Fprintf(conn, "\r\n")
		fmt.Fprintf(conn, body)
	case "POST":
		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Type: text/html\r\n")
		fmt.Fprintf(conn, "\r\n")
		fmt.Fprintf(conn, "HELLO THERE")
	default:
		fmt.Println("ERROR: CANNOT HANDLE REQUEST", m)
	}

}
