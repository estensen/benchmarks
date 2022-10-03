# benchmarks
Measure Go performance

## Run
```bash
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/crypto
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkHashString/SHA256-12         	 5318930	       221.2 ns/op	      32 B/op	       1 allocs/op
BenchmarkHashString/SHA256_directly-12         	 6261206	       214.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/BLAKE2b-12                 	 7737580	       148.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/MD5-12                     	10936093	       120.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/xxHash-12                  	72584184	        16.54 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/crypto	7.032s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/fibonacci
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkRecursiveFibonacci_10-12     	 4969282	       238.3 ns/op
BenchmarkRecursiveFibonacci_20-12     	   39207	     30074 ns/op
BenchmarkSequentialFibonacci_10-12    	218754934	         5.284 ns/op
BenchmarkSequentialFibonacci_20-12    	100000000	        10.12 ns/op
PASS
ok  	github.com/estensen/benchmarks/fibonacci	7.330s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/io
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkReaders/ReadAll-12         	14272842	        88.68 ns/op	     512 B/op	       1 allocs/op
BenchmarkReaders/Copy-12            	28020350	        41.73 ns/op	      48 B/op	       1 allocs/op
BenchmarkReaders/ReadFrom-12        	12287956	        98.41 ns/op	     512 B/op	       1 allocs/op
PASS
ok  	github.com/estensen/benchmarks/io	3.979s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/slices
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSliceConversion/bad-12         	  794428	      1492 ns/op	    4080 B/op	       8 allocs/op
BenchmarkSliceConversion/better-12      	 1876194	       642.3 ns/op	    1792 B/op	       1 allocs/op
BenchmarkSliceConversion/best-12        	 2009995	       585.5 ns/op	    1792 B/op	       1 allocs/op
BenchmarkSlicePointers-12               	  927853	      1238 ns/op	     800 B/op	     100 allocs/op
BenchmarkSliceNoPointers-12             	20135852	        58.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceHybrid-12                 	 2646964	       443.3 ns/op	    1792 B/op	       2 allocs/op
BenchmarkAppendSlice/no_pre-alloc-12    	 8458339	       142.8 ns/op	     248 B/op	       5 allocs/op
BenchmarkAppendSlice/pre-alloc-12       	38751016	        30.55 ns/op	      80 B/op	       1 allocs/op
PASS
ok  	github.com/estensen/benchmarks/slices	12.459s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/strings
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkConcatString/concat_string-12         	  106395	     10470 ns/op	   53480 B/op	      99 allocs/op
BenchmarkConcatString/concat_buffer-12         	 1000000	      1104 ns/op	    3008 B/op	       6 allocs/op
BenchmarkConcatString/concat_builder-12        	 1000000	      1039 ns/op	    3312 B/op	       8 allocs/op
BenchmarkCovertIntToString/convert_fmt-12      	10915176	       111.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkCovertIntToString/convert_strconv-12  	25996089	        45.27 ns/op	      24 B/op	       1 allocs/op
BenchmarkEqualStrings/equal_strings_op-12      	10771042	       112.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkEqualStrings/equal_strings_fold-12    	39006230	        31.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkMatchString/match_string-12           	  405636	      3158 ns/op	    3734 B/op	      27 allocs/op
BenchmarkMatchStringCompiled/match_string-12   	 2267305	       530.8 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/strings	12.754s
```
