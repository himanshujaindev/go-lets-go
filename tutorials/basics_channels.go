package main

import (
	"fmt"
	"time"
)

// channels + goroutines

func basics_channels() {
	var c = make(chan int)

	// c <- 1 // c: [1] // unbuffered channel with only one int value

	// // when writing to an unbuffered channel, code will block until read happens

	// var i = <-c // c: [] (1 is popped out)
	// fmt.Println(i)

	go process(c)
	// fmt.Println(<-c) // waiting ... until the value is set in the process fn

	for i := range c {
		fmt.Println(i)
	}

	var b = make(chan int, 5) // buffer chan
	go process(b)
	for i := range b { // process fn will exit and not wait until main fn is completed
		fmt.Println(i)
		time.Sleep(time.Second + 1) // some work ...
	}
}

func process(channel chan int) {
	// c <- 123
	defer close(channel) // else deadlock! ; defer = before fn exits
	for i := 0; i < 5; i++ {
		channel <- i
	}
	fmt.Println("Exiting process")
}
