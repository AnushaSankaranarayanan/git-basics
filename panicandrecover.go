package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("recovered in cleanup..", r)
	}
	wg.Done()
}
func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		if i == 2 {
			panic("oh dear...a 2")
		}
		fmt.Println(s)
	}
}

func main() {
	wg.Add(1)
	go say("There")
	wg.Add(1)
	go say("Hey")
	wg.Wait()
}
