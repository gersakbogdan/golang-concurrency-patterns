package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Message struct {
	str string
	wait chan bool
}

func boring(msg string, delay int) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)

	i := 0

	go func() {
		for {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			i++
			time.Sleep(time.Duration(rand.Intn(1e3 * delay)) * time.Millisecond)
			<-waitForIt
		}
	}()

	return c
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)

	go func() { for {c<- <-input1} }()
	go func() { for {c<- <-input2} }()
	
	return c
}

func main() {
	c := fanIn(boring("Multiplexing Joe", 2), boring("Multiplexing Ann", 1))

	for i := 0; i < 5; i++ {
		msg1 := <-c; fmt.Println(msg1.str)
		msg2 := <-c; fmt.Println(msg2.str)

		msg1.wait <- true
		msg2.wait <- true
	}

	fmt.Println("Multiplexing ... leaving!")
}