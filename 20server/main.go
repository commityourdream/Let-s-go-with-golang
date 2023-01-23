package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	PerformGetRequest()
	PerformPostRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content Length is:", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":0,
			"platform":"Google.com"
		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

//perform postform reques

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// form data
	data := url.Values{}
	data.Add("Firstname", "Amit")
	data.Add("Lastname", "Manjhi")
	data.Add("Email", "myselfakm95@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
