# go-tuning-examples

This projects is for some fun examples.

#`io_bound_tune_test.go`
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

#`cpu_bound_tune_test.go`
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
