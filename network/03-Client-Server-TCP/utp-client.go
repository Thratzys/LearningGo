package main

import (
	"bufio"
	"os"
	"fmt"
	"net"
)

func main() {
	tcpAddress, err := net.ResolveTCPAddr("tcp", "localhost:3000")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddress)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Write your text: ")
	data, _ := reader.ReadString('\n')

	_, err = conn.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	fmt.Println("Wrote:", data)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		panic(err)
	}

	fmt.Println("Reply:", string(reply))

	conn.Close()
}