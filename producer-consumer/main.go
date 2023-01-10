package main

import (
	"math/rand"
	"time"
)

// /no of Pizzas we've  made
const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error //when we've finished making pizzas
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano()) //it ensures we dont get the same result when we run the program

	// print out a message

	// create a producer

	// run the producer in the background

	// create and run consumer

	// print out the ending message
}
