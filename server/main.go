package main

import (
	"fmt"
	"net"
)

func initRealm() {
	var p string = "8080"
	listener, err := net.Listen("tcp", ":"+p)
	if err != nil {
		fmt.Println("Failed to listen on server", err)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port:", p)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Received: %s", buffer[:n])
		conn.Write([]byte("Message received:\n"))
		if err != nil {
			fmt.Println("Failed to write bytes:", err)
			return
		}
	}

}

func main() {
	initRealm()
}
