package main

import "time"

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBabers  int
	BabersDoneChan  chan bool
	ClientsChan     chan string
	Open            bool
}
