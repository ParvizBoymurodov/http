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
	const address ="0.0.0.0:9999"
	log.Print("server starting")
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("can't listen on %s: %v",address,err)
	}
	defer listener.Close()
	for  {
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
	reader:=bufio.NewReader(conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("can't read %v",err)
	}
	log.Printf(requestLine)
	parts := strings.Split(strings.TrimSpace(requestLine), " ")
	if len(parts) != 3 {
		return
	}

	method, request, protocol := parts[0], parts[1], parts[2]
	if method == "GET" && request == "/" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("http/files/file.html")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: text/html\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}

	if method == "GET" && request == "/html.html" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("http/files/html.html")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: text/html\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}

	if method == "GET" && request == "/2.png" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("http/img/2.png")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: image/png\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}

	if method == "GET" && request == "/1.jpg" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("http/img/1.jpg")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: image/jpg\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}

	if method == "GET" && request == "/text.txt" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("http/files/text.txt")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: text/html\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}

	if method == "GET" && request == "/html.pdf" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("http/files/html.pdf")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: application/pdf\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}
}