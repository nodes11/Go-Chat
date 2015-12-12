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
	con, err := net.Dial("tcp", "127.0.0.1:5555")
	if err != nil {
		panic(err)
	}

	for {
		messg := read()

		fmt.Fprintf(con, messg)

		nmessg, _ := bufio.NewReader(con).ReadString('\n')

		fmt.Printf(nmessg)
	}
}
