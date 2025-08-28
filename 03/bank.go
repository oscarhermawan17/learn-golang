package main

import "fmt"

func main() {
	var accountBalance = 1000.0

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

		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return // return the value, and looping would be stoped
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New Ammount: ", accountBalance)
		} else if choice == 3 {
			fmt.Print("Withdrawl ammount: ")
			var withdraw float64
			fmt.Scan(&withdraw)

			if withdraw <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				return
			}

			if withdraw > accountBalance {
				fmt.Println("Invalid amount")
				return
			}

			accountBalance -= withdraw
			fmt.Println("Balance updated! New Ammount: ", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			// return
			break // Just break the loop, after loop another code still running
		}
	}

	fmt.Println("Thank you for using Go Bank!")
	
}