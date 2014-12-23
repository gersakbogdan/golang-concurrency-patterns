package main

import (
	"fmt"
	"time"
	"math/rand"
)

var (
	Web1 = fakeSearch("web")
	Web2 = fakeSearch("web")
	Web3 = fakeSearch("web")

	Image1 = fakeSearch("image")
	Image2 = fakeSearch("image")
	Image3 = fakeSearch("image")

	Video1 = fakeSearch("video")
	Video2 = fakeSearch("video")
	Video3 = fakeSearch("video")
)

type Result string
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string) (results []Result) {
	c := make(chan Result)

	go func() { c <- First(query, Web1, Web2, Web3) }()
	go func() { c <- First(query, Image1, Image2, Image3) }()
	go func() { c <- First(query, Video1, Video2, Video3) }()

	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("Ok! Time's out")
			fmt.Println()
			return
		}
	}

	return
}

func First(query string, replica ...Search) Result {
	c := make(chan Result)

	for i := range replica {
		go func() {
			c <- replica[i](query)
		}()
	}

	return <-c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()

	results := Google("golang")

	elapsed := time.Since(start)

	fmt.Println(results)
	fmt.Println(elapsed)
}