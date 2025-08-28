package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, errRevenue := getUserInput("Revenue: ")

	if errRevenue != nil {
		fmt.Println("Error: ", errRevenue)
		return
	}

	expenses, errExpenses := getUserInput("Expenses: ")

	if errExpenses != nil {
		fmt.Println("Error: ", errExpenses)
		return
	}

	taxRate, errRate := getUserInput("Tax Rate: ")

	if errRate != nil {
		fmt.Println("Error: ", errRate)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be positive number")
	}

	return userInput, nil
}
