package main

import (
	"fmt"
)

// Go's functions start with the func keyword
// Followed by the name of the function
// We have to put the opening curly brace of the function in the same line as the func keyword strictly in go unlike other programming languages
func main(){
	fmt.Println("Hello, playground")
	for i:=0;i<5;i++ {
		sayMessage("Hello Go!",i)
	}
	greeting := "Hello"
	name := "Stacey"
	sayGreeting(greeting,name)
	fmt.Println(name)
	sayGreetingPointer(&greeting,&name)
	fmt.Println(name)
	s := sum(1,2,3,4,0)
	fmt.Println("The sum is ",*s)
	s2 := sum_with_named_return_value(1,2,3,4,5)
	fmt.Println("The sum is ",s2)

	d,err := divide(5.0,1.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
	// Anonymous function i.e., function without names 
	// If we don't call this function immediately, the compiler will give an error saying that we declared the function but didn't use it anywhere
	for i:=0;i<5;i++{
	func(i int) {
		fmt.Println(i)
	}(i) // Here we are spoon feeding the i value directly into the function as an argument instead of referring it from the scope of the loop to avoid any unprecedented errors
	 } // These two parantheses execute the anonymous function on the fly
	
	 // Type signature for an anonymous function variable is simply a func()
	 // Make sure that you already define the anonymous function variable before executing it
	 var f func() = func() {
		fmt.Println("Hello Go!")
	 }
	 f()
	 g := greeter {
		greeting: "Hello",
		name: "Go",
	 }
	 g.greet()
	 fmt.Println("The new name is:",g.name)

}

func sayMessage(msg string, idx int){
	fmt.Println(msg)
	fmt.Println("The value of the index is",idx)
}
// To remove verbosity from here, we can define the variables first and the type of those multiple variables at the end
// The compiler will automatically infer that every variable that's in that comma delimited list has the same type
func sayGreeting(greeting, name string){
	fmt.Println(greeting,name)
	// Changing the name over here doesn't change the name variable in the main function
	// Because when the value is actually passed to this function, it is passed a copy of the variable but not the original variable itself
	name = "Ted"
	fmt.Println(name)
}

// But with the help of pointers we can change the value of the variable in main
func sayGreetingPointer(greeting,name *string){
	fmt.Println(*greeting,*name)
	*name = "Ted"
	fmt.Println(*name)
}
// The spread operator before int tells go to take up all the last arguments that are passed in the caller and wrap them up into a slice that has the name of the variable
// If we are using variadic parameters along with some additional parameters, then the variadic parameter has to be at the end cause go compiler can't recognize the difference between normal parameters and variadic parameter unless u specify it at the end
// When we are returning the result variable from the local stack,go automatically promotes the variable to be available on the shared memory in computer also known as heap memory
func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _,v := range values {
		result += v
	}
	return &result
}
// Syntactic sugar for declaring a result variable
func sum_with_named_return_value(values ...int)(result int){
	fmt.Println(values)
	for _,v := range values {
		result += v
	}
	return
}

func divide(a,b float64) (float64,error) {
	if b == 0.0{
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b,nil
}

type greeter struct {
	greeting string
	name string
}

// Methods are working with unknown types or custom types
func (g greeter) greet() {
	fmt.Println(g.greeting,g.name)
	// This will not take effect because the greeter object we are receiving here is just a copy of the main greeter object present in main function
	// Can manipulate the g value if u pass in a pointer receiver there i.e., g *greeter in its paramteters
	g.name = ""
}