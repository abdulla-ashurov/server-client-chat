package main

import (
	"fmt"
	"io"
	"net/http"
)

//Create handler struct
type HttpHandler struct{}

//Implement `ServerHTTP` method on 'HttpHandler' struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	//Write 'Hello' using 'io.WriteString' function
	io.WriteString(res, "Hello")

	//Write 'World' using 'fmt.Fprint' function
	fmt.Fprint(res, " World! ")

	//Write `❤️` using simple `Write` call
	res.Write([]byte("❤️"))
	res.Write([]byte("  Bye World!"))
}

func main() {
	//Create a new handler
	handler := HttpHandler{}

	//listen and serve
	http.ListenAndServe(":80", handler)
}
