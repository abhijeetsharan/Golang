package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	abhijeet := User{"Abhijeet", "abhijeet@go.dev", true, 24}
	fmt.Println(abhijeet)
	fmt.Printf("abhijeet details are: %+v\n", abhijeet)
	fmt.Printf("Name is %v and email is %v.", abhijeet.Name, abhijeet.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
