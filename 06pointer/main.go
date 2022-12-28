package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println("Value of pointer is:", ptr)

	number := 26
	var ptr = &number
	fmt.Println("Value of actual pointer is:", ptr)
	fmt.Println("Value of actual pointer is:", *ptr)

	*ptr = *ptr + 4
	fmt.Println("New Value of number is:", number)
}
