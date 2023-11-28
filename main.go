package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"turtlequeue/commands"
	"turtlequeue/parser"
	"turtlequeue/server"
)

func decodeRespString(reqStr string, conn net.Conn) string {
	if strings.ToLower(reqStr) == "ping" {
		return "PONG"
	} else if strings.ToLower(reqStr) == "pong" {
		return "PING"
	} else if strings.HasPrefix(strings.ToLower(reqStr), "connect") {
		return "+OK"
	} else if strings.HasPrefix(strings.ToLower(reqStr), "sub") {
		var commandArgs map[string]string = parser.DecodeSubCommand(reqStr)
		return commands.SubscribeClient(commandArgs, conn)
	} else if strings.HasPrefix(strings.ToLower(reqStr), "pub") {
		var commandArgList map[string]string = parser.DecodeSubCommand(reqStr)
		return commands.PublishClient(commandArgList, conn)
	}
	return ""
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte(fmt.Sprintf("%s\r\n", commands.Intro(conn))))

	// Use a dynamic buffer to accumulate the data
	var buffer []byte
	tempBuffer := make([]byte, 1024*20)

	for {

		// Read a chunk of data from the client
		bytesRead, err := conn.Read(tempBuffer)
		if err != nil {
			fmt.Println("Error occurred", err)
			if err == io.EOF {
				// Connection closed by the client
				break
			}
			fmt.Println("Error reading from client:", err)
			return
		}

		// Append the received chunk to the buffer
		buffer = append(buffer, tempBuffer[:bytesRead]...)

		// Check if the message is complete (ending with "\r\n")
		if len(buffer) >= 2 && buffer[len(buffer)-2] == '\r' && buffer[len(buffer)-1] == '\n' {
			// Process the complete message (excluding "\r\n")
			message := string(buffer[:len(buffer)-2])
			decodedStr := decodeRespString(message, conn)
			//fmt.Println("Returning: ", decodedStr)
			//Echo the message back to the client
			conn.Write([]byte(fmt.Sprintf("%s\r\n", decodedStr)))

			// Clear the buffer for the next message
			buffer = nil
		}
	}

}

func main() {
	listener := server.CreateTCPConnection()

	for {
		// Accept incoming client connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle each client connection in a separate goroutine
		go handleClient(conn)
	}
}
