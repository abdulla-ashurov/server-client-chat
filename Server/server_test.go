package main

import (
	server "Chat_Server_Client/server/Functions"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	}

	//Check tests
	for _, value := range tests {
		assert.EqualValues(t, server.SaveUserMessage(&value), false, "INCORRECT!")
	}
}
