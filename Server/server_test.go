package main

import (
	server "Chat_Server_Client/server/Functions"
	"fmt"
	"strconv"
	"testing"
)

func regTest(t *testing.T) {
	user := []server.User{
		{Username: "Some"},
		{Username: "Some"},
		{Username: "Abdulla"},
		{Username: "Some"},
	}

	count := 1
	for _, value := range user {
		if server.Reg(value) != true {
			t.Errorf("Test " + strconv.Itoa(count) + ": Incorrect!")
		}
	}

	fmt.Print(user)
}
