package main

import (
	server "Chat_Server_Client/server/Functions"
	"encoding/json"
	"net/http"
)

//function registration
func RespondRegistration(res http.ResponseWriter, req *http.Request) {
	//Create user type of RegUser
	var user server.User

	//Read Json file body and write to user structure
	err := json.NewDecoder(req.Body).Decode(&user)

	//Check Decode
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	//msg := "Welcome, " + user.Username + "!"

	//registration user
	if server.CheckRegistration(user.Username) {
		//Respond to server
		res.WriteHeader(http.StatusOK)
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}

//Get all users
func RespondAllUsers(res http.ResponseWriter, req *http.Request) {
	//Transition a new line
	res.Write([]byte("\n"))

	//If we haven't any one, we print message about it
	if len(server.Users) > 0 {
		userNames := server.GetAllUsersName()
		for _, name := range userNames {
			res.Write([]byte(name))
		}
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}

}

//Send message
func SaveSendMessages(res http.ResponseWriter, req *http.Request) {
	//Create object of type SendUser
	var sendUser server.SendUser

	//Read json file body and write to user structure
	err := json.NewDecoder(req.Body).Decode(&sendUser)

	//Check Decode
	if err != nil {
		res.Write([]byte("Server: We didn't read your JSON file. Try later..."))
		res.WriteHeader(http.StatusBadRequest)
	}

	if server.SaveUserMessage(&sendUser) {
		res.WriteHeader(http.StatusOK)
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}

}

//Get user Message
func RespondtMessages(res http.ResponseWriter, req *http.Request) {
	//Create object of type User
	var user server.User

	//Read json file body and write to user structure
	err := json.NewDecoder(req.Body).Decode(&user)

	//Check Decode
	if err != nil {
		res.Write([]byte("Server: We didn't read your JSON file. Try later..."))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	//Create array of type []SendUser
	if value, ok := server.IsExist(user.Username); ok {
		res.WriteHeader(http.StatusOK)
		userMessages := server.GetUserMessages(value[0].Reciever)

		for i := 0; i < len(userMessages); i++ {
			res.Write([]byte(userMessages[i]))
		}
	} else {
		res.WriteHeader(http.StatusBadRequest)
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
	mux.HandleFunc("/reg", RespondRegistration)

	//Handle '/all'
	mux.HandleFunc("/all", RespondAllUsers)

	//Handle `/send`
	mux.HandleFunc("/send", SaveSendMessages)

	//Handle `/get`
	mux.HandleFunc("/get", RespondtMessages)

	//Start server
	http.ListenAndServe(":80", mux)

}
