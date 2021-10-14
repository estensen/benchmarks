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
BenchmarkHashString/SHA256-12                	 5068338	       226.1 ns/op	      32 B/op	       1 allocs/op
BenchmarkHashString/SHA256_directly-12         	 6359607	       187.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/BLAKE2b-12                 	 8534854	       140.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/MD5-12                     	11022841	       108.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/xxHash-12                  	73441771	        16.31 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/crypto	6.879s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/slices
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSliceConversion/bad-12         	  663001	      1539 ns/op	    4080 B/op	       8 allocs/op
BenchmarkSliceConversion/better-12      	 1835328	       658.6 ns/op	    1792 B/op	       1 allocs/op
BenchmarkSliceConversion/best-12        	 1916281	       727.1 ns/op	    1792 B/op	       1 allocs/op
PASS
ok  	github.com/estensen/benchmarks/slices	5.026s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/strings
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkConcatString/concat_string-12         	  110510	     10278 ns/op	   53480 B/op	      99 allocs/op
BenchmarkConcatString/concat_buffer-12         	 1000000	      1134 ns/op	    3440 B/op	       6 allocs/op
BenchmarkConcatString/concat_builder-12        	 1368289	       883.5 ns/op	    2032 B/op	       7 allocs/op
BenchmarkCovertIntToString/convert_fmt-12      	11082727	       107.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkCovertIntToString/convert_strconv-12  	25817779	        48.29 ns/op	      24 B/op	       1 allocs/op
BenchmarkMatchString/match_string-12           	  402396	      2863 ns/op	    3718 B/op	      27 allocs/op
BenchmarkMatchStringCompiled/match_string-12   	 2075523	       593.3 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/strings	10.189s
```
