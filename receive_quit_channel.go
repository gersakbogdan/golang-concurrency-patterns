package main

import (
	"fmt"
	"time"
	"math/rand"
)

func cleanup() {
	fmt.Println("Joe says: Ok, ok, just one sec ... cleaning up now!")
	time.Sleep(2 * time.Second)
}

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s, %d", msg, i):
				// nothing to do here
			case s := <-quit:
				fmt.Printf("Joe receive: %v\n", s)
				cleanup()
				quit <- "See you!"
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func main() {
	quit := make(chan string)
	joe := boring("Joe", quit)

	rand.Seed(time.Now().UnixNano())
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(i, ":", <-joe)
	}
	quit <- "Bye!"
	fmt.Printf("Joe says: %q\n", <-quit)
}