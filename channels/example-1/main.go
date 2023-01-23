package main

import (
	"fmt"
	"strings"
	"sync"
)

func shout(ping <-chan string, pong chan<- string) {
	defer wg.Done()
	for {
		//accept a value from ping
		s := <-ping //ping receives from the channel stored in a variable s

		///pong sends data to the channel received from ping defined with the variable s
		pong <- fmt.Sprintf("%s!!", strings.ToUpper(s))
	}
}

var wg sync.WaitGroup

func main() {

	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type Something and Press ENTER ,enter Q to quit ")
	for {
		//print a prompt
		fmt.Print("->")
		//get user input
		var userinput string
		//accepting user inputs and prints them in uppercase
		_, _ = fmt.Scanln(&userinput)
		if userinput == strings.ToLower("q") {
			break //quit the program by clixking q
		}

		//send data to ping channel,then ping receives the data userinput )
		ping <- userinput
		//wait on the response
		response := <-pong //(pong receives the data from ping with var response, then sends the data to the channel)
		fmt.Println("Response", response)

	}
	fmt.Println("All done, Closing Channels")
	close(ping)
	close(pong)

}
