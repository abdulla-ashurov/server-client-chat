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

func main() {
	//Create a new `ServeMux`
	//ServeMux can accept a function for a specific route and when incoming request
	//URL mathces that route, that will be executed
	mux := http.NewServeMux()

	//Handle `/` route
	//res - respond to client
	//req - client requst
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World!"))
	})

	//Handle '/reg'
	mux.HandleFunc("/reg", func(res http.ResponseWriter, req *http.Request) {
		//Create user type of RegUser
		var user RegUser

		//Read Json file body and write to user structure
		err := json.NewDecoder(req.Body).Decode(&user)

		//Check Decode
		if err != nil {
			res.Write([]byte("Server: We didn't read your JSON file. Try later..."))
			return
		}

		//Respond message to user
		msg := "Welcome, " + user.Username + "!"
		res.Write([]byte(msg))

	})

	//Start server
	http.ListenAndServe(":80", mux)
}
