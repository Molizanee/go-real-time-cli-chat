# Go Real-Time CLI Chat

A simple, lightweight, real-time chat application built with Go that runs in the command line. This project demonstrates the fundamentals of network programming in Go, including TCP connections, concurrency with goroutines, and basic client-server architecture.

## Overview

This chat application consists of a server that handles multiple client connections, allowing connected clients to broadcast messages to all other connected clients in real-time. The application implements:

- TCP socket connections for reliable communication
- Concurrent client handling with goroutines
- Thread-safe client management using mutex locks
- Simple broadcast messaging system

## Features

- Real-time messaging between multiple clients
- Automatic notification when users join or leave the chat
- Concurrent client handling (the server can manage multiple connections simultaneously)
- Clean disconnection handling and resource cleanup

## Project Structure

```
go-real-time-cli-chat/
├── client/
│   └── client.go       # Client structure definition
├── server/
│   └── server.go       # Server implementation with connection handling
├── main.go             # Application entry point
├── go.mod              # Go module definition
├── .gitignore          # Git ignore file
└── README.md           # This documentation
```

## How It Works

1. The server listens for incoming TCP connections on port 8080
2. When a client connects, the server:
   - Creates a new client instance
   - Adds it to the list of active clients
   - Spawns a dedicated goroutine to handle messages from that client
3. When a client sends a message, the server broadcasts it to all other connected clients
4. When a client disconnects, the server removes it from the client list and notifies other clients

## Technical Details

### Server

The server component manages client connections and message broadcasting:

- `Server` struct: Maintains a list of connected clients and a mutex for thread-safe operations
- `AddClient`: Safely adds a new client to the server's client list
- `RemoveClient`: Removes a disconnected client and notifies other clients
- `Broadcast`: Sends messages from one client to all other connected clients
- `HandleConnection`: Processes incoming data from each client connection

### Client

The client component is a simple structure representing a connected user:

- `Client` struct: Contains the network connection and client address information

### Main

The main package initializes the server and sets up the TCP listener to accept new connections.

## Getting Started

### Prerequisites

- Go 1.22.2 or later
- Git (for cloning the repository)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/go-real-time-cli-chat.git
   cd go-real-time-cli-chat
   ```

2. Build the application:
   ```bash
   go build
   ```

### Running the Server

Execute the compiled binary to start the server:

```bash
./real-time-cli-chat
```

The server will start listening on port 8080.

### Connecting as a Client

You can connect to the chat server using telnet or netcat:

```bash
# Using netcat
nc localhost 8080

# Using telnet
telnet localhost 8080
```

Once connected, simply type messages and press Enter to send them to all other connected clients.

## Testing

Currently, the application doesn't include formal tests. To test manually:

1. Start the server
2. Connect multiple clients using different terminal windows
3. Send messages from each client and verify they are received by other clients
4. Test disconnection by closing a client connection

## Future Improvements

Some potential enhancements for this project:

1. Add username support for better identification of chat participants
2. Implement private messaging between specific clients
3. Add a dedicated client application with a better user interface
4. Implement chat rooms or channels for group conversations
5. Add authentication and security features
6. Write comprehensive test suite for all components
7. Add command support (e.g., `/help`, `/list`, `/exit`)
8. Implement logging and message history

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- This project was created for learning purposes to demonstrate network programming concepts in Go
- Inspired by classic IRC and other command-line chat applications
