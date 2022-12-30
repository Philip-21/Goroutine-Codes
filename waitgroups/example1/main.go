package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {

	defer wg.Done() //defer deosnt exec until the func exits (it decements by 1)
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
	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d:, %s", i, x), &wg)
	}
	wg.Wait() //fetches data starting from the 1st index(0)

	wg.Add(1) //adding an item as it fetches to make the program execute

	printSomething("this is what is printed", &wg)

}
