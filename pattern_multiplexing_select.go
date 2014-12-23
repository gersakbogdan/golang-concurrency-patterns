package main

import (
	"fmt"
	"time"
	"math/rand"
)

func boring(msg string, delay int) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3 * delay)) * time.Millisecond)
		}
	}()

	return c
}

/*
	The select statement provides another way to handle multiple channels.
	It's like a switch, but each case is a communication:
	- All channels are evaluated.
	- Selection blocks until one communication can proceed, which then does.
	- If multiple can proceed, select chooses pseudo-randomly.
	- A default clause, if present, executes immediately if no channel is ready. 
*/
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	
	return c
}

func main() {
	c := fanIn(boring("Multiplexing Joe", 2), boring("Multiplexing Ann", 1))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Multiplexing ... leaving!")
}