/*
In Go language, a channel is a medium through which a goroutine communicates with another goroutine and this communication is lock-free.
Or in other words, a channel is a technique which allows to let one goroutine to send data to another goroutine.
*/

/* Video : https://youtu.be/i5HEE5Ikv3w */

package main

import (
	"fmt"
	"sync"
)

func main() {
	myChannel := make(chan int, 2) /* Create a channel of integers. 2 is the buffer value */
	wg := &sync.WaitGroup{}        /* Create a wait group */

	/* myChannel <- 5 */ /* Send 5 to the channel */
	/* fmt.Println(<-myChannel) */ /* Print the received value */
	/* Above 2 statements will give error because we can only write in channel when someone is listening to it, otherwise it will cause a deadlock */

	/* Add a goroutine to the wait group */
	wg.Add(2)
	/* <-chan : Receive Only */
	go func(myChannel <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()

		val, isChannelOpen := <-myChannel /* Receive a value from the channel */
		if isChannelOpen {
			fmt.Println(val) /* Print the received value */
		} else {
			fmt.Println("Channel is closed")
		}
	}(myChannel, wg)
	/* chan<- : Send Only */
	go func(myChannel chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		myChannel <- 5   /* Send 5 to the channel */
		myChannel <- 10  /* Send 10 to the channel */
		close(myChannel) /* Close the channel */
	}(myChannel, wg)
	/* Above statements work in passing and receiving the value as both goroutine are running in parallel */

	wg.Wait() /* Wait for the goroutine to finish */
}
