# benchmarks
Measure Go performance

## Run
```bash
✗ go version
go version go1.19 darwin/amd64
✗ go test -bench=. ./...
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/crypto
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkHashString/SHA256-12         	 5112163	       228.5 ns/op	      32 B/op	       1 allocs/op
BenchmarkHashString/SHA256_directly-12         	 6264190	       194.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/BLAKE2b-12                 	 8467216	       146.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/MD5-12                     	10971738	       110.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/xxHash-12                  	64882150	        16.34 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/crypto	7.411s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/fibonacci
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkRecursiveFibonacci_10-12     	 5033366	       239.6 ns/op
BenchmarkRecursiveFibonacci_20-12     	   41743	     30911 ns/op
BenchmarkSequentialFibonacci_10-12    	300784536	         4.120 ns/op
BenchmarkSequentialFibonacci_20-12    	210125055	         5.756 ns/op
PASS
ok  	github.com/estensen/benchmarks/fibonacci	7.824s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/slices
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSliceConversion/bad-12         	  760296	      1480 ns/op	    4080 B/op	       8 allocs/op
BenchmarkSliceConversion/better-12      	 1922475	       623.5 ns/op	    1792 B/op	       1 allocs/op
BenchmarkSliceConversion/best-12        	 1938962	       579.2 ns/op	    1792 B/op	       1 allocs/op
BenchmarkAppendSlice/no_pre-alloc-12    	 8268492	       145.9 ns/op	     248 B/op	       5 allocs/op
BenchmarkAppendSlice/pre-alloc-12       	40356052	        30.26 ns/op	      80 B/op	       1 allocs/op
PASS
ok  	github.com/estensen/benchmarks/slices	7.435s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/strings
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkConcatString/concat_string-12         	  114259	     10619 ns/op	   53480 B/op	      99 allocs/op
BenchmarkConcatString/concat_buffer-12         	 1000000	      1085 ns/op	    3008 B/op	       6 allocs/op
BenchmarkConcatString/concat_builder-12        	 1212597	      1033 ns/op	    3312 B/op	       8 allocs/op
BenchmarkCovertIntToString/convert_fmt-12      	11454099	       103.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkCovertIntToString/convert_strconv-12  	26831100	        44.18 ns/op	      24 B/op	       1 allocs/op
BenchmarkEqualStrings/equal_strings_op-12      	11320344	       100.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkEqualStrings/equal_strings_fold-12    	39015379	        32.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkMatchString/match_string-12           	  392599	      2865 ns/op	    3735 B/op	      27 allocs/op
BenchmarkMatchStringCompiled/match_string-12   	 2151550	       545.3 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/strings	12.804s
```
