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
		{"Abdulla", "Ulfat", "Me too!"},
		{"Andrey", "Ulfat", "Hi! How are you?"},
		{"Artur", "Ulfat", "I'm fine and you?"},
		{"Andrey", "Ulfat", "Me too!"},
	}

	//Check tests
	for _, value := range tests {
		assert.EqualValues(t, server.SaveUserMessage(&value), true, "INCORRECT!")

	}

	userMessages := server.GetUserMessages(tests[0].Reciever)
	for i := 0; i < len(userMessages); i++ {
		assert.EqualValues(t, tests[i], userMessages[i], "INCORRECT!")
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
		{Username: "Anna"},
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
				"Hi! How are you?",
				"Me too!"},
		},
		{
			SendMsg: []server.SendUser{
				{"John", "Tom", "Hi! How are you?"},
				{"Alfred", "Tom", "Me too!"}},
			Correct: []string{
				"Hi! How are you?",
				"Me too!"},
		},
		{
			SendMsg: []server.SendUser{
				{"John", "Anna", "Hi! How are you?"},
				{"Alfred", "Anna", "Me too!"}},
			Correct: []string{
				"Hi! How are you?",
				"Me too!"},
		},
	}

	//Save Users messages
	for i := 0; i < len(tests); i++ {
		for _, user := range tests[i].SendMsg {
			server.SaveUserMessage(&user)
		}
	}

	for i := 0; i < len(tests); i++ {

		userMessages := server.GetUserMessages(tests[i].SendMsg[0].Reciever)

		for j := 0; j < len(tests[i].Correct); j++ {
			//Check tests
			assert.Equal(t, userMessages[j].Message, tests[i].Correct[j], "INCORRECT!")
		}
	}
}

//Checking get Unregistrated user messages
func TestGetUnRegistratedUserMessages(t *testing.T) {

	//Tests
	tests := []server.SendUser{
		{"Abdulla", "F", "Hi! How are you?"},
		{"Ulfat", "F", "I'm fine and you?"},
		{"Abdulla", "F", ""},
		{"R", "F", "Hello!"},
		{"Abdulla", " ", "     "},
		{"F", "D", "He"},
	}

	for i := 0; i < len(tests); i++ {

		userMessages := server.GetUserMessages(tests[i].Reciever)

		//Check tests
		for j := 0; j < len(userMessages); j++ {
			//Check tests
			assert.Equal(t, userMessages[j], tests[i].Message, "INCORRECT!")
		}
	}
}
