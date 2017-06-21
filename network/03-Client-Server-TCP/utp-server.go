package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buffer := make([]byte, 1024)

	_, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buffer))

	conn.Write([]byte("Messagem recebida!"))
	conn.Close()
}