package client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Global variable of type string for get user's command
var CMD string

//Global variable of type string for save current user
var CurrentUser string

//Create User structure for save data about registration user
//We have one field -> username
type User struct {
	Username string `json:"username"` // In json file we'll use variable name -> Username
}

//Create SendUser structure for save data about Sendler
type SendData struct {
	Sender   string `json:"sender"`
	Reciever string `json:"reciever"`
	Message  string `json:"message"`
}

//Header about chat
func Intro() {
	fmt.Println("[A SIMPLE CHAT] - BE FRIENDLY!!! BE HAPPY!!!")
	fmt.Println("!!!PROMPT: help\tprint all comands")
}

//Information menu for user
func Menu() {
	fmt.Println("\treg\tregistration a new user[name]")
	fmt.Println("\tall\tshow all users in chatroom")
	fmt.Println("\tsend\tsend message to someone[whom, what]")
	fmt.Println("\tget\tget all my messages in chatroom")
	fmt.Println("\texit\tclose chatroom")
}

//Get user's command
func UserInput() {
	fmt.Scan(&CMD)
}

func RegNewUser(user *User) {

	//Get a new user name for registration
	fmt.Print("Input a new user name: ")
	fmt.Scan(&user.Username)

	//Save current user
	CurrentUser = user.Username

	//Convert struct to json type
	postBody, _ := json.Marshal(user)

	//Convert to *bytes.Buffer and initialization
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
	log.Println(string(body))
}

//Registration a new user
func Registration() {
	//Create object type of RegUser
	var user User

	//Check user, we have current user or not
	if CurrentUser == "" {
		RegNewUser(&user)
	} else {
		fmt.Println("You Signed Up! Your name:", CurrentUser)
	}

}

//Get all users
func GetAllUsers() {
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
	log.Println(string(body))

}

//Send message
func SendMessage() {
	//Create object of type SendUser
	var sendData SendData

	//Check user, we have current user or not
	if CurrentUser == "" {
		//Get the name user and save in server
		Registration()
	} else {
		//Set current user to sender
		sendData.Sender = CurrentUser
	}

	//Get Reciever
	fmt.Print("Input whom you want to send: ")
	fmt.Scan(&sendData.Reciever)

	//Get message for send
	fmt.Print("Message: ")
	inputReader := bufio.NewReader(os.Stdin)
	sendData.Message, _ = inputReader.ReadString('\n')

	//Convert struct to json type
	postBody, _ := json.Marshal(sendData)

	//Convert to *bytes.Buffer and initialization
	responseBody := bytes.NewBuffer(postBody)

	//Send request to server
	resp, err := http.Post("http://localhost:80/send", "application/json", responseBody)

	//Check request
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
	log.Println(string(body))
}

//Get user message
func GetAllUserMessages() {
	//Create object of type User structure
	var user User

	//Check user, we have current user or not
	if CurrentUser == "" {
		//Get the name user
		Registration()
	} else {
		//Set current user to sender
		user.Username = CurrentUser
	}

	//Convert struct to json type
	postBody, _ := json.Marshal(user)

	//Convert to *bytes.Buffer and initialization
	responseBody := bytes.NewBuffer(postBody)

	//Send request to server
	resp, err := http.Post("http://localhost:80/get", "application/json", responseBody)

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
	log.Println(string(body))

}
