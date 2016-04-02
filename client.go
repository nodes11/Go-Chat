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

func readName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Name: ")
	message, _ := reader.ReadString('\n')
	return message[0: len(message) - 1]
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

func clearScreen(){
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	fmt.Printf("Chat Room:\n")
	fmt.Printf("--------------------------------------------------------------------------------\n")
}

func main() {
	clearScreen();
	name = readName()
	clearScreen()

	name += ": "

	con, err := net.Dial("tcp", "127.0.0.1:5555")
	if err != nil {
		panic(err)
	}

	for {
		go getMessages(con)

		messg := name
		messg += read()
		fmt.Fprintf(con, messg)
	}
}
