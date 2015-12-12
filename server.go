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
		messg, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Printf(string('\n') + string(messg) + ">")
	}
}

func main() {
	listen, err := net.Listen("tcp", ":5555")
	if err != nil {
		panic(err)
	}

	con, _ := listen.Accept()

	for {
		go getMessages(con)

		messg := read()
		con.Write([]byte(messg))
	}
}
