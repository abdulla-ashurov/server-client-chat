package main

import (
	server "Chat_Server_Client/server/Functions"
	"fmt"
	"testing"
)

//Create structure for Test Regisration
type RegUser struct {
	User    string
	Correct bool
}

func TestReg(t *testing.T) {

	//Tests
	dataCorrect := []RegUser{
		{User: "Abdulla", Correct: true},
		{User: "Ulfat", Correct: true},
		{User: "Andrey", Correct: true},
		{User: "Sasha", Correct: true},
		{User: "Abdulla", Correct: false},
		{User: "Ulfat", Correct: false},
	}

	//Check tests
	for _, data := range dataCorrect {
		if server.Reg(data.User) == data.Correct {
			fmt.Println("OK!")
		} else {
			fmt.Println("Incorrect!")
		}
	}

}
