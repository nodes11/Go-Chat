package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(">")
	message, _ := reader.ReadString('\n')
	return message
}

func getMessages(c net.Conn) {
	for {
		nmessg, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Printf(string('\n') + nmessg + ">")
	}
}

func main() {
	con, err := net.Dial("tcp", "127.0.0.1:5555")
	if err != nil {
		panic(err)
	}

	for {
		go getMessages(con)

		messg := read()
		fmt.Fprintf(con, messg)
	}
}
