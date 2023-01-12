package main

import (
	"fmt"
	"math"
)

func main() {
	// Single line if block is never evaluated unless u surround it with flower brackets
	if true {
		fmt.Println("The test is true")

	}

	statePopulations := map[string]int{
		"California": 11111,
		"Texas": 22222,
		"Florida": 33333,
		"New York": 44444,
		"Pennsylvania": 55555,
		"Illinois": 66666,
		"Ohio": 77777,
	}
	// The part before the semicolon here is called the initializer of the if block 
	// Here pop and ok variables are only defined in the scope of the if statement only
	// They cannot be accessed outside of if block
	if pop,ok := statePopulations["Florida"];ok{
		fmt.Println(pop)
	}

	number := 50
	guess := 30
	// Have to write else beside the ending bracket of the corresponding if 
	if guess < 1{
		fmt.Println("The guess must be greateer than 1!")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
	if guess < number {
		fmt.Println("Too low")
	}
	if guess > number {
		fmt.Println("Too high")
	}
	if guess == number {
		fmt.Println("You got it!")
	}
}
	fmt.Println(number<=guess, number>=guess, number!=guess)
	fmt.Println(!true)
	//myNum == math.Pow(math.Sqrt(myNum),2)
	//Doesn't work with 0.123 because when go is working with floating point numbers,
	// They are not considered as exact representations of decimal values and so we have to be careful while working with them
	myNum := 0.123
	// 0.001 is the error parameter here, the less it is the more accurate our result will be
	if math.Abs(myNum / math.Pow(math.Sqrt(myNum),2) - 1) < 0.001{
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
}
// Shortcircuiting is valid in go
func returnTrue() bool {
	fmt.Println("returning true")
	return true
}