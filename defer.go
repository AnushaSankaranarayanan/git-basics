package main

import (
	"fmt"
	"sync"
	"time"
)

func foo() {
	defer fmt.Println("Done!")
	defer fmt.Println("Are we done yet?")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Doing some stuff, who knows what?")
}

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	foo()
	wg.Add(1)
	go say("There")
	wg.Add(1)
	go say("Hey")
	wg.Wait()
}
