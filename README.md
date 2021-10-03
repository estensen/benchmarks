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
BenchmarkHashString/SHA256-12         	 5217225	       219.0 ns/op	      32 B/op	       1 allocs/op
BenchmarkHashString/SHA256_directly-12         	 6433054	       183.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/BLAKE2b-12                 	 8553332	       137.3 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/crypto	4.303s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/slices
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSliceConversion/bad-12         	  810795	      1450 ns/op	    4080 B/op	       8 allocs/op
BenchmarkSliceConversion/better-12      	 1812314	       654.0 ns/op	    1792 B/op	       1 allocs/op
BenchmarkSliceConversion/best-12        	 1986199	       597.3 ns/op	    1792 B/op	       1 allocs/op
PASS
ok  	github.com/estensen/benchmarks/slices	5.820s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/strings
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkConcatString/concat_string-12         	  111790	     10052 ns/op	   53480 B/op	      99 allocs/op
BenchmarkConcatString/concat_buffer-12         	 1000000	      1115 ns/op	    3440 B/op	       6 allocs/op
BenchmarkConcatString/concat_builder-12        	 1486192	       801.3 ns/op	    2032 B/op	       7 allocs/op
BenchmarkCovertIntToString/convert_fmt-12      	11820615	        99.01 ns/op	      32 B/op	       2 allocs/op
BenchmarkCovertIntToString/convert_strconv-12  	25336412	        44.36 ns/op	      24 B/op	       1 allocs/op
BenchmarkMatchString/match_string-12           	  399728	      2739 ns/op	    3713 B/op	      27 allocs/op
BenchmarkMatchStringCompiled/match_string-12   	 2117730	       560.3 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/strings	9.803s
```
