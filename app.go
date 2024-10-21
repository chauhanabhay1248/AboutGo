package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	pout("Hello World !!")
	pout(investment_calculator_with_variables())
	pout(profile_calculator_with_user_input())
	pout(bank_case_statement())
}

func investment_calculator_with_variables() string {
	// var invAmount, yr float64 = 1000, 10
	// invAmount, yr := 1000.0, 10.0

	const inflationRate = 2.5
	var invAmount float64 = 1000
	expRate := 5.5
	yr := 10.0

	futureValue := invAmount * math.Pow(1+expRate/100, yr)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, yr)
	pout(futureValue)
	pout(futureRealValue)
	return "Investment Calculator Done \n"
}

func profile_calculator_with_user_input() string {
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
	return "Profit Calculator Done \n"
}

func bank_case_statement() string {
	fmt.Println("Welcome to Bank !!")

	currentAmount, err := readFromFile()
	if err != nil {
		fmt.Println("Error: ", err)
		panic("Can't continue because of error !!")
	}
	for {
		fmt.Println("Enter Choice: ")
		fmt.Println("1 for Show Amount")
		fmt.Println("2 for Deposit Amount")
		fmt.Println("3 for Withdraw Amount")
		fmt.Println("4 for Exit")

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Current Amount is ", currentAmount)
		} else if choice == 2 {
			fmt.Println("Enter Amount To Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			currentAmount = currentAmount + depositAmount
			writeIntoFile(currentAmount)
			fmt.Println("Updated Amount is ", currentAmount)
		} else if choice == 3 {
			fmt.Println("Enter Amount To Withraw: ")
			var withrawAmount float64
			fmt.Scan(&withrawAmount)
			if currentAmount-withrawAmount < 0 {
				fmt.Println("Not Enough Balance.")
				continue
			} else {
				currentAmount = currentAmount - withrawAmount
				writeIntoFile(currentAmount)
			}
			fmt.Println("Updated Amount is ", currentAmount)
		} else {
			fmt.Println("Thank you !!")
			break
		}
	}
	return "Bank Done \n"
}

func pout(val interface{}) {
	fmt.Println(val)
}

const fileName = "data.txt"

func writeIntoFile(data float64) {
	stringData := fmt.Sprint((data))
	os.WriteFile(fileName, []byte(stringData), 0644)
}

func readFromFile() (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("Failed to fetch data from file")
	}
	stringData := string(data)
	floatData, err := strconv.ParseFloat(stringData, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse string")
	}
	return floatData, nil
}
