package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error creating listener: %v", err)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error accepting connection: %v", err)
			continue
		}

		fmt.Println("Accepted connection from", conn.RemoteAddr())

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Fatalf("Error reading data: %v", err)
			return
		}

		message := string(buffer[:n])
		fmt.Println("Received message:", message)

		response := strings.ToLower(message)

		_, err = conn.Write([]byte(response))
		if err != nil {
			log.Fatalf("Error writing response: %v", err)
			return
		}
	}
}