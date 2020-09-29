package main

import (
	"fmt"
	"math"
	"math/rand"
)

func foo() {
	fmt.Println("Square root of 4 is", math.Sqrt(4))
	fmt.Println("Random number lesss than 100 is", rand.Intn(100))
}

func main() {
	fmt.Println("Welcome to go")
	foo()
}
