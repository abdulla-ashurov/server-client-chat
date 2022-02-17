package unit_test

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
		if server.Reg(value) {
			fmt.Println("Test " + strconv.Itoa(count) + ": Correct!")
		} else {
			fmt.Println("Test " + strconv.Itoa(count) + ": Incorrect!")
		}
	}

	fmt.Print(user)
}
