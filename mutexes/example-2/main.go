package main

import (
	"fmt"
	"sync"
)

//write a program that projects how much money someome will make in 52 weeks (1 year )

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	//variable for bank balance
	var bankBalance int

	var balance sync.Mutex

	//print out starting values
	fmt.Printf("Initial account balance: $%d.00", bankBalance)
	fmt.Println()

	//define weekly revenue
	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Crypto", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}
	wg.Add(len(incomes))

	//loop through 52 weeks and print out how much; keep a runnig total
	for i, income := range incomes { ///looping through income with the range of incomes

		go func(i int, income Income) {
			defer wg.Done()
			//loop through all 52 weeks
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance   //current bank balance
				temp += income.Amount //add to the current income for the current week
				bankBalance = temp    //add temp to the bank balance
				balance.Unlock()
				fmt.Printf("On Week %d, you earned $%d.00 from %s\n", week,
					income.Amount, income.Source)
			}
		}(i, income)
	}
	wg.Wait()

	//print final balance
	fmt.Printf("Final Bank Balance: $%d.00", bankBalance)
	fmt.Println()
}
