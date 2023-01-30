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
