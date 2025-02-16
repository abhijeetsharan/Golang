package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")

	content := "This needs to go in a file - Abhijeet Sharan"

	file, err := os.Create("./mygofile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mygofile.txt")
}

func readFile (filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside the file is: \n", string(databyte))
}

func checkNilError(err error){
	if err != nil {
		panic(err)
	}
}