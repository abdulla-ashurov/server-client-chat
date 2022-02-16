package main

import "fmt"

//Create Reg structure for save date about registration user
//We have one field -> username
type Reg struct {
	username string `json:"Username"` // In json file we'll use variable name -> Username
}

//Information menu for user
func menu() {
	fmt.Println("\thelp\tprint all comands")
	fmt.Println("\treg\tregistration a new user[name]")
	fmt.Println("\tall\tshow all users in chatroom")
	fmt.Println("\tsend\tsend message to someone[whom, what]")
	fmt.Println("\tget\tget all my messages in chatroom")
}

func main() {
	menu()
}
