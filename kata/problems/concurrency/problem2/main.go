/*
Problem 2: Sum aggregator
Create a goroutine that sends 5 numbers into a channel. The main goroutine should receive all of them
and print the total sum. No hardcoding the count — use channel closing to know when to stop.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)

	go func() {
		nums := [5]int{1, 2, 3, 4, 5}

		for _, num := range nums {
			fmt.Println(time.Now(), "Storing value: ", num)
			ch <- num
		}
		close(ch)

	}()

	fmt.Println(time.Now(), "Waiting for values")

	sum := 0
	for val := range ch { // blocks main go function
		sum += val
	}

	fmt.Println(time.Now(), "Total Sum: ", sum)
}
