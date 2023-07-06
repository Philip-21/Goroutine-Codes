package main

import "fmt"

/*send and receive values from channels*/

func Receiver(c <- chan int ){
	for val := range c{
		fmt.Println(val)
	}
}

func Sender ()<- chan int {
	c := make(chan int)

	go func (){
		for i :=0; i<10;i++{
			c <-i
		}
		close(c)
	}()
	return c
}


func main(){
	Sendvalue := Sender()
	Receiver(Sendvalue)
}