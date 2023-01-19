package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "This needs to be go in a file"
	file, err := os.Create("./sample.txt")

	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is:", length)
	defer file.Close()
	readFile("./sample.txt")
}

func readFile(filename string)  {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside a file is:\n",string(databyte))
	
}
