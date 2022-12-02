# benchmarks
Measure Go performance

## Run
```bash
➜  benchmarks git:(master) ✗ go version
go version go1.19.3 darwin/amd64
➜  benchmarks git:(master) ✗ go test -bench=. ./...
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/crypto
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkHashString/SHA256-12         	 5109166	       216.1 ns/op	      32 B/op	       1 allocs/op
BenchmarkHashString/SHA256_directly-12         	 6484764	       201.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/BLAKE2b-12                 	 8925474	       134.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/MD5-12                     	10723772	       105.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/xxHash-12                  	77144521	        15.37 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/crypto	6.877s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/fibonacci
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkRecursiveFibonacci_10-12     	 5291661	       230.6 ns/op
BenchmarkRecursiveFibonacci_20-12     	   40766	     28318 ns/op
BenchmarkSequentialFibonacci_10-12    	237843906	         5.049 ns/op
BenchmarkSequentialFibonacci_20-12    	123119113	         9.747 ns/op
PASS
ok  	github.com/estensen/benchmarks/fibonacci	8.103s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/ints
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkIncrement/val_plus_increment-12         	1000000000	         0.2438 ns/op	       0 B/op	       0 allocs/op
BenchmarkIncrement/val_plus_equal-12             	1000000000	         0.2420 ns/op	       0 B/op	       0 allocs/op
BenchmarkIncrement/val_plus_plus-12              	1000000000	         0.2415 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/ints	0.998s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/io
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkReaders/ReadAll-12         	13435972	        79.80 ns/op	     512 B/op	       1 allocs/op
BenchmarkReaders/Copy-12            	28339220	        40.29 ns/op	      48 B/op	       1 allocs/op
BenchmarkReaders/ReadFrom-12        	12377146	        96.80 ns/op	     512 B/op	       1 allocs/op
PASS
ok  	github.com/estensen/benchmarks/io	3.828s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/slices
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSliceConversion/bad-12         	  808124	      1483 ns/op	    4080 B/op	       8 allocs/op
BenchmarkSliceConversion/better-12      	 1867113	       639.3 ns/op	    1792 B/op	       1 allocs/op
BenchmarkSliceConversion/best-12        	 2013908	       583.5 ns/op	    1792 B/op	       1 allocs/op
BenchmarkSlicePointers-12               	  930475	      1201 ns/op	     800 B/op	     100 allocs/op
BenchmarkSliceNoPointers-12             	20478410	        57.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceHybrid-12                 	 2658843	       457.9 ns/op	    1792 B/op	       2 allocs/op
BenchmarkAppendSlice/no_pre-alloc-12    	 7958454	       141.6 ns/op	     248 B/op	       5 allocs/op
BenchmarkAppendSlice/pre-alloc-12       	38042602	        31.16 ns/op	      80 B/op	       1 allocs/op
PASS
ok  	github.com/estensen/benchmarks/slices	12.489s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/strings
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkConcatString/concat_string-12         	  113984	     10431 ns/op	   53480 B/op	      99 allocs/op
BenchmarkConcatString/concat_buffer-12         	 1000000	      1097 ns/op	    3008 B/op	       6 allocs/op
BenchmarkConcatString/concat_builder-12        	 1000000	      1048 ns/op	    3312 B/op	       8 allocs/op
BenchmarkConcatString/concat_grown_builder-12  	 1405492	       859.5 ns/op	    2656 B/op	       5 allocs/op
BenchmarkCovertIntToString/convert_fmt-12      	11961027	       101.5 ns/op	      32 B/op	       2 allocs/op
BenchmarkCovertIntToString/convert_strconv-12  	26691764	        45.26 ns/op	      24 B/op	       1 allocs/op
BenchmarkCount/split_strings_and_count-12      	12079470	        93.81 ns/op	      80 B/op	       1 allocs/op
BenchmarkCount/count_newlines-12               	290466226	         4.186 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqualStrings/equal_strings_op-12      	11169494	       105.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkEqualStrings/equal_strings_fold-12    	37732734	        29.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkFileSplit/split_string_to_parse-12    	35873632	        30.84 ns/op	      16 B/op	       1 allocs/op
BenchmarkFileSplit/cut_string_to_parse-12      	38161476	        30.89 ns/op	      16 B/op	       1 allocs/op
BenchmarkConcatFilename/concat_with_print-12   	 8727040	       135.5 ns/op	      48 B/op	       3 allocs/op
BenchmarkConcatFilename/concat-12              	61184294	        20.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkMatchString/match_string-12           	  370126	      3111 ns/op	    3741 B/op	      27 allocs/op
BenchmarkMatchStringCompiled/match_string-12   	 2361954	       489.7 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/strings	21.351s
```
