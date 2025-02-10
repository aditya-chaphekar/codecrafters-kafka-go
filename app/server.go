package main

import (
	"fmt"
	"net"
	"os"
)

type MessageSize struct {
	StartOffset int
	EndOffset   int
	Value       []byte
}

type CorrelationId struct {
	StartOffset int
	EndOffset   int
	Value       []byte
}

type ApiKey struct {
	StartOffset int
	EndOffset   int
	Value       []byte
}

type ApiVersion struct {
	StartOffset int
	EndOffset   int
	Value       []byte
}

type Request struct {
	MessageSize
	CorrelationId
	ApiKey
	ApiVersion
}

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

	req := new(Request)

	req.MessageSize.StartOffset = 0
	req.MessageSize.EndOffset = 4
	req.MessageSize.Value = buff[req.MessageSize.StartOffset:req.MessageSize.EndOffset]

	req.ApiKey.StartOffset = 4
	req.ApiKey.EndOffset = 6
	req.ApiKey.Value = buff[req.ApiKey.StartOffset:req.ApiKey.EndOffset]

	req.ApiVersion.StartOffset = 6
	req.ApiVersion.EndOffset = 8
	req.ApiVersion.Value = buff[req.ApiVersion.StartOffset:req.ApiVersion.EndOffset]

	req.CorrelationId.StartOffset = 8
	req.CorrelationId.EndOffset = 12
	req.CorrelationId.Value = buff[req.CorrelationId.StartOffset:req.CorrelationId.EndOffset]

	if err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
		os.Exit(1)
	}
	resp := make([]byte, 8)
	copy(resp, req.MessageSize.Value)
	copy(resp[4:], req.CorrelationId.Value)

	// add 2 bytes in response with value as 35
	errorByte := make([]byte, 2)
	errorByte[0] = 0
	errorByte[1] = 35
	resp = append(resp, errorByte...)

	conn.Write(resp)

}
