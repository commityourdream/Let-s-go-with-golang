package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://loc.dev:3000/learn?coursename=reactjs&paymentid=hgfb546yts"

func main() {
	// parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.User)
	// fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("The type of qparams are %T\n", qparams)
	//fmt.Println(qparams)

	for _, val := range qparams {
		fmt.Println("params is ", val)
	}

	partsofUrl := &url.URL{
		Scheme:"https",
		Host:"lco.dev",
		Path:"/tutcss",
		RawPath:"user-amit",
	}
	anotherURL := partsofUrl.String()
	fmt.Println(anotherURL)

}
