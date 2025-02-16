package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("New Email of this user is:", u.Email)
}

func main() {
	fmt.Println("Structs in Golang")

	abhijeet := User{"Abhijeet", "abhijeet@go.dev", true, 24}
	fmt.Println(abhijeet)
	fmt.Printf("abhijeet details are: %+v\n", abhijeet)
	fmt.Printf("Name is %v and email is %v. \n", abhijeet.Name, abhijeet.Email)
	abhijeet.GetStatus()
	abhijeet.NewMail()
	fmt.Printf("Name is %v and email is %v. \n", abhijeet.Name, abhijeet.Email)

}
