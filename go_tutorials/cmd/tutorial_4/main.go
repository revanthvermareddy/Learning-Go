package main

import (
	"errors"
	"fmt"
)

// arrays

func main() {
	var printValue = "Hello World!"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result int = intDivision(numerator, denominator)
	fmt.Println(result)

	// we can capture multiple return values like below
	// and with Printf() we can format the string and by using %v we can replace them with variables passed to the string
	var result2, remainder, err = intDivisionWithRemainder(numerator, denominator)

	// if err != nil {
	// 	fmt.Print(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of the integer division is %v", remainder)
	// } else {
	// 	fmt.Printf("The result of the integer division is %v with remainder %v", result2, remainder)
	// }

	// The switch statement equivalent for the above if else statements and we do not need to mention explicit break statements (it's implicit)

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", remainder)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v", result2, remainder)
	}

	// below is the conditional switch statement which looks at the variable

	switch remainder {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("The division was close")
	default:
		fmt.Println("The division was not close")
	}

	// && - is an d operator
	// || - is or operator

}

// this is another function declaration
func printMe(printValue string) {
	fmt.Println(printValue)
}

// function to do the int division and return result. Here we specify the return type of int
func intDivision(numerator int, denominator int) int {
	var result int = numerator / denominator
	return result
}

// function to do the int division and return both result and remainder. Here we specify multiple return types using (int, int)
// a design pattern in Go is when a func is expected to throw and error we return an variable of error type along with the other variables
func intDivisionWithRemainder(numerator int, denominator int) (int, int, error) {
	var err error // default value is nil
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
