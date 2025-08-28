package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error){
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, errBalance := strconv.ParseFloat(balanceText, 64)

	if errBalance != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText  := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() { 
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("---------")
		// panic("Cant continue, sorry") // panic will stop the program, and we can get the line number that caused the error
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)


		switch choice {
			case 1:
				fmt.Println("Your balance is", accountBalance)

			case 2:
				fmt.Print("Your deposit: ")
				var depositAmount float64
				fmt.Scan(&depositAmount)

				if depositAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					continue
				}

				accountBalance += depositAmount
				fmt.Println("Balance updated! New Ammount: ", accountBalance)
				writeBalanceToFile(accountBalance)

			case 3:
				fmt.Print("Withdrawl ammount: ")
				var withdraw float64
				fmt.Scan(&withdraw)

				if withdraw <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					continue
				}

				if withdraw > accountBalance {
					fmt.Println("Invalid amount")
					continue
				}

				accountBalance -= withdraw
				fmt.Println("Balance updated! New Ammount: ", accountBalance)
				writeBalanceToFile(accountBalance)

			default: 
				fmt.Println("Goodbye!")
				return
		}
	}
}