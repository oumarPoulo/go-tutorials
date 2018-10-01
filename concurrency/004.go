package main

import (
	"fmt"
	"time"
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
func main() {
	pingPong()
	var input string
	fmt.Scanln(&input)
}

func pingPong() {
	c := make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)
}

func pinger(c chan string) {
	for {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1) // used to print value after each second
	}
}

// ping
// pong
// ping
// pong
// ping
// pong
// ping
// pong
// ping
