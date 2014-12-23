package main

import (
	"fmt"
	"time"
	"math/rand"
)

/* Make the intervals between messages unpredictable (still under a second). */
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

/* The boring function runs on forever, like a boring party guest. */
func main() {
	boring("less boring")
}