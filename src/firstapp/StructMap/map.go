package main

import (
	"fmt"
)

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 11111,
		"Texas": 22222,
		"Florida": 33333,
		"New York": 44444,
		"Pennsylvania": 55555,
		"Illinois": 66666,
		"Ohio": 77777,
	}
	m := map[[3]int]string{}
	// Order of elements is not maintained in map unlike arrays and slices

	statePopulations["Georgia"] = 88888
	fmt.Println(statePopulations,m)
	delete(statePopulations,"Georgia")
	// Even though Georgia was deleted from map it still shows no error and prints out the value of 0 which is inappropriate
	fmt.Println(statePopulations["Georgia"])
	// Use comma and ok syntax if u r unsure about whether the key we are trying to access is present in a map or not
	pop, ok := statePopulations["Oho"]
	fmt.Println(pop,ok)
	fmt.Println(statePopulations)
	fmt.Println(len(statePopulations))
	sp := statePopulations
	// Impacting map sp will have impact on the parent map as well i.e., statePopulations
	delete(sp,"Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)
}