package main

import (
	server "CHAT_SERVER_CLIENT/server/Functions"
	"net/http"
)

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
	mux.HandleFunc("/reg", server.ResReg)

	//Handle '/all'
	mux.HandleFunc("/all", server.All)

	//Handle `/send`
	mux.HandleFunc("/send", server.Send)

	//Handle `/get`
	mux.HandleFunc("/get", server.GetMsg)

	//Start server
	http.ListenAndServe(":80", mux)
}
