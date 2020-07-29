// Write a program which prompts the user to enter a floating point number and prints
//  the integer which is a truncated version of the floating point number that was entered.
//  Truncation is the process of removing the digits to the right of the decimal place.

package main

import (
	"fmt"
	"math"
)

func main() {
	var truncNum float64
	var number float64

	fmt.Scan(&number)

	truncNum = math.Trunc(number)

	fmt.Printf("%f\n", truncNum)
}
