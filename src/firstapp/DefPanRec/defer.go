package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Defer functions are executed in LIFO order i.e., Last In and First Out
	// Last function to get deferred becomes the first function to execute
	defer fmt.Println("start")
	// Go calls the deferred functions at the end of execution of program
	// Defer moves the line after the main function but before the main function returned
	defer fmt.Println("middle")
	defer fmt.Println("end")

	res,err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	// If we don't use defer here before closing it will give an error that we are trying to close the response body before even reading it 
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s",robots)

	a := "start"
	// Here start is printed instead of what we expected i.e., end
	// When we defer a function, it takes the argument at the time, the defer is called but not at the time the called function is executed
	// Defer eagerly evaluates the variable
	defer fmt.Println(a)
	a = "end"
}