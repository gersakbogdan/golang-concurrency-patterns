package main

import (
	"fmt"
	"time"
	"math/rand"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	i := 0

	go func() {
		for {
			c <- fmt.Sprintf("%s %d", msg, i)
			i++
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

/*
	The output will always be in sync: Joe, Ann, Joe, Ann 
	because the channels are still blocking on send and receive
	so Ann can output if Joe is not ready
*/
func main() {
	joe := boring("Joe service")
	ann := boring("Ann service")

	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
		fmt.Println()
	}

	fmt.Println("Channels as a handle on a service ... leaving!")
}