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
var Messages = map[string][]SendUser{}

//function registration
func Reg(user string) bool {
	//Check we have this user or haven't
	if _, ok := users[user]; ok || user == "" {
		return false
	} else {
		//Save a new user in Map
		users[user] = id
		id++
		return true
	}
}

//function registration
func RespondRegistration(res http.ResponseWriter, req *http.Request) {
	//Create user type of RegUser
	var user User

	//Read Json file body and write to user structure
	err := json.NewDecoder(req.Body).Decode(&user)

	//Check Decode
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	//msg := "Welcome, " + user.Username + "!"

	//registration user
	if Reg(user.Username) {
		//Respond to server
		res.WriteHeader(http.StatusOK)
	} else {
		res.WriteHeader(http.StatusBadRequest)
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
func RespondAllUsers(res http.ResponseWriter, req *http.Request) {
	//Transition a new line
	res.Write([]byte("\n"))

	//If we haven't any one, we print message about it
	if len(users) > 0 {
		res.Write([]byte(PrintAll()))
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}

}

//Save User Messages in Map
func SaveMsg(sendUser *SendUser) bool {
	//Check user we have or haven't
	if sendUser.Message == "" {
		return false
	} else if _, ok := users[sendUser.Reciever]; ok {
		Messages[sendUser.Reciever] = append(Messages[sendUser.Reciever], *sendUser)
		return true
	} else {

		return false
	}
}

//Send message
func SaveSendMessages(res http.ResponseWriter, req *http.Request) {
	//Create object of type SendUser
	var sendUser SendUser

	//Read json file body and write to user structure
	err := json.NewDecoder(req.Body).Decode(&sendUser)

	//Check Decode
	if err != nil {
		res.Write([]byte("Server: We didn't read your JSON file. Try later..."))
		res.WriteHeader(http.StatusBadRequest)
	}

	if SaveMsg(&sendUser) {
		res.WriteHeader(http.StatusOK)
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}

}

func GetMessages(recievers []SendUser) string {
	//Check we have messages or not
	msg := ""
	if len(recievers) > 0 {
		//msg += "\n"
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
func RespondtMessages(res http.ResponseWriter, req *http.Request) {
	//Create object of type User
	var user User

	//Read json file body and write to user structure
	err := json.NewDecoder(req.Body).Decode(&user)

	//Check Decode
	if err != nil {
		res.Write([]byte("Server: We didn't read your JSON file. Try later..."))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	//Create array of type []SendUser
	if value, ok := Messages[user.Username]; ok {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(GetMessages(value)))
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
	//recievers := messages[user.Username]

}
