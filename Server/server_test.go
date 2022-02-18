package main

import (
	server "Chat_Server_Client/server/Functions"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Create structure for Test Regisration
type RegUser struct {
	User    string
	Correct bool
}

//Create structure for Save Messages
type SaveMsg struct {
	SendMsg server.SendUser
	Correct bool
}

//Create structure for Get Messages
type GetMsg struct {
	SendMsg []server.SendUser
	Correct []string
}

//Registraion Test
func TestReg(t *testing.T) {

	fmt.Println("REGISTRATION TESTING!")
	//Tests
	dataCorrect := []RegUser{
		{User: "Abdulla", Correct: true},
		{User: "Ulfat", Correct: true},
		{User: "Andrey", Correct: true},
		{User: "Sasha", Correct: true},
		{User: "Abdulla", Correct: false},
		{User: "Ulfat", Correct: false},
		{User: "", Correct: false},
	}

	//Check tests
	for _, data := range dataCorrect {
		assert.EqualValues(t, server.Reg(data.User), data.Correct, "INCORRECT!")
	}

	fmt.Println("OK!")
}

//Test for Send Messages
func TestSendMsg(t *testing.T) {

	fmt.Println("SEND MESSAGES TESTING!")

	//Tests
	data := []SaveMsg{
		{
			SendMsg: server.SendUser{"Abdulla", "Ulfat", "Hi! How are you?"},
			Correct: true,
		},
		{
			SendMsg: server.SendUser{"Ulfat", "Abdulla", "Hi! How are you?"},
			Correct: true,
		},
		{
			SendMsg: server.SendUser{"Abdulla", "Roma", "Hi! How are you?"},
			Correct: false,
		},
		{
			SendMsg: server.SendUser{"Ulfat", "U", "Hi! How are you?"},
			Correct: false,
		},
		{
			SendMsg: server.SendUser{"Ulfat", "Ukljkljk", ""},
			Correct: false,
		},
	}

	//Check tests
	for _, value := range data {
		assert.EqualValues(t, server.SaveMsg(&value.SendMsg), value.Correct, "INCORRECT!")
	}

	fmt.Println("OK!")
}

//Test for Get Messages
func TestGetMsg(t *testing.T) {

	fmt.Println("GET MESSAGES TESTING!")

	//Tests
	data := GetMsg{
		SendMsg: []server.SendUser{
			{"Abdulla", "Ulfat", "Hi! How are you?"},
			{"Sasha", "Ulfat", "Hi! How are you?"},
		},
		Correct: []string{
			"\nAbdulla: Hi! How are you?\n",
			"Sasha: Hi! How are you?\n",
		},
	}

	correct := ""
	for i := 0; i < len(data.Correct); i++ {
		correct += data.Correct[i]
	}

	//Check tests
	assert.Equal(t, server.GetMessages(data.SendMsg), correct, "INCORRECT!")

	fmt.Println("OK!")

}
