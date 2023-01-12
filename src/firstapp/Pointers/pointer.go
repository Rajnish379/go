package main

import (
	"fmt"
)

func main(){
	var a int = 42
	var b *int = &a
	// Here &a indicates the address of the memory location for variable a
	// Adding * to b will deference the pointer b and prints out the value of the variable at it's memory location
	fmt.Println(a,*b)
	a = 27
	fmt.Println(a,*b)
	*b = 26
	fmt.Println(a,*b)
	c := [3]int{1,2,3}
	d := &c[0]
	e := &c[1]

	// The addresses of d and e are 4 bytes apart because they are integer values and one integer value needs 4 bytes of space
	// Cannot do arithmetic operations between pointer addresses and integers 
	fmt.Println("%v %p %p\n",c,d,e)

	var ms *myStruct 
	ms = &myStruct{foo:42}
	fmt.Println(ms)

	var ns *myStruct
	// Uninitialized pointer will be initialized by go with the value of nil
	fmt.Println(ns)
	ns  =new(myStruct)
	//We have to add parentheses because in go dereferencing operator has lower preference than the dot operator
	// So we have to make sure that we dereference it in order to access the value there
	// (*ns).foo statement is similar to the one below because the compiler sees it in the same way
	ns.foo = 42
	fmt.Println(ns.foo)
	fmt.Println(ns)

	// In arrays the values of the array are considered intrinsic to the variable
	arr := [3]int{1,2,3}
	arr2 := arr
	fmt.Println(arr,arr2)
	arr[1] = 42
	fmt.Println(arr,arr2)

	// But slice is actually the projection of the underlying array
	// Slice contains the pointer to the first element that the slice is pointing to on the underlying array
	s := []int{1,2,3}
	t := s
	fmt.Println(s,t)
	s[1] = 42
	fmt.Println(s,t)

	// Maps also represent the underlying data similar to slices
	m := map[string]string{"foo":"bar","baz":"buz"}
	n := m
	fmt.Println(m,n)
	m["foo"] = "qux"
	fmt.Println(m,n)

}

type myStruct struct {
	foo int
}