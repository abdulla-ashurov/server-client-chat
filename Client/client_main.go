package main

import (
	client "CHAT_SERVER_CLIENT/client/Functions"
	"fmt"
)

func main() {

	//Header about chat
	client.Intro()

	for {
		//Get user's command
		client.UserInput()

		switch client.CMD {
		case "help":
			//Print Information about commands
			client.Menu()

		case "reg":
			//Registration user
			client.Reg()
		case "all":
			//Print all users
			client.GetAll()
		case "send":
			//Send message to some user
			client.Send()

		case "get":
			//Get all my messages
			client.GetMsg()
		case "exit":
			//Close the program
			fmt.Println("Good Bye! See you soon...")
			return

		default:
			//When user input incorrect command
			fmt.Println("Command not found: ", client.CMD)
		}
	}

}
