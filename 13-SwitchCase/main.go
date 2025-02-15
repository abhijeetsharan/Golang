package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	dice := rand.Intn(6) + 1

	fmt.Println("You rolled:", dice, "!")

	switch dice {
	case 1:
		fmt.Println("Bad luck! Try again.")
	case 2:
		fmt.Println("You got a small number.")
	case 3:
		fmt.Println("Not bad! Almost there.")
	case 4:
		fmt.Println("Good roll!")
	case 5:
		fmt.Println("You're getting close to the max!")
		fallthrough
	case 6:
		fmt.Println("WOW! You got the max!")
	default:
		fmt.Println("Something went wrong!")
	}
}