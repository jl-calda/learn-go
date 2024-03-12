package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Error reading balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Error parsing balance")
	}

	return balance, nil
}

func writeBalance(balance float64) {
	balanceText := fmt.Sprintf("%.2f", balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func bank() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------------------")
		panic("Cannot continue")
	}

	fmt.Println("Welcome to the Bank!")
	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			var depositAmount float64

			fmt.Print("How much would you like to deposit?")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount. Please try again.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your new balance is: ", accountBalance)
			writeBalance(accountBalance)
		case 3:
			var withdrawAmount float64

			fmt.Print("How much would you like to withdraw?")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid withdraw amount. Please try again.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds. Please try again.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Your new balance is: ", accountBalance)
			writeBalance(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for visiting the Bank!")
			return
		}

		// 	if choice == 1 {
		// 		fmt.Println("Your balance is: ", accountBalance)
		// 	} else if choice == 2 {
		// 		var depositAmount float64
		// 		fmt.Print("How much would you like to deposit?")
		// 		fmt.Scan(&depositAmount)
		// 		if depositAmount <= 0 {
		// 			fmt.Println("Invalid deposit amount. Please try again.")
		// 			continue
		// 		}
		// 		accountBalance += depositAmount
		// 		fmt.Println("Your new balance is: ", accountBalance)
		// 		writeBalance(accountBalance)
		// 	} else if choice == 3 {
		// 		var withdrawAmount float64
		// 		fmt.Print("How much would you like to withdraw?")
		// 		fmt.Scan(&withdrawAmount)
		// 		if withdrawAmount <= 0 {
		// 			fmt.Println("Invalid withdraw amount. Please try again.")
		// 			continue
		// 		}
		// 		if withdrawAmount > accountBalance {
		// 			fmt.Println("Insufficient funds. Please try again.")
		// 			continue
		// 		}
		// 		accountBalance -= withdrawAmount
		// 		fmt.Println("Your new balance is: ", accountBalance)
		// 		writeBalance(accountBalance)
		// 	} else {
		// 		fmt.Println("Goodbye!")
		// 		break
		// 	}

	}

	fmt.Println("Thank you for visiting the Bank!")
}
