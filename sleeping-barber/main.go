package main

import (
	"fmt"
	"math/rand"
	"time"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBabers  int
	BabersDoneChan  chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) cutHair(barber, client string) {
	fmt.Printf("%s is cutting %s's hair \n", barber, client)
	time.Sleep(shop.HairCutDuration)
	fmt.Printf("%s is finished cutting %s's hair. \n", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	fmt.Printf("%s is going home \n", barber)
	shop.BabersDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	fmt.Printf("Closing shop for the day \n")

	close(shop.ClientsChan)
	shop.Open = false

	//wait until all barbers are done
	//blocks until every single barber is done
	for a := 1; a <= shop.NumberOfBabers; a++ {
		<-shop.BabersDoneChan
	}
	close(shop.BabersDoneChan)

	fmt.Printf("Barbers shop closed for the day , everyone has gone home \n")

}

func (shop *BarberShop) addBarber(Barber string) {
	//add a barber to the shop evrytime the func is callled
	shop.NumberOfBabers++

	go func() {
		isSleeping := false
		fmt.Printf("%s goes to waiting room to check for clients \n", Barber)

		for {
			//if there are no clients , the barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				fmt.Printf("There are no clients, so %s takes a nap \n", Barber)
				isSleeping = true
			}
			//keep listening to the channel
			//get client from room when someone arrives
			client, shopOpen := <-shop.ClientsChan
			//shopOpen returns a bool value(standard way to check if the value received from the channel was sent to the channel)
			if shopOpen {
				if isSleeping {
					//customer wakes baber up if hes sleeping
					fmt.Printf("%s wakes %s up \n", client, Barber)
					isSleeping = false
				}
				//cut hair
				shop.cutHair(Barber, client)
			} else {
				//shop is closed, so send the barber Home and close th goroutine
				shop.sendBarberHome(Barber)
				return //closes the goroutine
			}

		}

	}()
}

func (shop *BarberShop) addClient(client string) {
	fmt.Printf("%s arrives", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			fmt.Printf("%s takes a seat in the waiting room \n", client)
		default: //default avoid deadlock (when a  goroutine stops running)
			fmt.Printf("The waiting rom is full, so %s leaves \n", client)
		}
	} else {
		fmt.Printf("The Shop is already closed, so %s leaves! \n", client)

	}
}

//Testing things out
func main() {

	//random number generator used with the arrivalRate so clients
	// dont arrive at the same interval
	rand.NewSource(time.Now().UnixNano())

	fmt.Printf("The Sleeping Barber Problem \n")

	//variable setup to define objects and  parameters
	seatingCapacity := 10
	arrivalRate := 100
	cutDuration := 20000 * time.Millisecond
	timeOpen := 10 * time.Second
	clientChan := make(chan string, seatingCapacity)
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
	fmt.Printf("The Shop is Open for the day! \n")
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
