/*
Problem 1: Simple producer consumer
Create a program with one goroutine that sends the numbers 1–10 to a channel,
and the main goroutine receives and prints them. Use a buffered channel of size 3.
Make sure main doesn't exit early.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println(time.Now(), "sending Order ", i)
			ch <- i
			fmt.Println(time.Now(), "sent Order ", i)
		}
		close(ch)
	}()

	fmt.Println(time.Now(), "Waiting for all orders")

	for receiver := range ch { // blocks main go routine until ch channel closed
		fmt.Println(time.Now(), "Received: Order ", receiver)
	}

	fmt.Println(time.Now(), "Received all orders")
}
