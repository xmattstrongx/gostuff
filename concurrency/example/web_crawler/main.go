package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, v *UrlVisitedTracker) {
	v.addUrlToList(url)
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Printf("FOUND For url: %s... found: %q with urls: %s\n", url, body, urls)
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if v.isAlreadyFound(u) {
			continue
		}
		v.wg.Add(1)
		go Crawl(u, depth-1, fetcher, v)
	}
	v.wg.Done()
	return
}

func (u *UrlVisitedTracker) isAlreadyFound(url string) bool {
	u.mux.Lock()
	defer u.mux.Unlock()
	if val, ok := u.visitedUrls[url]; ok && val > 0 {
		return true
	}
	return false
}

func (u *UrlVisitedTracker) addUrlToList(url string) {
	u.mux.Lock()
	defer u.mux.Unlock()
	u.visitedUrls[url]++
}

func main() {
	start := time.Now()
	v := UrlVisitedTracker{visitedUrls: make(map[string]int)}
	Crawl("http://golang.org/", 4, fetcher, &v)
	v.wg.Wait()
	fmt.Printf("Searchh took: %s\n", time.Since(start))
}

type UrlVisitedTracker struct {
	visitedUrls map[string]int
	mux         sync.Mutex
	wg          sync.WaitGroup
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
