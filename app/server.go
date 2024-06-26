package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	con, err := l.Accept()

	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer con.Close()

	for {
		var buff []byte

		_, err = con.Read(buff)
		if err != nil {

			fmt.Println("Error reading message: ", err.Error())

			os.Exit(1)

		}

		_, err = con.Write([]byte("+PONG\r\n"))

		if err != nil {

			fmt.Println("Error writing response message: ", err.Error())
		}
	}
}
