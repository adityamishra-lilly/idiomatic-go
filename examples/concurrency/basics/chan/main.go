package main

import (
	"fmt"
	"time"
)

/*
This program panics because there is no goroutine outside of `main` interacting with the `ch` channel:

	fatal error: all goroutines are asleep - deadlock!
*/

// func main() {
// 	ch := make(chan int)

// 	ch <- 10

// 	v := <-ch

// 	fmt.Println("received", v)
// }

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println(time.Now(), "taking a nap")

		time.Sleep(2 * time.Second)

		ch <- "hello"
	}()

	fmt.Println(time.Now(), "waiting for message")

	//blocks (waits) until some goroutine sends a value into ch
	v := <-ch

	fmt.Println(time.Now(), "received", v)
}
