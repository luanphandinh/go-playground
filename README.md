# go-tuning-examples

This projects is for some fun examples.
## Table of Contents
* [Concurrently](#concurrently)
    * [IO Bound tune](#io-bound-tune)   
    * [CPU Bound tune](#cpu-bound-tune)
* [Data Over Object Oriented Tune](#data-over-object-oriented-tune)
* [String concat tune](#string-concat-tune)

<a name="io-bound-tune"></a>
### IO Bound Tune (`io_bound_tune_test.go`)
```bash
GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s
goos: darwin
goarch: amd64
pkg: github.com/luanphandinh/go-playground
BenchmarkSequential            1        5396369911 ns/op
BenchmarkConcurrent            3        1481725928 ns/op : ~72.5% faslter
PASS
ok      github.com/luanphandinh/go-playground   14.246s
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
GOGC=off go test -cpu 1 -run none -bench Count -benchtime 3s
goos: darwin
goarch: amd64
pkg: github.com/luanphandinh/go-playground
BenchmarkCountSequential         5000000               644 ns/op
BenchmarkCountConcurrent          500000             11688 ns/op
PASS
ok      github.com/luanphandinh/go-playground   10.405s
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


<a name="data-over-object-oriented-tune"></a>
### Data Over Object Oriented Tune (`data_over_object_oriented_tune_test.go`)
```bash
GOGC=off go test -cpu 1 -run none -bench Oriented -benchtime 3s
goos: darwin
goarch: amd64
pkg: github.com/luanphandinh/go-tuning-examples
BenchmarkDataOriented               1000           8871345 ns/op
BenchmarkObjectOriented              300          17673123 ns/op
PASS
ok      github.com/luanphandinh/go-tuning-examples      16.929s
```

For a project that running large workload over time

* Instead of saving flag `active` inside a object (Object Oriented Programming)
-> Then traverse through the list object if active then doSomething()
```go
type Obj struct {
    active bool
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

A good notice that sometimes we don't need to put a boolean flag to indicate inside an object
It faster to loop through array for `bool`(1 byte)
instead of looping through array of heavy object (> 1 byte) just for a bool

* Guidance from: code::dive conference 2014 - Scott Meyers: Cpu Caches and Why You Care
https://www.youtube.com/watch?v=WDIkqP4JbkE

<a name="string-concat-tune"></a>
### Tune String Concat (`string_concat_tune_test.go`)
Using `+` for concat string is faster than using `fmt.Sprintf()`
```go
func doConcat() string {
    return "This" + "is" + "simple" + "concat" + "string"
}

func doFormat() string {
    return fmt.Sprintf("%s %s %s %s %s", "This", "is", "simple", "format", "string")
}
``` 
```bash
GOGC=off go test -cpu 1 -run none -bench String -benchtime 3s
goos: darwin
goarch: amd64
pkg: github.com/luanphandinh/go-tuning-examples
BenchmarkConcatString           5000000000               0.57 ns/op
BenchmarkConcatStringFormat     20000000               183 ns/op
PASS
ok      github.com/luanphandinh/go-tuning-examples      6.884s
```