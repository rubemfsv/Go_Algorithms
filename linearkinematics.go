// Let us assume the following formula for displacement s as a function of time t, acceleration a, 
// initial velocity vo, and initial displacement so.

// s =½ a t2 + vot + so

// Write a program which first prompts the user to enter values for acceleration, initial velocity, and
// initial displacement. Then the program should prompt the user to enter a value for time and the
// program should compute the displacement after the entered time.

// You will need to define and use a function called GenDisplaceFn() which takes three float64 
// arguments, acceleration a, initial velocity vo, and initial displacement so. GenDisplaceFn() should 
// return a function which computes displacement as a function of time, assuming the given values
// acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn()
// should take one float64 argument t, representing time, and return one float64 argument which is 
// the displacement travelled after time t.

// For example, let’s say that I want to assume the following values for acceleration, initial velocity,
// and initial displacement: a = 10, vo = 2, so = 1. I can use the following statement to call 
// GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//GenDisplaceFn definition
func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return 0.5*acceleration*math.Pow(time, 2) + initialVelocity*time + initialDisplacement
	}

	return fn
}

func main() {
	//var acceleration, initialVelocity, initialDisplacement, time float64

	for {
		fmt.Println("please input the acceleration, initialVelocity, initialDisplacement, time separated by space")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		input = strings.Trim(input, " ")

		fmt.Printf("input is: [%s]\n", input)

		acceleration, _ := strconv.ParseFloat(strings.Split(input, " ")[0], 64)
		initialVelocity, _ := strconv.ParseFloat(strings.Split(input, " ")[1], 64)
		initialDisplacement, _ := strconv.ParseFloat(strings.Split(input, " ")[2], 64)
		time, _ := strconv.ParseFloat(strings.Split(input, " ")[3], 64)

		fmt.Printf("acceleration: %f\n", acceleration)
		fmt.Printf("initialVelocity: %f\n", initialVelocity)
		fmt.Printf("initialDisplacement: %f\n", initialDisplacement)
		fmt.Printf("time: %f\n", time)

		fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

		//fmt.Println(fn(3))
		fmt.Println(fn(time))
	}
}