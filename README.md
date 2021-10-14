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
BenchmarkHashString/SHA256-12       		  	 5290388	       221.7 ns/op	      32 B/op	       1 allocs/op
BenchmarkHashString/SHA256_directly-12         	 6403832	       186.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/BLAKE2b-12                 	 8538073	       141.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/xxHash-12                  	73156243	        15.99 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/crypto	5.466s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/slices
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSliceConversion/bad-12         	  696118	      1460 ns/op	    4080 B/op	       8 allocs/op
BenchmarkSliceConversion/better-12      	 1824963	       670.1 ns/op	    1792 B/op	       1 allocs/op
BenchmarkSliceConversion/best-12        	 1913059	       612.9 ns/op	    1792 B/op	       1 allocs/op
PASS
ok  	github.com/estensen/benchmarks/slices	4.821s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/strings
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkConcatString/concat_string-12         	  116941	     10192 ns/op	   53480 B/op	      99 allocs/op
BenchmarkConcatString/concat_buffer-12         	  958395	      1127 ns/op	    3440 B/op	       6 allocs/op
BenchmarkConcatString/concat_builder-12        	 1430810	       843.0 ns/op	    2032 B/op	       7 allocs/op
BenchmarkCovertIntToString/convert_fmt-12      	11264161	       105.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkCovertIntToString/convert_strconv-12  	24974967	        46.08 ns/op	      24 B/op	       1 allocs/op
BenchmarkMatchString/match_string-12           	  393256	      2847 ns/op	    3720 B/op	      27 allocs/op
BenchmarkMatchStringCompiled/match_string-12   	 2064314	       556.2 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/strings	9.933s
```
