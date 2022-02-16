package main

import (
	"net/http"
)

//Create handler structure without fields
type HttpHandler struct {
}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

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
		res.Write([]byte("Welcome to registration!"))
	})

	//Start server
	http.ListenAndServe(":80", mux)
}
