package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
	i := 0
	var m, u string

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			f := strings.Fields(ln)
			m = f[0]
			u = f[1]
			fmt.Printf("***METHOD: %s\n", m)
			fmt.Printf("***URI: %s\n", u)
		}
		if ln == "" {
			break
		}
		i++
	}

	mux(conn, m, u)
}

func mux(conn net.Conn, m string, u string) {
	defer conn.Close()

	var body string
	switch m {
	case "GET":
		switch u {
		case "/":
			body = fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<meta http-equiv="X-UA-Compatible" content="ie=edge">
				<title>Hello</title>
			</head>
			<body>
				<h1>HOLY COW THIS IS LOW LEVEL</h1>
				<p><strong>METHOD:</strong> %s</p>
				<p><strong>URI:</strong> %s</p>
			</body>
			</html>
			`, m, u)
		case "/apply":
			body = fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<meta http-equiv="X-UA-Compatible" content="ie=edge">
				<title>Apply</title>
			</head>
			<body>
				<form action="/apply" method="POST">
					<input type="text" name="fname" placeholder="First Name">
					<input type="submit" value="Submit Application">
				</form>
			</body>
			</html>
			`)
		default:
			body = "<h1>NOT FOUND</h1>"
		}

	case "POST":
		switch u {
		case "/apply":
			body = "<h1>POST ACCEPTED</h1>"
		}
	default:
		log.Fatalln("UNKNOWN METHOD")
	}

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	io.WriteString(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
