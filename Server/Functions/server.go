package server

import (
	"encoding/json"
	"net/http"
)

//Create handler structure without fields
type HttpHandler struct {
}

//Create Reg structure for save date about registration user
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

//Global map for save users
var users = map[string]uint{}

//Count for id users
var id uint = 1

//Global map for save messages
var messages = map[string][]SendUser{}

//function registration
func Reg(user User) bool {
	//Check we have this user or haven't
	if _, ok := users[user.Username]; ok {
		return false
	} else {
		//Save a new user in Map
		users[user.Username] = id
		id++
		return true
	}
}

//function registration
func ResReg(res http.ResponseWriter, req *http.Request) {
	//Create user type of RegUser
	var user User

	//Read Json file body and write to user structure
	err := json.NewDecoder(req.Body).Decode(&user)

	//Check Decode
	if err != nil {
		res.Write([]byte("Server: We didn't read your JSON file. Try later..."))
		return
	}

	msg := "Welcome, " + user.Username + "!"

	//registration user
	if Reg(user) {
		//Respond to server
		res.Write([]byte(msg))
	} else {
		res.Write([]byte("Please use another nickname!"))
	}
}

//Print all users
func PrintAll() string {
	//Save all users in string variable
	allUsers := ""
	for name := range users {
		allUsers += name
	}

	//Return all users
	return allUsers
}

//Get all users
func All(res http.ResponseWriter, req *http.Request) {
	//Transition a new line
	res.Write([]byte("\n"))

	//If we haven't any one, we print message about it
	if len(users) > 0 {
		res.Write([]byte(PrintAll()))
	} else {
		res.Write([]byte("Nobody in chat right now. Try later..."))
	}

}

func SaveMsg(sendUser *SendUser) bool {
	//Check user we have or haven't
	if _, ok := users[sendUser.Reciever]; ok {
		messages[sendUser.Reciever] = append(messages[sendUser.Reciever], *sendUser)
		return true
	} else {

		return false
	}
}

//Send message
func Send(res http.ResponseWriter, req *http.Request) {
	//Create object of type SendUser
	var sendUser SendUser

	//Read json file body and write to user structure
	err := json.NewDecoder(req.Body).Decode(&sendUser)

	//Check Decode
	if err != nil {
		res.Write([]byte("Server: We didn't read your JSON file. Try later..."))
		return
	}

	if SaveMsg(&sendUser) {
		res.Write([]byte("OK"))
	} else {
		res.Write([]byte(sendUser.Reciever + " is not in the chat"))
	}

}

func GetMessages(recievers []SendUser) string {
	//Check we have messages or not
	msg := ""
	if len(recievers) > 0 {
		msg += "\n"
		//Respond all messages
		for i := 0; i < len(recievers); i++ {
			msg += recievers[i].Sender + ": " + recievers[i].Message + "\n"
		}
		return msg
	} else {
		msg = "Empty"
		return msg
	}
}

//Get user Message
func GetMsg(res http.ResponseWriter, req *http.Request) {
	//Create object of type User
	var user User

	//Read json file body and write to user structure
	err := json.NewDecoder(req.Body).Decode(&user)

	//Check Decode
	if err != nil {
		res.Write([]byte("Server: We didn't read your JSON file. Try later..."))
		return
	}

	//Create array of type []SendUser
	recievers := messages[user.Username]

	//Check message we have or haven't and print
	res.Write([]byte(GetMessages(recievers)))
}
