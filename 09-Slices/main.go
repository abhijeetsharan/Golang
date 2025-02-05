package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the Slices!")

	var fruitList = []string{"Apple", "Banana", "Cherry", "Mango"}
	fmt.Printf("Type of fruitList is %T\n ", fruitList)

	fruitList = append(fruitList, "Strawberry", "Date")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	// for i := 0; i < len(highScores); i++ {
	// 	fmt.Println(highScores[i])
	// }

	// how to remove a value from slices based on index

	// Removing from slice
	var courses = []string{"Reactjs", "javascript", "Golang", "Swift", "Ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
