package main

// Enter two numbers swap them numbers without using third variable.

import (
	"fmt"
)

// Function to Ask for a number
func getNumber(prompt string) int {

	// Print the prompt to the user
	var number int
	fmt.Print(prompt)

	// Collect the number
	fmt.Scanln(&number)

	return number
}

func swap(firstNumber, secondNumber int) (int, int) {
	return secondNumber, firstNumber
}

func main() {

	var firstNumber, secondNumber int

	firstNumber = getNumber("Enter first number : ")
	secondNumber = getNumber("Enter second number : ")
	fmt.Println("First number is ", firstNumber)
	fmt.Println("Second number is ", secondNumber)

	firstNumber, secondNumber = swap(firstNumber, secondNumber)
	fmt.Println("Now First number is ", firstNumber)
	fmt.Println("Now Second number is ", secondNumber)

}
