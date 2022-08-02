package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	num, err := rand.Int(rand.Reader, big.NewInt(27))
	if err != nil {
		panic(err)
	}
	n := num.Int64()

	fmt.Println("Hello", n, n)
}