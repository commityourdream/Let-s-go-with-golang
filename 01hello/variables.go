package main

import (
	"fmt"
	"reflect"
)

const LoginToken string = "gfdaswqzx" //public

func main() {
	var username string = "Amit"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)
	fmt.Println("Variable is of type:", reflect.ValueOf(username).Kind())

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)
	var abc float64 = 10.45
	fmt.Println(abc)
	fmt.Printf("Variable is of type: %T\n", abc)

	var randomvar int
	fmt.Println(randomvar)
	fmt.Printf("Variable is of type: %T\n", randomvar)

	//implicittype
	var website = "google.com"
	fmt.Println(website)
	fmt.Printf("Variable is type of %T\n", website)

	noOfUsers := 3000
	fmt.Println(noOfUsers)
	fmt.Printf("Variable is type of %T\n", noOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is type of %T\n", LoginToken)

}
