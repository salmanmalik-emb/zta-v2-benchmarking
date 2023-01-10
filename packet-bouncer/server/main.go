package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Launching server...")

	ln, err := net.Listen("tcp", ":8050")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Use bufio to read from conn
	for {
		reader := bufio.NewReader(conn)
		err := conn.SetReadDeadline(time.Now().Add(time.Minute))
		if err != nil {
			fmt.Println("error in setting read deadline", err.Error())
		}
		readBytes := make([]byte, 2048)
		noBytes, err := reader.Read(readBytes)
		if err != nil {
			fmt.Println("closing connection", err.Error())
			return
		}
		fmt.Println("Message Received len: ", noBytes)
		// Send message back to client
		_, err = conn.Write(readBytes[:noBytes])
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
