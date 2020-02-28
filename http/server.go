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
	//if strings.Contains(request,"?download"){
	//	writer := bufio.NewWriter(conn)
	//	bytes, err := ioutil.ReadFile("http/img/")
	//	_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
	//	_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
	//	_, _ = writer.WriteString("Content-Type: application/octet-stream\r\n")
	//	_, _ = writer.WriteString("Connection: Close\r\n")
	//	_, _ = writer.WriteString("\r\n")
	//	_, _ = writer.Write(bytes)
	//	err = writer.Flush()
	//	if err != nil {
	//		log.Printf("can't sent response: %v", err)
	//	}
	//	log.Printf("response on: %s", request)
	//	return
	//}
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

	if method == "GET" && (request == "/html.html" || request == "/html.html?download" ) && protocol == "HTTP/1.1" {
		content:="text/html"
		if strings.Contains(request,"?download"){
			content = "application/octet-stream"
		}
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("http/files/html.html")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: "+ content+"\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}

	if method == "GET" && (request == "/2.png" || request == "/2.png?download" ) && protocol == "HTTP/1.1" {
		content:="image/png"
		if strings.Contains(request,"?download"){
			content = "application/octet-stream"
		}
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("http/img/2.png")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: "+ content+"\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}

	if method == "GET" && (request == "/1.jpg" || request == "/1.jpg?download" )  && protocol == "HTTP/1.1" {
		content:="image/jpg"
		if strings.Contains(request,"?download"){
			content = "application/octet-stream"
		}
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("http/img/1.jpg")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: "+ content+"\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}

	if method == "GET" && (request == "/text.txt" || request=="/text.txt?download") && protocol == "HTTP/1.1" {
		content:="text/html"
		if strings.Contains(request,"?download"){
			content = "application/octet-stream"
		}
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("http/files/text.txt")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: "+ content+"\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}

	if method == "GET" && (request == "/html.pdf" || request=="/html.pdf?download") && protocol == "HTTP/1.1" {
		content:="application/pdf"
		if strings.Contains(request,"?download"){
			content = "application/octet-stream"
		}
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("http/files/html.pdf")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: "+ content+"\r\n")
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