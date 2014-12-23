package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(left, right chan int) {
	left<- <-right
}

func main() {
	const n = 10000

	leftmost := make(chan int)
	right := make(chan int)

	left := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right);
		left = right
	}

	rand.Seed(time.Now().UnixNano())
	say := rand.Intn(n)
	fmt.Println("Say: ", say)

	go func(c chan int) { c <- say }(right)
	
	fmt.Println("Receive: ", <-leftmost)
}