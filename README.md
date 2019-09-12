# benchmarks
Measure Go performance

## Run
```bash
$ go test -bench=. ./... 
goos: darwin
goarch: amd64
BenchmarkSliceConversion/bad-12                   706338              1525 ns/op            4080 B/op          8 allocs/op
BenchmarkSliceConversion/better-12               1686865               706 ns/op            1792 B/op          1 allocs/op
BenchmarkSliceConversion/best-12                 1848459               653 ns/op            1792 B/op          1 allocs/op
PASS
ok      _/Users/havard.anda.estensen/Developer/benchmarks/slices        4.945s

```
