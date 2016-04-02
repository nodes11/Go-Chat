package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
)

var name string
var messg string

var reader = bufio.NewReader(os.Stdin)

func readName() string {
	reader = bufio.NewReader(os.Stdin)
	fmt.Printf("Name: ")
	message, _ := reader.ReadString('\n')
	return message[0:len(message) - 1]
}

func read() string {
	reader = bufio.NewReader(os.Stdin)
	fmt.Printf(name)
	message, _ := reader.ReadString('\n')
	return message
}

func getMessages(c net.Conn) {
	for {
		messg, err := bufio.NewReader(c).ReadString('\n')
		if err != nil{
			break;
		}

		//scanner := bufio.NewScanner(os.Stdin)
		//tempText := scanner.Text()

		fmt.Printf(string('\n') + string(messg) + name)
	}


		fmt.Printf("Connection Closed")
}

func clearScreen(){
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	fmt.Printf("Chat Room:\n")
	fmt.Printf("--------------------------------------------------------------------------------\n")
}

func main() {
	//Get user name
	clearScreen();
	name = readName()
	clearScreen()



	name += ": "

	//Get a connection from a client
	listen, err := net.Listen("tcp", ":5555")
	if err != nil {
		panic(err)
	}

	con, _ := listen.Accept()

	//CHAT!!!
	for {

		go getMessages(con)

		messg = name

		messg += string(read())

		con.Write([]byte(messg))
	}
}
