# benchmarks
Measure Go performance

## Run
```bash
✗ go version
go version go1.18 darwin/amd64
✗ go test -bench=. ./...
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/crypto
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkHashString/SHA256-12                    5256291	       239.4 ns/op	      32 B/op	       1 allocs/op
BenchmarkHashString/SHA256_directly-12         	 6502294	       184.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/BLAKE2b-12                 	 8394993	       136.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/MD5-12                     	10999168	       105.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/xxHash-12                  	74355836	        15.80 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/crypto	6.933s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/fibonacci
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkRecursiveFibonacci_10-12         5325547	         221.6 ns/op
BenchmarkRecursiveFibonacci_20-12     	    43400	         27499 ns/op
BenchmarkSequentialFibonacci_10-12    	305352999	         3.869 ns/op
BenchmarkSequentialFibonacci_20-12    	213693325	         5.645 ns/op
PASS
ok  	github.com/estensen/benchmarks/fibonacci	7.499s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/slices
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSliceConversion/bad-12         	  825196	        1438 ns/op	    4080 B/op	       8 allocs/op
BenchmarkSliceConversion/better-12      	 1898743	       621.7 ns/op	    1792 B/op	       1 allocs/op
BenchmarkSliceConversion/best-12        	 2102976	       568.8 ns/op	    1792 B/op	       1 allocs/op
PASS
ok  	github.com/estensen/benchmarks/slices	5.876s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/strings
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkConcatString/concat_string-12         	  104883	      9672 ns/op	   53480 B/op	      99 allocs/op
BenchmarkConcatString/concat_buffer-12         	  973308	      1103 ns/op	    3440 B/op	       6 allocs/op
BenchmarkConcatString/concat_builder-12        	 1229082	      1014 ns/op	    3312 B/op	       8 allocs/op
BenchmarkCovertIntToString/convert_fmt-12      	11220274	       102.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkCovertIntToString/convert_strconv-12  	25074403	        44.37 ns/op	      24 B/op	       1 allocs/op
BenchmarkEqualStrings/equal_strings_op-12      	10785816	       109.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkEqualStrings/equal_strings_fold-12    	38511313	        30.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkMatchString/match_string-12           	  371172	      2822 ns/op	    3728 B/op	      27 allocs/op
BenchmarkMatchStringCompiled/match_string-12   	 2152146	       555.1 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/strings	12.404s
```
