package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	connection, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		message, _ := bufio.NewReader(connection).ReadString('\n')
		fmt.Print("->: " + message)

		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(connection, text+"\n")
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
