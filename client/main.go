package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func initFoo() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Failed to dial server", err)
		return
	}
	defer conn.Close()

	fmt.Println("Successfully dialed server")

	// send msg to server

	_, err = conn.Write([]byte("Hello server\n"))
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(conn)
	m, err := r.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read bytes", err)
		return
	}
	fmt.Println("Server:", m)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := []byte(scanner.Text() + "\n")
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("Failed to write to the server", err)
			return
		}

		m, err = r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Print("Server: ", m)
	}
}

func main() { initFoo() }
