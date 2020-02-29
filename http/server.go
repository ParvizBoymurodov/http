package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("http/log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Can't close file; %v", err)
		}
	}()
	log.SetOutput(file)
    host:="0.0.0.0"
	port, ok := os.LookupEnv("PORT")
	if !ok{
		port="9999"
	}
	err= start(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatal(err)
	}
	log.Print("server starting")
}
func start(address string)(err error)  {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("can't listen on %s: %v", address, err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("can't accept conection %v", conn)
			continue
		}
		handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("can't close connection:%v", err)
		}
	}()
	reader := bufio.NewReader(conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("can't read %v", err)
	}
	log.Printf(requestLine)
	parts := strings.Split(strings.TrimSpace(requestLine), " ")
	if len(parts) != 3 {
		return
	}

	method, request, protocol := parts[0], parts[1], parts[2]

	if method == "GET" && protocol == "HTTP/1.1" {
		switch request {
		case "/":
			WriteBuffer(conn, "http/files/file.html", request, typeHtml)
		case "/html.html":
			WriteBuffer(conn, htmlWay, request, typeHtml)
		case "/html.html?download":
			WriteBuffer(conn, htmlWay, request, typeHtml)
		case "/2.png":
			WriteBuffer(conn, pngWay, request, png)
		case "/2.png?download":
			WriteBuffer(conn, pngWay, request, png)
		case "/1.jpg":
			WriteBuffer(conn, jpgWay, request, jpg)
		case "/1.jpg?download":
			WriteBuffer(conn, jpgWay, request, jpg)
		case "/text.txt":
			WriteBuffer(conn, txtWay, request, typeHtml)
		case "/text.txt?download":
			WriteBuffer(conn, txtWay, request, typeHtml)
		case "/html.pdf":
			WriteBuffer(conn, pdfWay, request, pdf)
		case "/html.pdf?download":
			WriteBuffer(conn, pdfWay, request, pdf)
		default:
			return
		}

	}
}

func WriteBuffer(conn net.Conn, fileName, request, contentType string) {
	if strings.Contains(request, "?download") {
		contentType = "application/octet-stream"
	}
	writer := bufio.NewWriter(conn)
	bytes, err := ioutil.ReadFile(fileName)
	_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
	_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
	_, _ = writer.WriteString("Content-Type: " + contentType + "\r\n")
	_, _ = writer.WriteString("Connection: Close\r\n")
	_, _ = writer.WriteString("\r\n")
	_, _ = writer.Write(bytes)
	err = writer.Flush()
	if err != nil {
		log.Printf("can't sent response: %v", err)
	}
	log.Printf("response on: %s", request)
}
