package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBabers  int
	BabersDoneChan  chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(Barber string) {
	//add a barber to the shop evrytime the func is callled
	shop.NumberOfBabers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to waiting room to check for clients ", Barber)

		for {
			//if there are no clients , the barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				color.Yellow("There are no clients, so %s takes a nap", Barber)
				isSleeping = true
			}
			//keep listening to the channel
			//get client from room when someone arrives
			client, shopOpen := <-shop.ClientsChan
			//shopOpen returns a bool value(standard way to check if the value received from the channel was sent to the channel)
			if shopOpen {
				if isSleeping {
					//customer wakes baber up if hes sleeping
					color.Yellow("%s wakes %s up", client, Barber)
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

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is finished cutting %s's hair.", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home ", barber)
	shop.BabersDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	color.Cyan("Closing shop for the day")

	close(shop.ClientsChan)
	shop.Open = false

	//wait until all barbers are done
	//blocks until every single barber is done
	for a := 1; a <= shop.NumberOfBabers; a++ {
		<-shop.BabersDoneChan
	}
	close(shop.BabersDoneChan)

	color.Green("Barbers shop closed for the day , everyone has gone home")

}

func (shop *BarberShop) addClient(client string) {
	color.Green("%s arrives", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%s takes a seat in the waiting room", client)
		default: //default avoid deadlock (when a  goroutine stops running)
			color.Red("The waiting rom is full, so %s leaves", client)
		}
	} else {
		color.Red("The Shop is already closed, so %s leaves!", client)

	}
}
