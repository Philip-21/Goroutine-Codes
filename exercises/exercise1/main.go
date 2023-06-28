package main

import (
	"fmt"
	"strconv"
)

/*Launch multiple Goroutines and each goroutine adding values to a Channel*/

func main() {
	ch := make(chan string)

	for i := 0; i < 20; i++ {
		//start the goroutine
		go func(i int) {
			//a for loop that control's the number of times a specific value is added to the channel
			for j := 0; j < 10; j++ {
				ch <- "Goroutine Value" + strconv.Itoa(i)
			}
		}(i)
		
	}
	// controls how many times the main Goroutine retrieves values from the channel and prints them.
	for k :=0; k<100; k++ {
		fmt.Println(k , <-ch)
	}
}
/*
i is responsible for assigning a unique identifier to each Goroutine and determining the value it adds to the channel.
j ensures that each Goroutine adds its value to the channel exactly 10 times.
k controls how many times the main Goroutine retrieves values from the channel and prints them.
*/