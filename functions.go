package main

import (
	"fmt"
)

type transformFn func(int) int

func functionsGo() {
	// 1. functions as values
	numbers := []int{1, 2, 3, 4, 5}
	dob := createTransformer(2)
	trip := createTransformer(3)
	doubled := transformNumbers(&numbers, dob)
	tripled := transformNumbers(&numbers, trip)
	halfed := transformNumbers(&numbers, func(number int) int { return number / 2 })
	fmt.Println(numbers)
	fmt.Println(doubled)
	fmt.Println(tripled)
	fmt.Println(halfed)
}

// func doubleNumbers(numbers *[]int) []int {
// 	dNumbers := []int{}

// 	for _, value := range *numbers {
// 		dNumbers = append(dNumbers, value*2)
// 	}

// 	return dNumbers
// }

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}

	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformerFunction() func(int) int {
	return double
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func recurssion() {
	numbers := []int{2, 3, 4, 5}
	fmt.Println(factorial(5))
	fmt.Println(factorialRec(5))
	fmt.Println(variadic(10, 2, 3, 4, 5))
	fmt.Println(variadic(10, numbers...))

}

func factorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result *= i
	}
	return result
}

func factorialRec(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorialRec(number-1)
}

func variadic(starting int, number ...int) int {
	total := 0
	for _, value := range number {
		total += value * starting
	}
	return total
}
