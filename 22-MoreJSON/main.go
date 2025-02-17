package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int	
	Platform string	`json:"website"`
	Password string	`json:"-"`
	Tags     []string	`json:"tags,omitempty"` 
}

func main() {
	fmt.Println("Welcome to JSON")
	EncodeJson()
}

func EncodeJson(){
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "Udemy", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "Udemy", "abcd123", []string{"fullstack", "js"}},
		{"DSA Bootcamp", 299, "Udemy", "hit123", nil},
	}

	//package this dat as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
