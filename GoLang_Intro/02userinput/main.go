package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to the user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for goLang:")

	//comma ok || err err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating:", input)
	fmt.Printf("Type of Varibale is %T\n", input)
}
