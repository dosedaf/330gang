package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')

		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		response := make([]byte, 1024)
		n, err := conn.Read(response)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		fmt.Println(string(response[:n]))
	}
}
