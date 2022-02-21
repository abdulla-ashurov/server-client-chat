package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var address string = "http://localhost:80/"
var appJson string = "application/json"

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

type GetUserMessage struct {
	SendMsg []SendUser
	Correct []string
}

func respondToRegistration(user User) int {

	//Conver struct to JSON type
	postBody, _ := json.Marshal(user)

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(address+"reg", appJson, responseBody)

	//Check require
	if err != nil {
		fmt.Println("Sorry, We didn't send your request. Try later...")
		return resp.StatusCode
	}

	//We close the require. Defer - works at the last
	defer resp.Body.Close()

	//Check response body
	if err != nil {
		fmt.Println("Sorry, We didn't read response to your request. Try later...")
		return resp.StatusCode
	}

	return resp.StatusCode
}

//Test Registration for Unregistrate Users
func TestIntegrationRegistationUnregistrateUsers(t *testing.T) {
	users := []User{
		{"Abdulla"},
		{"Ulfat"},
		{"Roma"},
		{"Dasha"},
		{"Sasha"},
	}

	//Check Integration Test for registration
	for i := 0; i < len(users); i++ {
		assert.EqualValues(t, respondToRegistration(users[i]), http.StatusOK, "ERROR!")
	}
}

//Test Registation for Registrate Users
func TestIntegrationRegistationRegistrateUsers(t *testing.T) {
	users := []User{
		{"John"},
		{"Mary"},
		{"Julie"},
		{"Kostya"},
		{"Darya"},
	}

	//Check Integration Test for registration
	for i := 0; i < len(users); i++ {
		assert.EqualValues(t, respondToRegistration(users[i]), http.StatusOK, "ERROR!")
	}

	//Incorrect Tests
	incorrectTests := []User{
		{"John"},
		{"Mary"},
		{"Julie"},
		{"Kostya"},
		{"Darya"},
		{"  Darya"},
		{""},
		{" "},
		{" Gosha"},
		{"    "},
		{"      "},
	}

	//Check Integration Test for registration
	for i := 0; i < len(incorrectTests); i++ {
		assert.EqualValues(t, respondToRegistration(incorrectTests[i]), http.StatusBadRequest, "ERROR!")
	}
}

func sendMessageToServer(sendUser SendUser) int {

	//Convert struct to JSON type
	postBody, _ := json.Marshal(sendUser)

	//Convert to *bytes.Buffer and initialization
	responseBody := bytes.NewBuffer(postBody)

	//Send request to server
	resp, err := http.Post(address+"send", appJson, responseBody)

	//Check request
	if err != nil {
		fmt.Println("Sorry, We didn't send your request. Try later...")
		return resp.StatusCode
	}

	//We close the require. Defer - works at the last
	defer resp.Body.Close()

	//Check reesponse body
	if err != nil {
		fmt.Println("Sorry, We didn't read response to your request. Try later...")
		return resp.StatusCode
	}

	return resp.StatusCode

}

//Checking send messages registreted users
func TestIntegrationSendMessageToRegistreteUsers(t *testing.T) {

	//Registrate users
	users := []User{
		{Username: "Abdulla"},
		{Username: "Ulfat"},
		{Username: "Andrey"},
		{Username: "Artur"},
		{Username: "Dasha"},
		{Username: "Sasha"},
	}

	//Registrate new users
	for _, user := range users {
		respondToRegistration(user)
	}

	//Tests
	tests := []SendUser{
		{"Abdulla", "Ulfat", "Hi! How are you?"},
		{"Ulfat", "Abdulla", "I'm fine and you?"},
		{"Abdulla", "Ulfat", "Me too!"},
		{"Andrey", "Artur", "Hi! How are you?"},
		{"Artur", "Andrey", "I'm fine and you?"},
		{"Andrey", "Artur", "Me too!"},
		{"Sasha", "Dasha", "Hi! How are you?"},
		{"Dasha", "Sasha", "I'm fine and you?"},
		{"Sasha", "Dasha", "Me too!"},
		{"Abdulla", "Sasha", "Hi!"},
	}

	//Check tests
	for _, value := range tests {
		assert.EqualValues(t, sendMessageToServer(value), http.StatusOK, "INCORRECT!")
	}

	//Check order user messages
	/*for _, value := range tests {
		assert.EqualValues(t, server.ge)
	}*/
}

//Checking send messages to Unregistrated users
func TestIntegrationSendMessageToUnRegistrateUsers(t *testing.T) {

	//Tests
	tests := []SendUser{
		{"Abdulla", "Nobody", "Hi! How are you?"},
		{"Ulfat", "Nobody", "I'm fine and you?"},
		{"Abdulla", "Ulfat", ""},
		{"R", "Abdulla", "Hello!"},
		{"Abdulla", " ", "     "},
		{"F", "D", "He"},
		{"Abdulla", "Ulfat", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. In vulputate mi commodo tortor aliquam consectetur. Vivamus fermentum sit amet mi a venenatis. Proin ut commodo nulla. Suspendisse pretium pharetra."},
	}

	//Check tests
	for _, value := range tests {
		assert.EqualValues(t, sendMessageToServer(value), http.StatusBadRequest, "INCORRECT!")
	}
}

func GetAllUserMessages(name string) int {

	var user User
	user.Username = name

	//Convert struct to JSON type
	postBody, _ := json.Marshal(user)

	//Convert to *bytes.Buffer and initialization
	responseBody := bytes.NewBuffer(postBody)

	//Send request to server
	resp, err := http.Post(address+"get", appJson, responseBody)

	//Check require
	if err != nil {
		fmt.Println("Sorry, We didn't send your request. Try later...")
		return resp.StatusCode
	}

	//We close the require. Defer - works at the last
	defer resp.Body.Close()

	//Check response body
	if err != nil {
		fmt.Println("Sorry, We didn't read response to your request. Try later...")
		return resp.StatusCode
	}

	return resp.StatusCode

}

func TestIntegrationGetAllRegistrateUserMessages(t *testing.T) {

	data := []User{
		{"Abdulla"},
		{"Ulfat"},
		{"Kim"},
		{"Chin"},
		{"In"},
		{"GoLang"},
	}

	for _, user := range data {
		respondToRegistration(user)
	}

	//Tests
	messages := []SendUser{
		{"Abdulla", "Ulfat", "Hi! How are you?"},
		{"Ulfat", "Abdulla", "I'm fine and you?"},
		{"Abdulla", "Ulfat", "Me too!"},
		{"Andrey", "Artur", "Hi! How are you?"},
		{"Artur", "Andrey", "I'm fine and you?"},
		{"Andrey", "Artur", "Me too!"},
		{"Sasha", "Dasha", "Hi! How are you?"},
		{"Dasha", "Sasha", "I'm fine and you?"},
		{"Sasha", "Dasha", "Me too!"},
		{"Abdulla", "Sasha", "Hi!"},
	}

	//Check tests
	for _, value := range messages {
		sendMessageToServer(value)
	}

	for _, value := range messages {
		assert.EqualValues(t, GetAllUserMessages(value.Reciever), http.StatusOK, "ERROR!")
	}

}

func TestIntegrationGetAllUnregistrateUserMessages(t *testing.T) {

	data := []User{
		{"Abdulla"},
		{"Ulfat"},
		{"Kim"},
		{"Chin"},
		{"In"},
		{"GoLang"},
	}

	for _, user := range data {
		respondToRegistration(user)
	}

	//Tests
	messages := []User{
		{"H"},
		{"G"},
		{"K"},
		{"     "},
		{""},
	}

	for _, value := range messages {
		assert.EqualValues(t, GetAllUserMessages(value.Username), http.StatusBadRequest, "ERROR!")
	}

}
