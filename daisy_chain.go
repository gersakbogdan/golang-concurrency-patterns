package main

import (
	"fmt"
)

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 100000

	leftmost := make(chan int)
	right := make(chan int)
	left := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func(c chan int) {
		c <- 0
	}(right)

	fmt.Println(<-leftmost)
}