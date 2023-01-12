package main

import (
	"fmt"
	"sync"
	"time"
)
const (
	logInfo = "INFO"
	logWarning = "WARNING"
	logError = "ERROR"
)
type logEntry struct {
	time time.Time
	severity string
	message string
}

var logCh = make(chan logEntry, 50)
// These type of empty struct channels have no data stored in it because they are allocated zero memory by go
// But they can be used to signal the receiver that the data has arrived 
var doneCh = make(chan struct{})

var wg = sync.WaitGroup{}
// Channels are designed to synchronize data transmission between multiple goroutines
func main() {
	// Buffered channel by increasing capacity of the channel to 50
	// Unbuffered channel can only store 1 int at a time
	// ch := make(chan int,50)
	//for j:= 0;j<5;j++{
	// wg.Add(2)
	//Need to have the same amout of senders and receivers for channel otherwise execution fails
	// This function will only receive data from channel 
	// go func(ch <- chan int) {
		// Pulling data out of the channel because the arrow is pointing outwards from the channel
		// i := <- ch
		// fmt.Println(i)
		// i = <- ch
		// fmt.Println(i)
		// Instead of writing the receiving variables multiple times, we can use a loop to receive all the values at once
		// Problem with this is we are causing deadlock because the for loop doesn't know when to exit and it keeps on monitoring for values from channel
		// for i:= range ch {
		// 	fmt.Println(i)
		// // }
		// for {
		// 	if i,ok := <- ch; ok {
		// 		fmt.Println(i)
		// 	} else {
		// 		break
		// 	}
		// }
	// 	// ch <- 27
	// 	wg.Done()
	// }(ch)
	// // This function will only send data into the channel
	// go func(ch chan<- int) {
	// 	// Sending data into the channel because the arrow symbol is to the right side of the channel 
		
	// 	ch <- 42
	// 	ch <- 66
	// 	// Closing the channel will make the length of the channel finite i.e., it will intimate the channel that we are done working with it 
	// 	// While closing a channel, we have to absolutely make sure that we are really closing it
	// 	// App panics when we try to send a value to a closed channel
	// 	close(ch)
	// 	// fmt.Println(<-ch)
	// 	wg.Done()
	// }(ch)
	// //}
	// wg.Wait()
	go logger()
	logCh <- logEntry{time.Now(),logInfo,"App is starting"}

	logCh <- logEntry{time.Now(),logInfo,"App is shutting down"}
	time.Sleep(100*time.Millisecond)
	doneCh <- struct{}{}
}

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"),entry.severity,entry.message)
		case <-doneCh:
			break
		}
	}
}