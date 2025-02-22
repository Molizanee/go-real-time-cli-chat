package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"real-time-cli-chat/client"
	"strings"
	"sync"
)

type Server struct {
	clients []client.Client
	mutex sync.Mutex
}

func NewServer() *Server {
	return &Server{
		clients: make([]client.Client, 0),
	}
}

func (s *Server) AddClient(conn net.Conn) {
	// Prevent other goroutines from modifying the clients slice while creating a new client
	s.mutex.Lock()
	defer s.mutex.Unlock()

	client := client.Client{
		Conn: conn,
		Addr: conn.RemoteAddr().String(),
	}

	s.clients = append(s.clients, client)
	fmt.Printf("Accepted connection from %s\n", client.Addr)
}

func (s *Server) RemoveClient(conn net.Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, client := range s.clients {
		if client.Conn == conn {
			log.Printf("Client disconnected: %s", client.Addr)
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
			s.Broadcast(conn, fmt.Sprintf("%s has left the chat\n", client.Addr))
			break
		}
	}
}

func (s *Server) Broadcast(senderConn net.Conn, message string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	senderAddr := senderConn.RemoteAddr().String()
	response := strings.ToLower(message)	

	for _, client := range s.clients {
		if client.Addr == senderAddr {
			continue
		}
		_, err := client.Conn.Write([]byte(response))
		if err != nil {
			log.Printf("Error writing to client: %v", err)
			return
		}
	}
}

func (s *Server) HandleConnection(conn net.Conn) {

	defer func() {
		s.RemoveClient(conn)
		conn.Close()
	}()	

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
					log.Printf("Client %s disconnected", conn.RemoteAddr().String())
					return
			}
			if netErr, ok := err.(net.Error); ok {
					log.Printf("Network error from %s: %v", conn.RemoteAddr().String(), netErr)
					return
			}			
			log.Printf("Error reading from %s: %v", conn.RemoteAddr().String(), err)
			return
		}

		message := string(buffer[:n])
		s.Broadcast(conn, message)
	}
}