package main

import (
	"fmt"
	"log"
	"net"
	"real-time-cli-chat/server"
)

func main() {

	srv := server.NewServer()

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

		srv.AddClient(conn)

		go srv.HandleConnection(conn)
	}
}
