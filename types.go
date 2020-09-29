package main

import (
	"fmt"
)

//func add(x float64, y float64) float64 {
func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}
func main() {
	//var num1, num2 float64 = 5.6, 9.5
	//var num2 float64 = 9.5
	num1, num2 := 5.6, 9.5
	var x int = 6
	var b float64 = float64(x)
	fmt.Println("Converted", b)
	fmt.Println(add(num1, num2))
	w1, w2 := "Hey", "there"
	fmt.Println(multiple(w1, w2))
}
