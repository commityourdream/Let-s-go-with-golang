package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags, omiempty"`
}

func main() {
	EncodeJson()
	decodeJson()

}

func EncodeJson() {
	ucourses := []course{
		{"ReactJs", 200, "Udemy.in", "abc123", []string{"web-dev", "js"}},
		{"Python", 200, "Udemy.in", "def321", []string{"back-end", "py"}},
		{"GoLang", 200, "Udemy.in", "ghi456", []string{"back-dev", "gl"}},
	}
	// package this data as JSON data

	finalJson, err := json.MarshalIndent(ucourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func decodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename":"Python",
		"Price":200,
		"website":"Udemy.in",
		"tags":["back-end","py"]
	}
		`)

	var onlinecourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &onlinecourse)
		fmt.Printf("%#v\n", onlinecourse)
	} else {
		fmt.Println("Json was not valid")
	}
	// some case where you want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is:%T\n", k, v, v)
	}

}
