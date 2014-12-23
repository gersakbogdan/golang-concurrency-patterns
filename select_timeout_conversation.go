package main

import (
	"fmt"
	"time"
	"math/rand"
)

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

/*
	Conversation timeout example.
	Timeout select case is triggered after time expires, so the conversation is stopped!
	This is possible because the timer is created outside the loop
*/
func main() {
	joe := boring("Joe")
	timeout := time.After(5 * time.Second)

	for {
		select {
		case s := <-joe:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You are too slow. Conversation stopped!")
			return
		}
	}
}