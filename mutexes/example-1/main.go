package main

import (
	"fmt"
	"sync"
)

// race is ued in this exampe to access the data safely
var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock() //the calling goroutine blocks until the mutex is available.
	msg = s
	m.Unlock()
}

func main() {
	msg := "Hello, World"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, universe", &mutex)
	go updateMessage("hello, cosmos", &mutex)
	wg.Wait()
	fmt.Println(msg)
}

//type go run -race .  to check if it executes well
//it prints out the data without showing error
