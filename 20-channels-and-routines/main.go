package main

import (
	"fmt"
	"net/http"
	"time"
)

type logWriter struct{}

func main() {
	links := []string{
		"http://amazon.com",
		"http://canva.com",
		"http://facebook.com",
		"http://golang.com",
		"http://google.com",
		"http://stackoverflow.com",
	}

	// making a channel:
	// a channel lets you send data between threads (go routines).
	c := make(chan string)

	for _, l := range links {
		// a go routine is initalised with `go`.
		// think of a routine as a separate thread
		go checkLink(l, c)
	}

	// waiting for a response from the channel is a blocking call.
	// fmt.Println(<- c)

	// for loop:
	// for i := 0; i < len(links); i++
	// infinite loop:
	//for {
	//	// this will keep on fetching links, return the link, fetch again etc
	//	go checkLink(<-c, c)
	//}

	// this code is euivilent to the `for {` loop above ^
	// code clarity: iterating over a channel should be done like this, not just
	// using `for {`:
	for l := range c {
		// note: putting the time.Sleep() here will cause the main thread to block.
		// this is bad because it means we can only do something once every X seconds.

		// this will keep on fetching links, return the link, fetch again etc
		// think of this as an anonymous functinon:
		// we dont want to reference the same address in memory for l:
		go func(link string) {
			// never ever share the same variable with the child routine.
			// we need a separate `link` var as `l` is used by the main routine.
			checkLink(link, c)
			time.Sleep(5 * time.Second)
		}(l) // you need the `()` at the end to call the function.
	}
}

func checkLink(link string, c chan string) {
	// putting the sleep here works better than the blocking the main thread.
	// https://golang.org/pkg/time/#example_Sleep
	// time.Sleep(5 * time.Second)
	// but this still sucks because the functinos job is to check link not sleep.
	// https://golang.org/pkg/time/#example_Sleep
	time.Sleep(5 * time.Second)

	// `http.Get` is a blocking call
	_, err := http.Get(link)
	if err != nil {
		println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "link is up!")
	c <- link
}