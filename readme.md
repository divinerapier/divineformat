# DIVINEFORMAT

## Benchmark

``` 
$ go test -run=none -bench=. -benchmem -cpuprofile=cpu.out  -memprofile mem.out
goos: darwin
goarch: amd64
pkg: github.com/divinerapier/divineformat
BenchmarkJSON-4          5000000               406 ns/op             132 B/op          3 allocs/op
BenchmarkSTD-4           1000000              1074 ns/op             320 B/op          1 allocs/op
```