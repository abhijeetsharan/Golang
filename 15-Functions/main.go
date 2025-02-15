package main

import "fmt"

func greet(){
	fmt.Println("Hello !")
}

// Variadic Functions (Accepting Multiple Arguments)
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func main() {
	fmt.Println("Welcome to the functions in Golang")
	greet()
	result := adder(5 ,10)
	fmt.Println(result)
	fmt.Println(sum(1,2,3,4,5,6))

	// Anonymous Functions
	func(msg string) {
		fmt.Println(msg)
	}("Hello, Go!!!")

	// Defer (Delays Execution Until funtion ends) for cleanup tasks, like closing files or database connections.
	defer fmt.Println("This prints last")
	fmt.Println("This prints first")
}

