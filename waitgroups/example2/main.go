package main

import (
	"fmt"
	"sync"
)

// func main() {

// 	// challenge: modify this code so that the calls to updateMessage() on lines
// 	// 18, 21, and 24 run as goroutines, and implement wait groups so that
// 	// the program runs properly, and prints out three different messages.
// 	// Then, write a test for all three functions in this program: updateMessage(),
// 	// printMessage(), and main().

// 	msg = "Hello, world!"

// 	updateMessage("Hello, universe!")
// 	printMessage()

// 	updateMessage("Hello, cosmos!")
// 	printMessage()

// 	updateMessage("Hello, world!")

// 	printMessage()
// }

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	msg = "Hello World"

	wg.Add(1)
	go updateMessage("Hello, Universe")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, Cosmos")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, World")
	wg.Wait()
	printMessage()
}
