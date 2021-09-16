# benchmarks
Measure Go performance

## Run
```bash
❯ go version
go version go1.17 darwin/amd64
❯ go test -bench=. ./...
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/crypto
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkHashString/SHA256-12           	 5485546	       214.8 ns/op	      32 B/op	       1 allocs/op
BenchmarkHashString/SHA256_directly-12         	 6502267	       180.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/BLAKE2b-12                 	 8641450	       137.4 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/crypto	4.179s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/slices
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSliceConversion/bad-12         	  822848	      1446 ns/op	    4080 B/op	       8 allocs/op
BenchmarkSliceConversion/better-12      	 1829031	       646.6 ns/op	    1792 B/op	       1 allocs/op
BenchmarkSliceConversion/best-12        	 2004114	       592.5 ns/op	    1792 B/op	       1 allocs/op
PASS
ok  	github.com/estensen/benchmarks/slices	5.897s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/strings
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkConcatString/concat_string-12         	  111310	      9900 ns/op	   53480 B/op	      99 allocs/op
BenchmarkConcatString/concat_buffer-12         	 1000000	      1094 ns/op	    3440 B/op	       6 allocs/op
BenchmarkConcatString/concat_builder-12        	 1507430	       809.6 ns/op	    2032 B/op	       7 allocs/op
BenchmarkMatchString/match_string-12           	  416467	      2736 ns/op	    3714 B/op	      27 allocs/op
BenchmarkMatchStringCompiled/match_string-12   	 2204418	       543.8 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/strings	7.424s
```
