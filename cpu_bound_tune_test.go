// Example of CPU-Bound tuning
package go_playground

import (
	"sync"
	"sync/atomic"
	"testing"
)

// Go from 1 to 2000 and increase count
func count() int {
	count := 0
	for i := 0; i < 2000; i++ {
		count++
	}

	return count
}

// Same goal to count to 2000
// But we will use 2 goroutines(each 1000 loop) to calculate concurrently
func countConcurrent() int {
	var count int32
	var wg sync.WaitGroup
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&count, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return int(count)
}

func BenchmarkCPUCountSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		count()
	}
}

// When running CPU Bound
// 		This func will spend times on context-switching between goroutines
//		Could be performance affect with 1 CPU Core
func BenchmarkCPUCountConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countConcurrent()
	}
}
