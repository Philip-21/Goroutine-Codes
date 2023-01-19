package main

import (
	"fmt"
	"time"
)

//using the select statement for channels(synonymous to switch statement)

func server1(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "This is from the server 1"
	}
}

func server2(ch chan string) {

	for {
		time.Sleep(3 * time.Second)
		ch <- "This is from the server 2"
	}
}
func main() {
	fmt.Println("Select statement with channels")
	fmt.Println("-----------------")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	for {
		//use the select statement in receiving data
		//if theres is more than one case that the select can match
		// it chooses one case at random it can print either case3 or 4
		//and either case 1 or 2 picking them at random
		select {
		case s1 := <-channel1:
			fmt.Println("Case One:", s1)
		case s2 := <-channel1:
			fmt.Println("Case Two:", s2)
		case s3 := <-channel2:
			fmt.Println("Case three:", s3)
		case s4 := <-channel2:
			fmt.Println("Case four:", s4)
			//default:
			//default avoid deadlock (when a  goroutine stops running)
		}

	}
}
