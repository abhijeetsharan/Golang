package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	n, _ := rand.Int(rand.Reader, big.NewInt(10000)) 
	fmt.Println("Secure Random Number:", n)
}
