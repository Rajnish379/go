package main

import (
	"fmt"
)

func main(){
	//[...] indicates the compiler to initialize the array with the size just enough to hold the number of elements present in the flower brackets
	grades := [...]int{90,80,70}
	fmt.Printf("Grade: %v\n",grades)
	var students [5]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Lisaj"
	students[2] = "Lisal"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[1])
	fmt.Printf("Number of Students: %v\n",len(students))

	//var identityMatrix [3][3]int = [3][3]int{ [3]int{1,0,0},[3]int{0,1,0},[3]int{0,0,1}}
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1,0,0}
	identityMatrix[1] = [3]int{0,1,0}
	identityMatrix[2] = [3]int{0,0,1}

	fmt.Println(identityMatrix)
	a := [...]int{1,2,3}
	// A copy of the array is made here i.e., b and a are two completely different arrays now
	b := a
	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)
	// c is pointing to the same array a, so whatever changes are reflected in array c will be reflected in the parent array a too 
	c := &a
	c[2] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

    // Slice declaration
	d := []int{1,2,3}
	// Slices are natural reference types, i.e., they refer to the same underlying data when initialized to a different variable
	e := d
	e[1] = 5
	fmt.Println(d)
	fmt.Println(e)
	fmt.Printf("Length: %v\n",len(d))
	fmt.Printf("Capacity: %v\n",cap(d))
	
	// This slicing syntax also works with arrays 
	// a1 := [...]int{1,2,3,4,5,6,7,8,9,10}
	a1 := []int{1,2,3,4,5,6,7,8,9,10}
	b1 := a1[:] // slice of all elements
	c1 := a1[3:] // slice from 4th element to end
	d1 := a1[:6] // slice first 6 elements
	e1 := a1[3:6] // slice the 4th, 5th and 6th elements
	// All 5 slices will change the value of the 5th index because they refer to same underlying data
	a1[5] = 42
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)
	fmt.Println(e1)
	makeSlice := make([]int,3,100)
	//Empty slice with length and capacity zero
	makeSlice1 := []int{}
	fmt.Println(makeSlice)
	fmt.Printf("Length: %v\n",len(makeSlice))
	fmt.Printf("Capacity: %v\n",cap(makeSlice))
	fmt.Println(makeSlice1)
	fmt.Printf("Length: %v\n",len(makeSlice1))
	fmt.Printf("Capacity: %v\n",cap(makeSlice1))
	makeSlice1 = append(makeSlice1, 1)
	fmt.Println(makeSlice1)
	fmt.Printf("Length: %v\n",len(makeSlice1))
	fmt.Printf("Capacity: %v\n",cap(makeSlice1))
	// makeSlice1 = append(makeSlice1,2,3,4,5)
	// Here the spread operator after the slice breaks the entire slice into it's individual components
	makeSlice1 = append(makeSlice1,[]int{2,3,4,5}...)
	fmt.Println(makeSlice1)
	fmt.Printf("Length: %v\n",len(makeSlice1))
	fmt.Printf("Capacity: %v\n",cap(makeSlice1))
	q1 := []int{1,2,3,4,5}
	fmt.Println(q1)
	q2 := append(q1[:2],q1[3:]...)
	fmt.Println(q2)
	fmt.Println(q1)






}