package main

import (
	"fmt"
	//"net/http"
	"log"
)

func main() {
	// a, b := 1,0
	// ans := a/b
	// fmt.Println(ans)
	// In the above code, as the compiler doesn't know how to divide any number with a zero, it enters into a panic state
	// fmt.Println("start")
	// Customizing the panic message
	// panic("something bad happened")
	// fmt.Println("end")
	// fmt.Println("Raine")
	// http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request) {
	// 	w.Write([]byte("Hello Go!"))
	// })
	// err := http.ListenAndServe(":8000",nil)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// Panics happen after the deferred statements are executed
	fmt.Println("start")
	panicker()
	fmt.Println("end")


}
// Function recovering from panic stops executing even though it recovered
// But functions that call the panicking function i.e., functions that are higher up in the stack (function which calls the recovered function) still continue to function
func panicker(){
	fmt.Println("about to panic")
	defer func(){
		// recover() tells us whether the application is panicking or not
		// It returns nil if the application is not panicking and returns the error i.e., the cause of panic if the app is panicking
		if err := recover(); err != nil {
			log.Println("Error:",err)
			// If you are trying to recover from panic and you think you can't handle it
			// Then you can rethrow that panic 
			panic(err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}