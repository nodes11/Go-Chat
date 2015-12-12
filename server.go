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

func main() {
	listen, err := net.Listen("tcp", ":5555")
	if err != nil {
		panic(err)
	}

	con, _ := listen.Accept()

	for {
		messg, _ := bufio.NewReader(con).ReadString('\n')
		fmt.Printf(string(messg))

		messg = read()
		con.Write([]byte(messg))
	}
}
