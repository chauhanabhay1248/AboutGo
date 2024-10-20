package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World !!")
	investment_calculator_with_variables()
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
