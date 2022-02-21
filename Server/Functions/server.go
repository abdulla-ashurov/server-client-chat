package server

import (
	"strings"
)

type HttpHandler struct {
}

//Create Reg structure for save date about registration user
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
var Users []string

//Global map for save messages
var messages = map[string][]SendUser{}

//Exist user or not
func IsExist(user string) ([]SendUser, bool) {
	if value, ok := messages[user]; ok {
		return value, true
	}
	return nil, false
}

//Check username for space, like "Sasha ", we get only "Sasha" without space
func spaceName(username string, index int) string {

	usernameNoSpace := ""

	for i := 0; i < index; i++ {
		usernameNoSpace += string(username[i])
	}

	return usernameNoSpace
}

//function registration
func CheckRegistration(user string) bool {

	//Check length username
	if len(user) == 0 {
		return false
	}

	//Check if username first letter is " " we respond error
	if user[0] == ' ' {
		return false
	}

	//Check username for space, like "Sasha ", we get only "Sasha" without space
	index := strings.Index(user, " ")
	if index != -1 {
		user = spaceName(user, index)
	}

	//Check we have this user or haven't
	for _, i := range Users {
		if i == user || user == "" {
			return false
		}
	}
	//Save a new user in Array
	Users = append(Users, user)
	return true
}

//Return all users
func GetAllUsersName() []string {
	return Users
}

//Save User Messages in Map
func SaveUserMessage(sendUser *SendUser) bool {

	if len(sendUser.Message) > 200 {
		return false
	}

	//Check we have message or haven't
	if len(sendUser.Message) == 0 {
		return false
	}

	//Check we have correct message or haven't: "    " -> ignore, "    hello world!" -> get only "hello world!"
	if sendUser.Message[0] == ' ' {
		sendUser.Message = strings.TrimLeft(sendUser.Message, " ")
	}

	//Check we have message or haven't
	if sendUser.Message == "" {
		return false
	} else {
		//Check we have sender or haven't
		for _, i := range Users {
			if i == sendUser.Sender {
				//Check we have reciever or haven't
				for _, i := range Users {
					if i == sendUser.Reciever {
						messages[sendUser.Reciever] = append(messages[sendUser.Reciever], *sendUser)
						//fmt.Println(messages[sendUser.Reciever])
						return true
					}
				}
			}
		}

	}
	return false
}

func GetUserMessages(reciever string) []string {

	//Create slice
	var userMessages []string

	//Check we have user or haven't
	for i := 0; i < len(Users); i++ {
		if Users[i] == reciever {
			//Check we have messages or not
			if sender, ok := messages[reciever]; ok {
				for _, i := range sender {
					msg := i.Sender + ": " + i.Message
					userMessages = append(userMessages, msg)
				}
				return userMessages
			} else {
				return userMessages
			}
		}
	}
	return userMessages
}
