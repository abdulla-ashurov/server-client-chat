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
type RegUser struct {
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
	var user RegUser

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
	//Print all users in map
	res.Write([]byte("\n"))
	for name := range users {
		res.Write([]byte(name + "\n"))
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
	messages[sendUser.Sender] = append([]SendUser{}, toKey)

	//Test Output Map
	/*for key, value := range messages {
		fmt.Println(key, value)
	}*/
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

	//Start server
	http.ListenAndServe(":80", mux)
}
