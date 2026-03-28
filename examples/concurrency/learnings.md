## Channels
Channels are a powerful concurrency primitive in Go that allow goroutines to communicate and synchronize with each other. They provide a way to send and receive values between goroutines, enabling safe and efficient communication without the need for explicit locks or shared memory.

### Unbuffered Channels
An unbuffered channel is a channel that has no capacity to hold values. When a goroutine sends a value to an unbuffered channel, it will block that goroutine until another goroutine receives that value from the channel. Similarly, when a goroutine receives a value from an unbuffered channel, it will block that goroutine until another goroutine sends a value to the channel. This synchronization mechanism ensures that the sender and receiver are coordinated, making it easier to manage concurrent operations.

### Buffered Channels
A buffered channel is a channel that has a specified capacity to hold values. When a goroutine sends a value to a buffered channel, it will block that goroutine only if the channel is full (i.e., it has reached its capacity). If there is space in the channel, the send operation will proceed without blocking. Similarly, when a goroutine receives a value from a buffered channel, it will block that goroutine only if the channel is empty. Buffered channels can help improve performance by allowing goroutines to send and receive values without blocking as long as there is available capacity in the channel.



## Deadlock vs  Race Condition

### Deadlock
Everyone is waiting, nobody can proceed.
All goroutines are blocked waiting on each other in a circular dependency. The program freezes completely and never makes progress. For example, if goroutine A holds a lock that goroutine B needs, and goroutine B holds a lock that goroutine A needs, both goroutines will be stuck waiting for each other indefinitely, resulting in a deadlock.
### Race Condition
Everyone is running, but in the wrong order.
Multiple goroutines access shared state concurrently and the outcome depends on who wins the CPU scheduling lottery. The program keeps running but produces unpredictable or incorrect results. For example, if two goroutines are incrementing the same counter variable without proper synchronization, the final value of the counter may be incorrect due to race conditions.

