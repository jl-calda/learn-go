package main

import (
	"errors"
	"fmt"
)

func showErrorMessage(err error) {
	fmt.Println("-------ERROR-------")
	fmt.Println(err)
	fmt.Println("-------------------")
}

func profitCalc() {
	revenue, err1 := getUserInput("Enter the Revenue:")
	if err1 != nil {
		showErrorMessage(err1)
		return
	}

	expenses, err2 := getUserInput("Enter the Expenses:")
	if err2 != nil {
		showErrorMessage(err2)
		return
	}

	taxRate, err3 := getUserInput("Enter the Tax Rate:")
	if err3 != nil {
		showErrorMessage(err3)
		return
	}

	ebt := getEBT(revenue, expenses)
	profit := getProfit(ebt, taxRate)
	ratio := getRatio(ebt, profit)

	fmt.Printf("Earnings Before Tax: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.2gf\n", ratio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Invalid input")
	}

	return userInput, nil
}

func getEBT(revenue float64, expenses float64) float64 {
	return revenue - expenses
}

func getProfit(ebt float64, taxRate float64) float64 {
	return ebt * (1 - taxRate/100)
}

func getRatio(ebt float64, profit float64) float64 {
	return ebt / profit
}
