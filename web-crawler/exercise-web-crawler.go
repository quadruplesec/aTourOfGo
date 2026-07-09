package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	v  map[string]int
	mu sync.Mutex
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	var c SafeCounter = SafeCounter{v: make(map[string]int)}

	var crawler func(url string, depth int, exit chan bool)
	crawler = func(url string, depth int, exit chan bool) {
		if depth <= 0 {
			exit <- true
			return
		}

		// I tried defering here but it caused a deadlock
		// The children wanted the lock that the parent was holding...
		c.mu.Lock()
		if c.v[url] > 0 {
			exit <- true
			c.mu.Unlock()
			return
		}
		c.v[url]++
		c.mu.Unlock()

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			exit <- true
			return
		}

		childrenDone := make(chan bool)
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			go crawler(u, depth-1, childrenDone)
		}

		// Wait for all children goroutines to finish
		for i := 0; i < len(urls); i++ {
			<-childrenDone
		}

		// Signal to the parent that we finished (same thing is done above)
		exit <- true
	}

	exit := make(chan bool)
	go crawler(url, depth, exit)
	<-exit

	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
