# go-playground

This projects is for some fun examples.

#`io_bound_tune_test.go`
The result is very much depend on the server response, but ultimately the result from running concurrent is remarkable faster.
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
