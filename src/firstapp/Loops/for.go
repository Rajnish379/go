package main 

import (
	"fmt"
)

func main() {
	// Comma operator is not available in go
	// In go the increment option is a statement but not an expression
	// Can't do i,j = i++,j++ because of above comment
	for i,j:=0,0;i<5;i,j=i+1,j+2 {
		fmt.Println(i,j)
	}
	for i:=0;i<5;i++ {
		fmt.Println(i)
		//Bad practice to manipulate the counter variable inside a loop
		//But here these lines indicate that go allows us to do so
		if i%2 == 0{
			i /= 2
		} else {
			i = 2*i + 1
		}
	// Here i is scoped to the main block instead of the for loop 
	i := 0
	// Can't leave the first semicolon in for loop out  cause go will think that it doesn't have an initializer if we do so
	for ; i<5;i++{
		fmt.Println(i)
	}
	// Failed i of for loop gets printed here i.e., 5
	fmt.Println(i)
	}
	// Infinite loop example 
	// for j:= 0;j<5; {
	// 	fmt.Println(j)
	// }
	// Go doesn't have do while loops but we can have the same functionality of the do while if u look at code below
	for j:= 0;j<5; {
		fmt.Println(j)
		j++
	}
	// Special case of for loop by initializing in main block and only specifying the terminating condition without specifying the initializer and the increment condition in for loop 
	k:=0
	for k<7 {
		fmt.Println(k)
		k++
	}
	l := 0
	for {
		fmt.Println(l)
		l++
		if l == 4 {
			break
		}
	}
	for m:= 0;m < 10;m++ {
		if m%2 == 0 {
			continue
		}
		fmt.Println(m)
	}
	// Nested for loop 
// Label helps us to tell the break statement in the inner nested loop to tell where it should break to 
Loop:
	for p := 1; p<=3; p++ {
		for q := 1; q<=3; q++ {
			fmt.Println(p*q)
			if p*q >= 3 {
				// This break doesn't break out of the entire nested loop. INstead it only breaks out of the only closest loop it finds i.e., the one having q only and the p loop just restarts 
				break Loop
			}
		}
	}
	// This for looping also works for arrays, not only slices
	s := []int{1,2,3}
	for k,v := range s {
		fmt.Println(k,v)
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
	for k,v := range statePopulations {
		fmt.Println(k,v)
	}
	// Printing out each individual characeter of string
	str := "Hello Go!"
	for k,v:= range str {
		fmt.Println(k,string(v))
	}
	// If we don't want to use the key in a map in go
	// _ gives hint to compiler that we don't want to use key
	for _,v := range statePopulations {
		fmt.Println(v)
	}
}