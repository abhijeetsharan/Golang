package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println("List of all languages:", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "GO")
	fmt.Println("List of all languages:", languages)

	// loops are interesting in golang
	for key, value := range languages {
		fmt.Println(key, "shorts for:", value)
	}
}
