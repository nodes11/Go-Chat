package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
)

var chatRoomName string
var messg string
var people []net.Conn
var reader = bufio.NewReader(os.Stdin)

//Reads in user input... Kinda useless for the server
func read() string {
	reader = bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')
	return message
}

//Listens to clients and sends out the incoming messages
func getMessages(c net.Conn) {
	for {
		//Read a message from the client
		messg,err := bufio.NewReader(c).ReadString('\n')

		if err != nil {
			break;
		}

		for i:=0; i < len(people); i++{
			if (c != people[i]){
				people[i].Write([]byte(messg))
			}
		}

	}
}

//Clear the screen
func clearScreen(){
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

//Handles the connection process
func main() {
	//Get user name and clear screen
	clearScreen()
	fmt.Printf("Chat Name: ")
	chatRoomName = read()
	clearScreen()
	fmt.Printf("Chat Room: %s", chatRoomName)
	fmt.Printf("--------------------------------------------------------------------------------\n")

	//Get a connection from a client
	listen, err := net.Listen("tcp", ":5555")
	if err != nil {
		panic(err)
	}

	//Chat
	for {
		con, _ := listen.Accept()
		if con != nil {
				fmt.Printf("New person joined!\n")
		}

		//Add them to the list
		people = append(people, con)

		con.Write([]byte(chatRoomName))

		//Start listening for messages
		go getMessages(con)

		//con = nil
	}
}
