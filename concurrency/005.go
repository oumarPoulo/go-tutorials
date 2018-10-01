package main

import (
	"fmt"
)

// Channels can be thought as pipes using which Goroutines communicate.
// Similar to how water flows from one end to another in a pipe,
// data can be sent from one end and received from the another end using channels.
// The syntax to send and receive data from a channel are given below
// data := <- a // read from channel a
// a <- data // write to channel a
// Sends and receives to a channel are blocking by default.
// What does this mean? When a data is sent to a channel,
// the control is blocked in the send statement until some other Goroutine reads from that channel.
// Similarly when data is read from a channel, the read is blocked until some Goroutine writes data to that channel.
// This property of channels is what helps Goroutines communicate effectively without the use of explicit locks
// or conditional variables that are quite common in other programming languages.
// Buffered Channels : c := make(chan int, 1)
// This creates a buffered channel with a capacity of 1.
// Normally channels are synchronous; both sides of the channel will wait until the other side is ready.
// A buffered channel is asynchronous; sending or receiving a message will not wait unless the channel is already full.
// Closing channels and for range loops on channels
// Receivers can use an additional variable while receiving data from the channel to check whether the channel has been closed.
func main() {
	producerReceiver1()
	fmt.Println("With for loop on channel msg")
	producerReceiver2()
	var input string
	fmt.Scanln(&input)
}

// Senders have the ability to close the channel to notify receivers that no more data will be sent on the channel.
func producer(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

// The function has an infinite for loop which checks whether the channel is closed using the variable ok.
// If ok is false it means that the channel is closed and hence the loop is broken
func producerReceiver1() {
	c := make(chan int)
	go producer(c)
	for {
		v, ok := <-c
		if !ok {
			break
		}
		fmt.Println("Received", v, ok)
	}
}

// The for range loop receives data from the ch channel until it is closed. Once ch is closed, the loop automatically exits
func producerReceiver2() {
	c := make(chan int)
	go producer(c)
	for v := range c {
		fmt.Println("Received", v)
	}
}

// Received 0 true
// Received 1 true
// Received 2 true
// Received 3 true
// Received 4 true
// Received 5 true
// Received 6 true
// Received 7 true
// Received 8 true
// Received 9 true
// With for loop on channel msg
// Received 0
// Received 1
// Received 2
// Received 3
// Received 4
// Received 5
// Received 6
// Received 7
// Received 8
// Received 9
