package main

import (
	server "Chat_Server_Client/server/Functions"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetUserMessages struct {
	SendMsg []server.SendUser
	Correct []string
}

//Checking Unregistrate Users to sign up
func TestUnregistrateUsers(t *testing.T) {

	//Tests
	tests := []server.User{
		{Username: "Abdulla"},
		{Username: "Ulfat"},
		{Username: "Andrey"},
		{Username: "Artur"},
		{Username: "Dasha"},
		{Username: "Sasha"},
		{Username: "S "},
	}

	//Check correct Tests
	for _, user := range tests {
		assert.EqualValues(t, server.CheckRegistration(user.Username), true, "INCORRECT!")
	}
}

//Checking Registrated Users to sign up
func TestRegistrateUsers(t *testing.T) {

	//Correct Tests
	tests := []server.User{
		{Username: "Abdulla"},
		{Username: "Ulfat"},
		{Username: "Andrey"},
		{Username: "Artur"},
		{Username: "Dasha"},
		{Username: "Sasha"},
	}

	//Registrate new users
	for _, user := range tests {
		server.CheckRegistration(user.Username)
	}

	//Incorrect Tests
	incorrectTests := []server.User{
		{Username: "Abdulla"},
		{Username: "Ulfat"},
		{Username: "Andrey"},
		{Username: "Artur"},
		{Username: "Dasha"},
		{Username: "Sasha"},
		{Username: ""},
		{Username: " "},
		{Username: " Gosha"},
		{Username: "    "},
		{Username: "Abdulla "},
	}

	//Check incorrect users
	for _, user := range incorrectTests {
		assert.EqualValues(t, server.CheckRegistration(user.Username), false, "INCORRECT!")
	}
}

//Checking send messages registreted users
func TestSendMessageToRegistreteUsers(t *testing.T) {

	//Registrate users
	users := []server.User{
		{Username: "Abdulla"},
		{Username: "Ulfat"},
		{Username: "Andrey"},
		{Username: "Artur"},
		{Username: "Dasha"},
		{Username: "Sasha"},
	}

	//Registrate new users
	for _, user := range users {
		server.CheckRegistration(user.Username)
	}

	//Tests
	tests := []server.SendUser{
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
		assert.EqualValues(t, server.SaveUserMessage(&value), true, "INCORRECT!")
	}

}

//Checking send messages to Unregistrated users
func TestSendMessageToUnRegistrateUsers(t *testing.T) {

	//Tests
	tests := []server.SendUser{
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
		assert.EqualValues(t, server.SaveUserMessage(&value), false, "INCORRECT!")
	}
}

//Checking get registrated user messages
func TestGetRegistratedUserMessages(t *testing.T) {

	//Registrate users
	users := []server.User{
		{Username: "Tom"},
		{Username: "John"},
		{Username: "Alfred"},
	}

	//Registrate new users
	for _, user := range users {
		server.CheckRegistration(user.Username)
	}

	//Tests
	tests := []GetUserMessages{
		{
			SendMsg: []server.SendUser{
				{"Tom", "John", "Hi! How are you?"},
				{"Alfred", "John", "Me too!"}},
			Correct: []string{
				"Tom: Hi! How are you?",
				"Alfred: Me too!"},
		},
		{
			SendMsg: []server.SendUser{
				{"John", "Tom", "Hi! How are you?"},
				{"Alfred", "Tom", "Me too!"}},
			Correct: []string{
				"John: Hi! How are you?",
				"Alfred: Me too!"},
		},
	}

	//Save Users messages
	for i := 0; i < len(tests); i++ {
		for _, user := range tests[i].SendMsg {
			server.SaveUserMessage(&user)
		}
	}

	for i := 0; i < len(tests); i++ {

		correct := ""
		for j := 0; j < len(tests[i].Correct[i]); j++ {
			correct += tests[i].Correct[j] + "\n"
		}

		//Check tests
		assert.Equal(t, server.GetUserMessages(tests[i].SendMsg[0].Reciever), correct, "INCORRECT!")
	}

}
