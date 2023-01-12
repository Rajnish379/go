package main

import (
	"fmt"
	"strconv"
)

var i int = 42

// Can declare multiple variables inside a var block to avoid using var keyword multiple times
var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

var (
	counter int = 0
)

func main() {
	// Block scope takes precedence over package scope
	var i int
	i = 45
	// i := 46
	// i = 46
	// A variable can be reassigned but not redeclared in a block scope
	var j int = 34
	// k can be assigned any type from int, float32, float64
	var k float32 = 67.5
	// l can only be an int or a float64 type because there is no explicit variable type declaration
	l := 66
	var m float32
	m = float32(i)

	var n string
	// We have to use Itoa here because the string conversion converts the integer to its corresponding ascii character
	n = string(j)
	fmt.Printf("%v,%T\n", n, n)
	n = strconv.Itoa(j)
	fmt.Printf("%v, %T\n", n, n)
	fmt.Println(i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Println(k)
	fmt.Println(l)
	// Can convert int to float32 but not vice versa
	fmt.Printf("%v ,%T", m, m)

}
