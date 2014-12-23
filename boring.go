package main

import (
	"fmt"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}

/* The boring function runs on forever, like a boring party guest. */
func main() {
	boring("boring")
}