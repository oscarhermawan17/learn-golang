package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinalcials(revenue, expenses, taxRate)

	fmt.Printf("%v\n", ebt)
	fmt.Printf("%v\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinalcials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64{
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}