# Fibonacci Benchmark Test

1. Object
    - BenchmarkFib - Origin 
    - BenchmarkFib1 - Linear Alg. solution
    - BenchmarkFib2 - Formula

2. Result
```
$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: fibonacci
BenchmarkFib-8               500           3210880 ns/op               0 B/op          0 allocs/op
BenchmarkFib1-8         20000000                70.8 ns/op             0 B/op          0 allocs/op
BenchmarkFib2-8         50000000                35.0 ns/op             0 B/op          0 allocs/op
PASS
ok      fibonacci       5.215s
```