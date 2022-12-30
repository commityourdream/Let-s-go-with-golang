package main

import "fmt"



func main()  {
	
	// no inherietance, super parent concept in golang

	Amit := User{"Amit","amit@go.dev",true,25}
	fmt.Println(Amit)
	fmt.Printf("Amit details are :%+v\n",Amit)
	fmt.Printf("Name is %v and email is %v.",Amit.Name,Amit.Email )

}
type User struct{
	Name string
	Email string
	Status bool
	Age int
}