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
 	When main returns, the program exits and takes the boring function down with it.
	We can hang around a little, and on the way show that both main and the launched
	goroutine are running. 
*/
func main() {
	go boring("less boring")
	fmt.Println("Listening...")
	time.Sleep(2 * time.Second)
	fmt.Println("Still boring ... leaving.")
}