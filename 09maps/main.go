package main

import (
	"fmt"
)


func main()  {
	
	languages := make(map[string]string)
	languages["JS"]="JavaScript"
	languages["PY"]="Python"
	languages["RB"]="Ruby"

	fmt.Println("List Of all Languages:",languages)
	fmt.Println("JS shorts for:",languages["JS"])

	delete(languages,"RB")
	fmt.Println("List Of all Languages:",languages)

	// loops

	for key, value := range languages{
		fmt.Printf("For key %v , Value is %v\n", key,value)
		
	}
}