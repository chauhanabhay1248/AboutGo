package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World !!")
	investment_calculator_with_variables()
	profile_calculator_with_user_input()
}

func investment_calculator_with_variables() {
	// var invAmount, yr float64 = 1000, 10
	// invAmount, yr := 1000.0, 10.0

	const inflationRate = 2.5
	var invAmount float64 = 1000
	expRate := 5.5
	yr := 10.0

	futureValue := invAmount * math.Pow(1+expRate/100, yr)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, yr)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func profile_calculator_with_user_input() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	printEbt := fmt.Sprintf("ebt: %f\n", ebt)
	fmt.Print(printEbt)
	fmt.Printf("Profit: %v\n", profit)
	fmt.Printf("Ratio: %f\n", ratio)
}
