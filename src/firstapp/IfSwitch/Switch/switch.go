package main

import (
	"fmt"
)

func main() {
	// We can't fall through in go but we can add multiple conditions in a single case
	// Cannot have overlapping cases when u r using the multiple condition syntax in go
	// Switch with initializer
	switch i:=2+3;i {
	case 1,5,10:
		fmt.Println("one, five or ten")
	case 2,4,6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	j:=10
	// Tagless switch
	// It can overlap unlike the tagged switch
	// If the cases overlap then only the first case which matches the condition will execute
	// No need of curly braces for every case block because go assumes by default that all the statements under a case statement are considered as part of that particular case block 
	// Unlike other programming languages, in go, break statement in switch blocks is already implied by default
	// Have to implicitly specify fallthrough if u want to execute multiple caseblocks
	// fallthrough is logicless i.e., even if the condition is false in the next block of fallthrough, the case block still executes
	// Dev's responsibility to take care of fallthrough keyword
	switch {
	case j <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough
	
	case j<=20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		break
		fmt.Println("This will print too")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	case [2]int:
		fmt.Println("i is [2]int")
	default:
		fmt.Println("i is another type")
	}

}