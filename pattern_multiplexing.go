package main

import (
	"fmt"
	"time"
	"math/rand"
)

func boring(msg string, delay int) <-chan string {
	c := make(chan string)
	i := 0

	go func() {
		for {
			c <- fmt.Sprintf("%s %d", msg, i)
			i++
			time.Sleep(time.Duration(rand.Intn(1e3 * delay)) * time.Millisecond)
		}
	}()

	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() { for {c<- <-input1} }()
	go func() { for {c<- <-input2} }()
	
	return c
}

func main() {
	c := fanIn(boring("Multiplexing Joe", 2), boring("Multiplexing Ann", 1))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Multiplexing ... leaving!")
}