package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.tutorialspoint.com/go/index.htm"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of response is %T\n", response)
	defer response.Body.Close()
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content.txt)
}
