package main

import (
	"fmt"
	"time"
)

func concurrency() {
	// dones := make([]chan bool, 4)

	done := make(chan bool)
	// dones[0] = make(chan bool)
	go greet("Hello", done)
	// dones[1] = make(chan bool)
	go greet("World", done)
	// dones[2] = make(chan bool)
	go slowGreet("Slow Hello", done)
	// dones[3] = make(chan bool)
	go greet("Goodbye", done)

	// for _, done := range dones {
	// 	<-done
	// }
	for range done {
		// fmt.Println(doneChan)
	}

}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(phrase)
	doneChan <- true
	close(doneChan)
}

func greet(phrase string, doneChan chan bool) {
	fmt.Println(phrase)
	doneChan <- true
}
