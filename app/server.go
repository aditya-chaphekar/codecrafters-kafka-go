package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	buff := make([]byte, 1024)
	_, err = conn.Read(buff)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		os.Exit(1)
	}
	conn.Write([]byte{0, 0, 0, 0, 0, 0, 0, 7})
}
