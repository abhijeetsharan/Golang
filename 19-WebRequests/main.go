package main

import (
	"fmt"
	"io"
	"net/http"
)

// Define the URL to make the HTTP GET request to
const url = "https://jsonplaceholder.typicode.com/todos/1"

func main() {
	// Print a message indicating the start of the web request
	fmt.Println("Web Request")

	// Make an HTTP GET request to the specified URL
	response, err := http.Get(url)

	// Check if there was an error during the request
	if err != nil {
		panic(err) // If there's an error, terminate the program
	}

	// Print the type of the response object
	fmt.Printf("Response is of type : %T \n", response)

	// Ensure the response body is closed after the function exits
	defer response.Body.Close() // It's the caller's responsibility to close the connection

	// Read all the data from the response body
	databytes, err := io.ReadAll(response.Body)

	// Check if there was an error while reading the response body
	if err != nil {
		panic(err) // If there's an error, terminate the program
	}

	// Convert the byte slice to a string and print the content
	content := string(databytes)
	fmt.Println(content)
}