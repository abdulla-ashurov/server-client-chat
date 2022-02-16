package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Global for get user's command
var cmd string

//Create Reg structure for save data about registration user
//We have one field -> username
type RegUser struct {
	Username string `json:"username"` // In json file we'll use variable name -> Username
}

//Create AllUsers structure for save

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

//Registration a new user
func reg() {
	//Create object type of RegUser
	var user RegUser

	//Get a new user name for registration
	fmt.Print("Input a new user name: ")
	fmt.Scan(&user.Username)

	//Convert struct to json type
	postBody, _ := json.Marshal(user)

	//Convert to bytes and initialization
	responseBody := bytes.NewBuffer(postBody)

	//Send request to server
	resp, err := http.Post("http://localhost:80/reg", "application/json", responseBody)

	//Check require
	if err != nil {
		fmt.Println("Sorry, We didn't send your request. Try later...")
		return
	}
	//We close the require. Defer - works at the last
	defer resp.Body.Close()

	//Read the responce body
	body, err := ioutil.ReadAll(resp.Body)

	//Check responce body
	if err != nil {
		fmt.Println("Sorry, We didn't read responce to your reqest. Try later...")
		return
	}

	//Print responce
	log.Print(string(body))
}

//Get all users
func getAll() {
	//Send get require
	resp, err := http.Get("http://localhost:80/all")

	//Check require
	if err != nil {
		fmt.Println("Sorry, We didn't send your request. Try later...")
		return
	}

	//We close the require. Defer - works at the last
	defer resp.Body.Close()

	//Read the responce body
	body, err := ioutil.ReadAll(resp.Body)

	//Check responce body
	if err != nil {
		fmt.Println("Sorry, We didn't read responce to your request. Try later...")
	}

	//Print responce
	log.Print(string(body))

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

		case "reg":
			//Registration user
			reg()
		case "all":
			//Print all users
			getAll()
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
