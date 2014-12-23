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
			time.Sleep(time.Duration(rand.Intn(1e3) * i / 2) * time.Millisecond)
		}
	}()

	return c
}

/*
	Timeout for each message.
	Timeout select case is triggered only if there is a message which takes more then 1sec

	Check select_timeout_conversation for a different way of using time.After
*/
func main() {
	joe := boring("Joe")

	for {
		select {
		case s := <-joe:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("You are too slow!")
			return
		}
	}
}