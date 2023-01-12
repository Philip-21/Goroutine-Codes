package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done() //defer doesnt exec until the func exits (it decrements by 1)
	fmt.Println(s)
}

func main() {

	var wg sync.WaitGroup

	words := []string{
		"Alpha",
		"Beta",
		"Delta",
		"gamma",
		"Epselon",
		"theta",
	}
	wg.Add(len(words)) //ading counters for the goroutine

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d:, %s", i, x), &wg)
	}
	wg.Wait() //waits for the goroutine to exec until the counter hits 0

	wg.Add(1) //adding  a count of 1 for a goroutine
	printSomething("this is what is printed", &wg)

}
