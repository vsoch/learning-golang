package main

import (
	"fmt"
)

func main() {

	fmt.Print("Enter a number to sum to: ")
	var number, sum int
	fmt.Scan(&number)

	for n := 1; n <= number; n++ {
		sum += n
	}

	fmt.Println("The sum is", sum)
}
