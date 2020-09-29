package main

import "fmt"

func multiplyBy5(c chan int, someValue int) {
	c <- someValue * 5
}
func main() {
	c := make(chan int)
	go multiplyBy5(c, 5)
	go multiplyBy5(c, 3)
	v1 := <-c
	v2 := <-c
	fmt.Println(v1, v2)
}
