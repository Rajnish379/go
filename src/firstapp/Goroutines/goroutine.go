package main

import (
	"fmt"
	"runtime"
	"sync"
	//"time"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	// Using goroutine to invoke a function by adding go keyword as prefix to the function call 
	// go sayHello()
	// time.Sleep(100* time.Millisecond)
	//var msg = "Hello"
	// wg.Add(1)
	//go func(msg string) {
		// Gppdbye gets printed out before Hello 
		// Race condition between main function and goroutine
		//fmt.Println(msg)
		// This Done call decrements the number of weight groups by one and when the number of weight groups becomes zero then the sync method understands that its time to wrap up the execution of the application
		//wg.Done()
	//}(msg) //Passing arguments to goroutines is a good way of passing data to the goroutines and it helps the goroutine to couple with the data in a better way
	//msg = "Goodbye"
	//wg.Wait()
	runtime.GOMAXPROCS(100)
	for i := 0; i<10;i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	//runtime.GOMAXPROCS(1) This will make the entire application single threaded i.e., there will not be parallelism at all but only concurrency
	fmt.Printf("Threads: %v",runtime.GOMAXPROCS(-1)) //GOMaxprocs -1 is basically not changing the number of threads i.e., it is notifying us on how many threads are available for us to execute
	wg.Wait()

}

func sayHello() {
	fmt.Printf("Hello #%v\n",counter)
	m.RUnlock()
	wg.Done()
}

func increment(){
	counter++
	m.Unlock()
	wg.Done()
}