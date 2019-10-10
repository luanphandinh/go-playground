# go-tuning-examples

This projects is for some fun examples.

### IO Bound Tune (`io_bound_tune_test.go`)
* The result is very much depend on the server response, but ultimately the result from running concurrent is remarkable faster.
* When making a request through network, that have significant latency. Therefor running IO Bound processes in concurrent will have positive impact on performance
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

### CPU Bound Tune (`cpu_bound_tune_test.go`)
The result is very much depend on local machine, but ultimately the result from running sequential CPU Bound task with single core is remarkable faster.
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
### Data Over Object Oriented Tune (`data_over_object_oriented_tune_test.go`)
This example comparing:
For a project that running large amount of working over time

* Instead of saving flag `active` inside a object (Object Oriented Programming)
-> Then traverse through the list object if active then doSomething()

* We then have an array of active flag (true or false) (Data Oriented Programming)
Loop through the array then doSomething()
It could help acquire a positive performance impact

A good notice that sometimes we don't need to put a boolean flag to indicate inside an object
It faster to loop through array for `bool`(1 byte)
instead of looping through array of heavy object (> 1 byte) just for a bool

* Guidance from: code::dive conference 2014 - Scott Meyers: Cpu Caches and Why You Care
https://www.youtube.com/watch?v=WDIkqP4JbkE
```bash
GOGC=off go test -cpu 1 -run none -bench Oriented -benchtime 3s
goos: darwin
goarch: amd64
pkg: github.com/luanphandinh/go-tuning-examples
BenchmarkDataOriented               1000           9062617 ns/op
BenchmarkObjectOriented              300          18142517 ns/op
PASS
ok      github.com/luanphandinh/go-tuning-examples      16.732s
```