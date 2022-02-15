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

		//get response headers
		header := res.Header()

		//set content type header
		header.Set("Content-type", "application/json")

		//reset data header (inline call)
		res.Header().Set("Data", "01/01/2022")

		//set status header
		res.WriteHeader(http.StatusBadRequest) // http.StatusBadRequest == 400

		//respond with a JSON string
		fmt.Fprint(res, `{"status":"FAILURE"}`)
	})

	//Listen and serve using `http.DefaultServerMux`
	http.ListenAndServe(":80", nil)
}
