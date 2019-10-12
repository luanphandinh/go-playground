// Example of IO-Bound tuning
// When sending multiple get request to server
// Could speed up by using go routines.
package go_playground

import (
	"net/http"
	"sync"
	"testing"
)

// example benchmark:
// GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s
// goos: darwin
// goarch: amd64
// pkg: github.com/luanphandinh/go-playground
// BenchmarkSequential            1        5396369911 ns/op
// BenchmarkConcurrent            3        1481725928 ns/op : ~72.5% faslter
// PASS
// ok      github.com/luanphandinh/go-playground   14.246s
//
// When making a request through network, that have significant latency
// Therefor running IO Bound processes in concurrent will have positive impact on performance
func request() *http.Request {
	url := "https://baconator-bacon-ipsum.p.rapidapi.com/?type=all-meat"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-host", "baconator-bacon-ipsum.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "722490e69dmshc56b43fc01c02c8p14c364jsn2e38d21dc2a1")
	return req
}

// example benchmark:
// GOGC=off go test -cpu 1 -run none -bench Fetch -benchtime 3s
// goos: darwin
// goarch: amd64
// pkg: github.com/luanphandinh/go-tuning-examples
// BenchmarkIOFetchSequential            2000           3256988 ns/op
// BenchmarkFetchConcurrent            2000           6039788 ns/op
// PASS
// ok      github.com/luanphandinh/go-tuning-examples      19.562s
//
// As of this example
// When making a request on local machine, there is no positive impact
// In fact it have negative impact since the latency is too small
func localRequest() *http.Request {
	url := "http://localhost:3000"
	req, _ := http.NewRequest("GET", url, nil)
	return req
}

// Sequential fetch
func fetch() {
	for i := 0; i < 5; i++ {
		// replace with localRequest() to see local benchmark
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
			// replace with localRequest() to see local benchmark
			res, _ := http.DefaultClient.Do(request())
			defer res.Body.Close()
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkIOFetchSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fetch()
	}
}

func BenchmarkIOFetchConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fetchConcurrent()
	}
}
