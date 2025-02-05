package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println("Current Time is: ", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.October, 15, 23, 23, 0, 0, time.UTC)
	fmt.Println("Created Date is: ", createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
