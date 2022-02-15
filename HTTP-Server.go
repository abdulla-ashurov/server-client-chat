package main

import (
	"fmt"
	"net/http"
)

//The problem with ServeHTTP method of HttpHandler is that is responds to all the request.
//ServeMux can accept a function for a specific route and when incoming request URL matches that route, that fubctuib will be executed.

func main() {
	//Create a new `ServeMux`
	mux := http.NewServeMux()

	//handle `/` route
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})

	//handle '/hello/golang' route
	mux.HandleFunc("/hello/golang", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang!")
	})

	//Listen and serve using `ServeMux`
	http.ListenAndServe(":80", mux)
}
