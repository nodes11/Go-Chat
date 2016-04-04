package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
)

var name string
var reader = bufio.NewReader(os.Stdin)
var chatRoomName string

func readName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Name: ")
	message, _ := reader.ReadString('\n')
	return message[0 : len(message)-1]
}

func read() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(name)
	message, _ := reader.ReadString('\n')
	return message
}

func getMessages(c net.Conn) {
	for {
		nmessg, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Printf(string('\n') + nmessg + name)
	}
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func main() {
	clearScreen()
	name = readName()
	clearScreen()

	name += ": "

	con, err := net.Dial("tcp", "68.234.246.89:5555")
	if err != nil {
		panic(err)
	}

	chatRoomName, _ := bufio.NewReader(con).ReadString('\n')

	fmt.Printf("Chat Room: %s", chatRoomName)
	fmt.Printf("--------------------------------------------------------------------------------\n")

	for {
		go getMessages(con)

		messg := name
		messg += read()
		fmt.Fprintf(con, messg)
	}
}
