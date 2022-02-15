package main

import (
	"fmt"
	"net/http"
)

//Create handler struct
type HttpHandler struct{}

//Implement `ServerHTTP` method on 'HttpHandler' struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	//Create response binary data
	data := []byte("Hello World!") //slice of bytes

	//Write `data` to response
	res.Write(data)
	fmt.Print(*req)
}

func main() {
	//Create a new handler
	handler := HttpHandler{}

	//listen and serve
	http.ListenAndServe(":80", handler)
}
