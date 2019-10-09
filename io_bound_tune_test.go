// Example of IO-Bound tuning
// When sending multiple get request to server
// Could speed up by using go routines.
package go_playground

import (
	"net/http"
	"sync"
	"testing"
)

// Simple get request
func request() *http.Request {
	url := "https://baconator-bacon-ipsum.p.rapidapi.com/?type=all-meat"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-host", "baconator-bacon-ipsum.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "722490e69dmshc56b43fc01c02c8p14c364jsn2e38d21dc2a1")
	return req
}

// Sequential fetch
func fetch() {
	for i := 0; i < 5; i++ {
		res, _ := http.DefaultClient.Do(request())
		defer res.Body.Close()
	}
}

// Concurrent fetch
func fetchConcurrent() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			res, _ := http.DefaultClient.Do(request())
			defer res.Body.Close()
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkFetchSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fetch()
	}
}

func BenchmarkFetchConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fetchConcurrent()
	}
}
