# benchmarks
Measure Go performance

## Run
```bash
➜  benchmarks git:(master) ✗ go version
go version go1.19.2 darwin/amd64
➜  benchmarks git:(master) ✗ go test -bench=. ./...
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/crypto
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkHashString/SHA256-12         	 5181872	       221.4 ns/op	      32 B/op	       1 allocs/op
BenchmarkHashString/SHA256_directly-12         	 6321712	       188.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/BLAKE2b-12                 	 8517476	       138.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/MD5-12                     	11017078	       108.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/xxHash-12                  	75200337	        15.85 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/crypto	6.947s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/fibonacci
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkRecursiveFibonacci_10-12     	 5222611	       234.6 ns/op
BenchmarkRecursiveFibonacci_20-12     	   40123	     28632 ns/op
BenchmarkSequentialFibonacci_10-12    	233641792	         5.159 ns/op
BenchmarkSequentialFibonacci_20-12    	100000000	        10.09 ns/op
PASS
ok  	github.com/estensen/benchmarks/fibonacci	7.087s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/io
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkReaders/ReadAll-12         	13276198	        79.89 ns/op	     512 B/op	       1 allocs/op
BenchmarkReaders/Copy-12            	26890315	        40.41 ns/op	      48 B/op	       1 allocs/op
BenchmarkReaders/ReadFrom-12        	12344115	        95.75 ns/op	     512 B/op	       1 allocs/op
PASS
ok  	github.com/estensen/benchmarks/io	3.888s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/slices
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSliceConversion/bad-12         	  804675	      1479 ns/op	    4080 B/op	       8 allocs/op
BenchmarkSliceConversion/better-12      	 1890062	       634.0 ns/op	    1792 B/op	       1 allocs/op
BenchmarkSliceConversion/best-12        	 2059556	       586.6 ns/op	    1792 B/op	       1 allocs/op
BenchmarkSlicePointers-12               	  921242	      1232 ns/op	     800 B/op	     100 allocs/op
BenchmarkSliceNoPointers-12             	20532378	        58.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceHybrid-12                 	 2664216	       446.8 ns/op	    1792 B/op	       2 allocs/op
BenchmarkAppendSlice/no_pre-alloc-12    	 8218360	       142.9 ns/op	     248 B/op	       5 allocs/op
BenchmarkAppendSlice/pre-alloc-12       	37594345	        31.32 ns/op	      80 B/op	       1 allocs/op
PASS
ok  	github.com/estensen/benchmarks/slices	12.519s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/strings
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkConcatString/concat_string-12         	  105766	     10367 ns/op	   53480 B/op	      99 allocs/op
BenchmarkConcatString/concat_buffer-12         	 1000000	      1094 ns/op	    3008 B/op	       6 allocs/op
BenchmarkConcatString/concat_builder-12        	 1000000	      1050 ns/op	    3312 B/op	       8 allocs/op
BenchmarkCovertIntToString/convert_fmt-12      	11278172	       113.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkCovertIntToString/convert_strconv-12  	26065640	        46.33 ns/op	      24 B/op	       1 allocs/op
BenchmarkCount/split_strings_and_count-12      	12433666	        96.11 ns/op	      80 B/op	       1 allocs/op
BenchmarkCount/count_newlines-12               	246752883	         4.572 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqualStrings/equal_strings_op-12      	10184306	       125.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkEqualStrings/equal_strings_fold-12    	38079099	        34.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkMatchString/match_string-12           	  348342	      3310 ns/op	    3742 B/op	      27 allocs/op
BenchmarkMatchStringCompiled/match_string-12   	 2120810	       584.3 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/strings	14.921s
```
