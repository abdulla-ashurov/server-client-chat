package main

import "fmt"

//Global for get user's command
var cmd string

//Create Reg structure for save date about registration user
//We have one field -> username
type Reg struct {
	username string `json:"Username"` // In json file we'll use variable name -> Username
}

//Header about chat
func intro() {
	fmt.Println("[A SIMPLE CHAT] - BE FRIENDLY!!! BE HAPPY!!!")
	fmt.Println("!!!PROMPT: help\tprint all comands")
}

//Information menu for user
func menu() {
	fmt.Println("\treg\tregistration a new user[name]")
	fmt.Println("\tall\tshow all users in chatroom")
	fmt.Println("\tsend\tsend message to someone[whom, what]")
	fmt.Println("\tget\tget all my messages in chatroom")
	fmt.Println("\texit\tclose chatroom")
}

//Get user's command
func userInput() {
	fmt.Scan(&cmd)
}

func main() {

	//Header about chat
	intro()

	for {
		//Get user's command
		userInput()

		switch cmd {
		case "help":
			//Print Information about commands
			menu()

		case "all":
			//Print all users

		case "send":
			//Send message to some user

		case "get":
			//Get all my messages

		case "exit":
			//Close the program
			fmt.Println("Good Bye! See you soon...")
			return

		default:
			//When user input incorrect command
			fmt.Println("Command not found: ", cmd)
		}
	}

}
