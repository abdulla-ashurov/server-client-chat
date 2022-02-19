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
