# DIVINEFORMAT

## Benchmark

``` 
$ go test -run=none -bench=. -benchmem -cpuprofile=cpu.out  -memprofile mem.out
goos: darwin
goarch: amd64
pkg: github.com/divinerapier/divineformat
BenchmarkJSON_Chain-4            5000000               362 ns/op         814.81 MB/s         132 B/op          3 allocs/op
BenchmarkJSON_ManyTimes-4        5000000               413 ns/op         713.40 MB/s         132 B/op          3 allocs/op
BenchmarkSTD_Once-4              1000000              1066 ns/op         281.42 MB/s         320 B/op          1 allocs/op
BenchmarkSTD_ManyTimes-4          200000              8231 ns/op          36.45 MB/s        9508 B/op         97 allocs/op
```