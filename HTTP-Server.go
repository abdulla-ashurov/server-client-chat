package main

import (
	"fmt"
	"net/http"
)

//The problem with ServeHTTP method of HttpHandler is that is responds to all the request.
//ServeMux can accept a function for a specific route and when incoming request URL matches that route, that fubctuib will be executed.

func main() {
	// We can pass nil as the value of handler, It's equal http.DefaultServeMux which is a global ServeMux instance.

	//handle `/` route to `http.DefaultServerMux`
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})

	//handle '/hello/golang' route to `http.DefaultServerMux`
	http.HandleFunc("/hello/golang", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang!")
	})

	//Listen and serve using `http.DefaultServerMux`
	http.ListenAndServe(":80", nil)
}
