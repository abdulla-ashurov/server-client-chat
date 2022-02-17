package main

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
func reg(res http.ResponseWriter, req *http.Request) {
	//Create user type of RegUser
	var user User

	//Read Json file body and write to user structure
	err := json.NewDecoder(req.Body).Decode(&user)

	//Check Decode
	if err != nil {
		res.Write([]byte("Server: We didn't read your JSON file. Try later..."))
		return
	}

	//@TODO user maximum one time registration, add same users in map -> error

	//Respond message to user
	msg := "Welcome, " + user.Username + "!"
	res.Write([]byte(msg))

	//Save a new user in Map
	users[user.Username] = id
	id++
}

//Get all users
func all(res http.ResponseWriter, req *http.Request) {

	res.Write([]byte("\n"))

	//If we haven't any one, we print message about it
	if len(users) > 0 {
		//Print all users in map
		for name := range users {
			res.Write([]byte(name + "\n"))
		}
	} else {
		res.Write([]byte("Nobody in chat right now. Try later..."))
	}

}

//Send message
func send(res http.ResponseWriter, req *http.Request) {
	//Create object of type SendUser
	var sendUser SendUser

	//Read json file body and write to user structure
	err := json.NewDecoder(req.Body).Decode(&sendUser)

	//Check Decode
	if err != nil {
		res.Write([]byte("Server: We didn't read your JSON file. Try later..."))
		return
	}

	//Save user message in Map
	toKey := SendUser{
		Sender:   sendUser.Sender,
		Reciever: sendUser.Reciever,
		Message:  sendUser.Message,
	}
	messages[sendUser.Reciever] = append([]SendUser{}, toKey)

	//Test Output Map
	/*for key, value := range messages {
		fmt.Println(key, value)
	}*/
}

//Get user Message
func getMsg(res http.ResponseWriter, req *http.Request) {
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
	for i := 0; i < len(recievers); i++ {
		msg := recievers[i].Sender + ": " + recievers[i].Message
		res.Write([]byte(msg + "\n"))
	}
}

func main() {
	//Create a new `ServeMux`
	//ServeMux can accept a function for a specific route and when incoming request
	//URL mathces that route, that will be executed
	mux := http.NewServeMux()

	//Handle `/` route, where -> res - respond to clien, req - client request
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World!"))
	})

	//Handle '/reg'
	mux.HandleFunc("/reg", reg)

	//Handle '/all'
	mux.HandleFunc("/all", all)

	//Handle `/send`
	mux.HandleFunc("/send", send)

	//Handle `/get`
	mux.HandleFunc("/get", getMsg)

	//Start server
	http.ListenAndServe(":80", mux)
}
