package main

import (
	"fmt"
	"time"
	"math/rand"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

/*
	The go statement runs the function as usual, but doesn't make the caller wait.
	It launches a goroutine.
	The functionality is analogous to the & on the end of a shell command. 
*/
func main() {
	go boring("less boring")
}