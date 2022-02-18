package main

import (
	client "Chat_Server_Client/client/Functions"
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
			client.Registration()
		case "all":
			//Print all users
			client.GetAllUsers()
		case "send":
			//Send message to some user
			client.SendMessage()

		case "get":
			//Get all my messages
			client.GetAllUserMessages()
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
