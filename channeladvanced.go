package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func multiplyBy5(c chan int, someValue int) {
	defer wg.Done()
	c <- someValue * 5
}
func main() {
	c := make(chan int,5)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go multiplyBy5(c, i)
	}
	wg.Wait()
	close(c)
	for item := range c {
		fmt.Println(item)
	}

}
