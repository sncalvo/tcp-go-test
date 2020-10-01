package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
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
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()
	defer atomic.AddInt64(&count, -1)

	atomic.AddInt64(&count, 1)
	connection.Write([]byte("Hello\n"))
	fmt.Print(".")
	for {
		netData, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		message := strings.TrimSpace(string(netData))
		if message == "STOP" {
			return
		}

		if message != "MESSAGES" {
			mutex.Lock()

			messages = append(messages, message)

			mutex.Unlock()
		} else {
			connection.Write([]byte(strings.Join(messages, ", ")))
		}
		fmt.Println(message)
		counter := strconv.FormatInt(count, 10) + "\n"
		connection.Write([]byte("\n" + string(counter)))
	}
}
