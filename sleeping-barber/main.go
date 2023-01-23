package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// variables
var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// seed our random number generator used with the arrivalRate
	//so clients dont arrive at the same interval
	rand.Seed(time.Now().UnixNano())

	// print welcome message
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("-------------------------")

	// create channels if we need any
	clientChan := make(chan string, seatingCapacity) //a buffered channel
	doneChan := make(chan bool)
	// create the barbershop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBabers:  0,
		ClientsChan:     clientChan,
		BabersDoneChan:  doneChan,
		Open:            true,
	}
	color.Green("The Shop is Open for the day!")
	// add barbers
	shop.addBarber("Frank")
	shop.addBarber("Harry")
	shop.addBarber("Dennis")
	shop.addBarber("Hilary")

	// start the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		//make the goroutine open for a particular time
		<-time.After(timeOpen)
		shopClosing <- true
		//Closing the shop for the day
		shop.closeShopForDay()
		closed <- true

	}()

	// add clients
	i := 1

	go func() {
		for {
			//get a random number wit average arrival rate
			randomMilliSeconds := rand.Int() % (12 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliSeconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// block until the barbershop is closed
	<-closed
}
