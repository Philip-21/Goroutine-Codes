package main

import (
	"fmt"
	"sync"
)

//write a program that projects how much money someome will make in 52 weeks (1 year )

var wg sync.WaitGroup

type income struct {
	Source string
	Amount int
}

func main() {
	//variable for bank balance
	var bankBalance int

	//print out starting values
	fmt.Printf("Initial account balance: $%d.00", bankBalance)
	fmt.Println()

	//define weekly revenue
	incomes := []income{
		{Source: "Main job", Amount: 500},
		{Source: "Crypto", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	//loop through 52 weeks and print out how much; keep a runnig total

	//print final balance
}
