# go-tuning-examples

This projects is for some fun examples.
## Table of Contents
* [Concurrently](#concurrently)
    * [IO Bound tune](#io-bound-tune)   
    * [CPU Bound tune](#cpu-bound-tune)
* [2D Array, Row or Col Travel faster ????](#row-or-column-traverse-tune)
* [Data Over Object Oriented Tune](#data-over-object-oriented-tune)
* [String concat tune](#string-concat-tune)
* [References](#references)

<a name="io-bound-tune"></a>
### IO Bound Tune (`io_bound_tune_test.go`)
```bash
GOGC=off go test -cpu 1 -run none -bench IO -benchtime 3s
goos: darwin
goarch: amd64
pkg: github.com/luanphandinh/go-tuning-examples
BenchmarkIOFetchSequential             1        5965586590 ns/op
BenchmarkIOFetchConcurrent             3        1137740573 ns/op
PASS
ok      github.com/luanphandinh/go-tuning-examples      13.592s
```
When make a request through network that have significant latency -> Current routine will enter wait state and wait for the response
* Sequential fetch making 5 request with `5 * latency` adding to execution time
```go
func fetch() {
    for i := 0; i < 5; i++ {
        // replace with localRequest() to see local benchmark
        res, _ := http.DefaultClient.Do(request())
        defer res.Body.Close()
    }
}
```
* Concurrent fetch will help improve latency, when 1 routine enter `wait` state, others `runnable` routine will be taken into `executing` state
```go
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
```
* Therefor running IO Bound processes concurrently will have positive impact on performance

<a name="cpu-bound-tune"></a>
### CPU Bound Tune (`cpu_bound_tune_test.go`)
Unlike IO Bound if we use concurrent in CPU Bound, the result might not have a good impact
```bash
GOGC=off go test -cpu 1 -run none -bench CPUCount -benchtime 3s
goos: darwin
goarch: amd64
pkg: github.com/luanphandinh/go-tuning-examples
BenchmarkCPUCountSequential     10000000               614 ns/op
BenchmarkCPUCountConcurrent       500000             11235 ns/op
PASS
ok      github.com/luanphandinh/go-tuning-examples      13.088s
```

* Sequential count
```go
func count() int {
    count := 0
    for i := 0; i < 2000; i++ {
        count++
    }
    
    return count
}
```
* Because `switching context` take times, and it tooks longer than CPU Bound process,
so running concurrently could cause performance impact
```go
// 2 `goroutine` run concurrently 
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
```
The result is very much depend on local machine, but ultimately the result from running sequential CPU Bound task with single core is remarkable faster.

<a name="row-or-column-traverse-tune"></a>
### 2D Array, Row or Col Travel faster
```bash
GOGC=off go test -cpu 1 -run none -bench Traverse -benchtime 3s
goos: darwin
goarch: amd64
pkg: github.com/luanphandinh/go-tuning-examples
BenchmarkRowTraverse          50          85126221 ns/op
BenchmarkColTraverse           5         820629562 ns/op
PASS
ok      github.com/luanphandinh/go-tuning-examples      13.050s
```
[As explained by Scott Meyers: Cpu Caches and Why You Care.](#references)
* Even though the traverse code between column and row are pretty much identical.
* But reading from `right` to `left` is more `"cached friendly"`, the under hardware with cached mechanism can help to predict the next index of array you are refer to
* Traverse by `column` could cause miss cached.

<a name="data-over-object-oriented-tune"></a>
### Data Over Object Oriented Tune (`data_over_object_oriented_tune_test.go`)
```bash
GOGC=off go test -cpu 1 -run none -bench Oriented -benchtime 3s
goos: darwin
goarch: amd64
pkg: github.com/luanphandinh/go-tuning-examples
BenchmarkObjectOriented             1000           5983468 ns/op
BenchmarkDataArrayOriented          1000           5924401 ns/op
BenchmarkDataMapOriented            1000           6224669 ns/op
PASS
ok      github.com/luanphandinh/go-tuning-examples      20.475s
```

For a project that running large workload over time

* Instead of saving flag `active` inside a object (Object Oriented Programming)
-> Then traverse through the list object if active then doSomething()
```go
type Obj struct {
    active bool
    heavy  string
}

// Using flag inside object
var objects []*Obj

func objectsDoSomething() {
    for _, obj := range objects {
        if obj.active {
            obj.doSomething()
        }
    }
}
```

* We then have an array of active flag (true or false) (Data Oriented Programming)
Loop through the array then doSomething()
It could help acquire a positive performance impact
```go
// Using array of flags
var flags []bool

func dataDoSomething() {
    for i, val := range flags {
        if val {
            objects[i].doSomething()
        }
    }
}
```

A good notice that sometimes we don't need to put a boolean inside an object
It faster to loop through array for `bool`(1 byte)
instead of looping through array of heavy object (> 1 byte) just for a bool

* Guidance from: code::dive conference 2014 - Scott Meyers: Cpu Caches and Why You Care
https://www.youtube.com/watch?v=WDIkqP4JbkE

<a name="string-concat-tune"></a>
### Tune String Concat (`string_concat_tune_test.go`)
* Using `byte` with `append` is fastest.
* Using `+` for concat string is faster than using `fmt.Sprintf()`
* Also the more you use `fmt.Sprintf()` or `+=` the more resource wasted
```go

func doConcat(str1 string, str2 string, str3 string) string {
	a := str1
	a += str2
	a += str3

	return a
}

func doByteConcatWithKnownLength(str1 string, str2 string, str3 string) string {
	length := len(str1) + len(str2) + len(str3)
	a := make([]byte, 0, length)
	a = append(a, str1...)
	a = append(a, str2...)
	a = append(a, str3...)

	return string(a)
}

func doFormat(str1 string, str2 string, str3 string) string {
	return fmt.Sprintf("%s %s %s", str1, str2, str3)
}
``` 
```bash
GOGC=off go test -cpu 1 -run none -bench ConcatString -benchtime 3s
goos: darwin
goarch: amd64
pkg: github.com/luanphandinh/go-tuning-examples
BenchmarkConcatString                           50000000                75.8 ns/op            24 B/op          2 allocs/op
BenchmarkConcatStringFormat                     20000000               207 ns/op              64 B/op          4 allocs/op
BenchmarkConcatStringFormatMultiple             20000000               339 ns/op              80 B/op          7 allocs/op
BenchmarkConcatStringByBytes                    100000000               53.7 ns/op            24 B/op          2 allocs/op
BenchmarkConcatStringByBytesKnownLength         200000000               28.8 ns/op            16 B/op          1 allocs/op
PASS
```

<a name="references"></a>
### References:
* William Kennedy: [Scheduling In Go : Part I - OS Scheduler](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html)
* William Kennedy: [ Scheduling In Go : Part II - Go Scheduler](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)
* William Kennedy: [Scheduling In Go : Part III - Concurrency](https://www.ardanlabs.com/blog/2018/12/scheduling-in-go-part3.html)
* Scott Meyers: [Cpu Caches and Why You Care](https://www.youtube.com/watch?v=WDIkqP4JbkE)
  
