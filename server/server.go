package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	count    int64 = 0
	messages       = make([]string, 0)
	mutex          = &sync.Mutex{}
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	listener, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()

	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	for {
		netData, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		message := strings.TrimSpace(string(netData))

	}
}
