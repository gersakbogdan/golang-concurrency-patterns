package main

import (
	"fmt"
	"time"
	"math/rand"
)

/* 
	We run the goroutine here so we can return the channel and wait for message in main method.
	Try to remote the "go" before func() to see what's happend
*/
func boring(msg string) <-chan string { // returns a receive only channelr
	c := make(chan string)
	i := 0

	// this runs in background so our boring method is not blocked and we can return the channel
	go func() {
		for {
			c <- fmt.Sprintf("%s %d", msg, i)
			i++
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

			/* [2]
			if i == 10 {
				close(c)
				break;
			}
			*/
		}
	}()

	return c
}

func main() {
	c := boring("Pattern generator")
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}

	/* [2]
	for v := range c {
		fmt.Println(v)
	}
	*/

	fmt.Println("Pattern generator ... leaving!")
}