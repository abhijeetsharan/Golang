package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for i := 0; i<len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days{
	// 	fmt.Printf("%d %s\n", index, day)
	// }

	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 5 {
			break
		}
		fmt.Println(rougueValue)
		rougueValue++
	}
}