package main

import (
	"fmt"
	"time"
	"math/rand"
)

var (
	Web = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
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
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))

	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()

	results := Google("golang")

	elapsed := time.Since(start)

	fmt.Println(results)
	fmt.Println(elapsed)
}