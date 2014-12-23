package main

import (
	"fmt"
	"time"
	"math/rand"
)

func boring(msg string, c chan string) {
	i := 0
	for {
		// "c <- " - this is blocking, boring waits for a receiver to be ready
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		i++
	}
}

/*
	A channel connects the main and boring goroutines so they can communicate. 
*/
func main() {
	c := make(chan string)

	go boring("using channels", c)

	/*
		at this moment boring method is blocked on line 13 before time.Sleep call
		and waits for a receiver to be ready.
	*/
	time.Sleep(3 * time.Second)

	for i := 0; i < 5; i++ {
		/* "<-c": here main function is waiting for a value to be sent. (blocking) */
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("Still boring ... leaving!")
}