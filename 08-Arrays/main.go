package main

import "fmt"

func main() {
	fmt.Println("Welcome to the arrays")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Cherry"
	fruitList[3] = "Mango"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit length is: ", len(fruitList))

	var vegList = [5]string{"Potato", "beans", "mushroom"}
	fmt.Println("Veg list is: ", vegList)
	fmt.Println("Veg list  length is: ", len(vegList))

}