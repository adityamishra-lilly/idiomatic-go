package main

import (
	"fmt"
	"time"
)

/*
Deadlock:
2026-03-28 10:46:36.5002233 +0000 GMT m=+0.000547601 0 sending
2026-03-28 10:46:36.5024455 +0000 GMT m=+0.002769801 0 sent
2026-03-28 10:46:36.5024455 +0000 GMT m=+0.002769801 1 sending
2026-03-28 10:46:36.5024455 +0000 GMT m=+0.002769801 1 sent
2026-03-28 10:46:36.5024455 +0000 GMT m=+0.002769801 2 sending
2026-03-28 10:46:38.5005054 +0000 GMT m=+2.000829701 waiting for messages
2026-03-28 10:46:38.5005054 +0000 GMT m=+2.000829701 received 0
2026-03-28 10:46:38.5005054 +0000 GMT m=+2.000829701 received 1
2026-03-28 10:46:38.5005054 +0000 GMT m=+2.000829701 received 2
2026-03-28 10:46:38.5005054 +0000 GMT m=+2.000829701 exiting
*/
// func main() {
// 	ch := make(chan int, 2)

// 	go func() {
// 		for i := 0; i < 3; i++ {
// 			fmt.Println(time.Now(), i, "sending")
// 			ch <- i
// 			fmt.Println(time.Now(), i, "sent")
// 		}

// 		fmt.Println(time.Now(), "All completed")
// 	}()

// 	time.Sleep(2 * time.Second)

// 	fmt.Println(time.Now(), "waiting for messages")

// 	fmt.Println(time.Now(), "received", <-ch)

// 	fmt.Println(time.Now(), "received", <-ch)

// 	fmt.Println(time.Now(), "received", <-ch)

// 	fmt.Println(time.Now(), "exiting")

// }

func main() {
	ch := make(chan int, 2)
	exit := make(chan struct{})

	go func() {
		for i := 1; i < 5; i++ {
			fmt.Println(time.Now(), i, "sending")
			ch <- i
			fmt.Println(time.Now(), i, "sent")
			time.Sleep(1 * time.Second)
		}

		fmt.Println(time.Now(), "All completed, leaving")
		close(ch) // ← critical: signals "no more values coming"
	}()

	go func() {
		/*
			"Select" is effective when using multiple channels
		*/
		// for{
		// 	select {
		// 	case v, open := <- ch:
		// 		if !open {
		// 			close(exit)
		// 			return
		// 		}
		// 		fmt.Println(time.Now(), "received", v)
		// 	}
		// }

		// In cases where only one channel is used
		for v := range ch { // keeps receiving until ch is closed
			fmt.Println(time.Now(), "received", v)
		}

		close(exit) // ← signals main that everything is done

	}()

	fmt.Println(time.Now(), "waiting for full completion")
	<-exit
	fmt.Println(time.Now(), "exiting")
}
