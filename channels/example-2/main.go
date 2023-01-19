package main

import "fmt"

func Unbuffered() {
	datachan := make(chan int)
	go func() {
		datachan <- 234 //send data into chan
		close(datachan)
	}()
	val := <-datachan //receives data from the channel and stores in a var
	fmt.Printf("val=%d\v", val)

}

func Buffered() {

	bufchan := make(chan int, 2)
	go func() {
		for i := 0; i < 200; i++ {
			//the buffer size are 2 values sent to the channel
			//once a value is received fro the chan the buff reduces to 2
			fmt.Println("The buffer Size is :", len(bufchan))
			bufchan <- i
		}
		close(bufchan)
	}()
	//receiving data from the channel
	for val := range bufchan {
		fmt.Println("The value received are", val)
	}

}

func main() {
	Unbuffered()
	Buffered()
}
