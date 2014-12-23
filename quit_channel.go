package main

import (
	"fmt"
	"time"
	"math/rand"
)

func boring(msg string, quit chan bool) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s, %d", msg, i):
				// nothing to do here
			case <-quit:
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func main() {
	quit := make(chan bool)
	joe := boring("Joe", quit)

	rand.Seed(time.Now().UnixNano())
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(i, ":", <-joe)
	}
	quit <- true
}