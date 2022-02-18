package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var address string = "http://localhost:80/"

//Create User structure for save data about registration user
//We have one field -> username
type User struct {
	Username string `json:"username"` // In json file we'll use variable name -> Username
}

//Create SendUser structure for save data about Sendler
type SendUser struct {
	Sender   string `json:"sender"`
	Reciever string `json:"reciever"`
	Message  string `json:"message"`
}

//Respond registarion user
func respondRegistration(user User) int {
	//Convert struct to JSON type
	postBody, _ := json.Marshal(user)

	//Convert to *bytes.Buffer and initialization
	responseBody := bytes.NewBuffer(postBody)

	//Integration test  for Registration User
	resp, err := http.Post(address+"reg", "application/json", responseBody)

	//Check require
	if err != nil {
		fmt.Println("Sorry, We didn't send your request. Try later...")
		return resp.StatusCode
	}
	//We close the require. Defer - works at the last
	defer resp.Body.Close()

	//Check responce body
	if err != nil {
		fmt.Println("Sorry, We didn't read responce to your reqest. Try later...")
		return resp.StatusCode
	}

	//Return status code
	return resp.StatusCode
}

//Registration function
func testNewUsers() {

	fmt.Println("Integration Test For Registration Start")

	//Create objects type of User
	users := []User{
		{"Abdulla"},
		{"Ulfat"},
		{"John"},
		{"Abdulla"},
		{""},
		{"Ulfat"},
	}

	//Answers
	answers := []int{200, 200, 200, 400, 400, 400}

	//Check Integration tests for registration
	for i := 0; i < len(users); i++ {
		if respondRegistration(users[i]) != answers[i] {
			panic("ERROR!")
		}
	}

	fmt.Println("Integration Test For Registration Succesfully Finished")
}

//Get All users Test
/*func getAllUsers() int {

	fmt.Println("Integration Test For Get All Users Start")

	//Send get require
	resp, err := http.Get(address+"all")

	//Check require
	if err != nil {
		fmt.Println("Sorry, We didn't send your request. Try later...")
		return resp.StatusCode
	}

	//We close the require. Defer - works at the last
	defer resp.Body.Close()

	//Check responce body
	if err != nil {
		return resp.StatusCode
	}

	fmt.Println("Integration Test For Get All Users Succesfully Finished")

	//Print responce
	return resp.StatusCode
}*/

//Send Message To Server
func sendMessageToServer(sendUser SendUser) int {
	//Convert struct to json type
	postBody, _ := json.Marshal(sendUser)

	//Convert to *bytes.Buffer and initialization
	responseBody := bytes.NewBuffer(postBody)

	//Send request to server
	resp, err := http.Post(address+"send", "application/json", responseBody)

	//Check request
	if err != nil {
		fmt.Println("Sorry, We didn't send your request. Try later...")
		return resp.StatusCode
	}

	//We close the require. Defer - works at the last
	defer resp.Body.Close()

	//Check responce body
	if err != nil {
		fmt.Println("Sorry, We didn't read responce to your request. Try later...")
		return resp.StatusCode
	}

	//Print responce
	return resp.StatusCode
}

//Send Function Test
func testSendMessage() {

	fmt.Println("\nIntegration Test For Send Message Start")

	sendUsers := []SendUser{
		{"Abdulla", "Ulfat", "Hi! How are you?"},
		{"Ulfat", "Abdulla", "Hi! I'm fine and you?"},
		{"Abdulla", "Ulfat", "Me too!"},
		{"Abdulla", "John", "Hi! How are you?"},
		{"John", "Abdulla", "I'm fine!"},
	}

	for i := 0; i < len(sendUsers); i++ {
		if sendMessageToServer(sendUsers[i]) != 200 {
			panic("ERROR!")
		}
	}

	//Incorrect Values
	sendUsers2 := []SendUser{
		{"Abdulla", "U", "Hi! How are you?"},
		{"Ulfat", "Aba", "Hi! I'm fine and you?"},
		{"Abdulla", "Ut", "Me too!"},
		{"Abdulla", "Jo", "Hi! How are you?"},
		{"John", "Abdulla", ""},
	}

	for i := 0; i < len(sendUsers2); i++ {
		if sendMessageToServer(sendUsers2[i]) != 400 {
			panic("ERROR!")
		}
	}

	fmt.Println("Integration Test For Send Message Succesfully Finished")
}

//Get All User Messages
func GetAllMessages(user User) int {

	//Convert struct to json type
	postBody, _ := json.Marshal(user)

	//Convert to *bytes.Buffer and initialization
	responseBody := bytes.NewBuffer(postBody)

	//Send request to server
	resp, err := http.Post(address+"get", "application/json", responseBody)

	//Check require
	if err != nil {
		fmt.Println("Sorry, We didn't send your request. Try later...")
		return resp.StatusCode
	}
	//We close the require. Defer - works at the last
	defer resp.Body.Close()

	//Check responce body
	if err != nil {
		fmt.Println("Sorry, We didn't read responce to your reqest. Try later...")
		return resp.StatusCode
	}

	//Print responce
	return resp.StatusCode
}

//Test Get All User Messages
func testGetAllMessages() {
	fmt.Println("\nIntegration Test For Get All Messages Start")

	//Create objects type of User
	users := []User{
		{"Abdulla"},
		{"Ulfat"},
		{"John"},
		{"Ks"},
		{""},
		{"fdsfds"},
	}

	//Answers
	answers := []int{200, 200, 200, 400, 400, 400}

	//Check Integration tests for registration
	for i := 0; i < len(users); i++ {
		if GetAllMessages(users[i]) != answers[i] {
			panic("ERROR!")
		}
	}

	fmt.Println("Integration Test For Get All Messages Succesfully Finished")
}

func main() {

	//Check Integration tests for Registration a new user
	testNewUsers()

	//Check Integration tests for Send Message
	testSendMessage()

	//Check Integration tests for Get All Messages
	testGetAllMessages()

}
