package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	serverAddr := "127.0.0.1:4222" // Replace with your server address and port
	messageText := "Hello World"   // Replace this with your message text
	dynamicData := []string{"SS", "AA", "BB"}

	// Loop to send messages with different dynamic_data values
	for _, data := range dynamicData {
		message := fmt.Sprintf("PUB %s 11\r\n%s\r\n", data, messageText)
		fmt.Println(message)
		sendMessage(serverAddr, message)
		time.Sleep(1 * time.Second) // Adjust this delay as needed for load testing
	}
}

func sendMessage(addr, message string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	_, err = fmt.Fprintf(conn, message)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
	fmt.Println("Message sent:", message)
}
