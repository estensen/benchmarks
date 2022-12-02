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
BenchmarkHashString/SHA256-12         	         4991473	       216.2 ns/op	      32 B/op	       1 allocs/op
BenchmarkHashString/SHA256_directly-12         	 6328057	       182.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/BLAKE2b-12                 	 8789499	       132.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/MD5-12                     	11097973	       105.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/xxHash-12                  	75919070	        15.38 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/crypto	6.904s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/encoding
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkByteToHex/encode_bytes_to_hex_string_with_fmt.Sprintf-12                   	12420480	        86.67 ns/op	      26 B/op	       2 allocs/op
BenchmarkByteToHex/encode_bytes_to_hex_string_with_hex.EncodeToString-12            	81995619	        15.71 ns/op	       2 B/op	       1 allocs/op
BenchmarkByteToHex/encode_bytes_to_hex_string_with_fmt.Sprintf#01-12                	11079276	       106.2 ns/op	      40 B/op	       2 allocs/op
BenchmarkByteToHex/encode_bytes_to_hex_string_with_hex.EncodeToString#01-12         	43097331	        26.82 ns/op	      16 B/op	       1 allocs/op
BenchmarkByteToHex/encode_bytes_to_hex_string_with_fmt.Sprintf#02-12                	 9979802	       116.5 ns/op	      48 B/op	       2 allocs/op
BenchmarkByteToHex/encode_bytes_to_hex_string_with_hex.EncodeToString#02-12         	34111448	        34.15 ns/op	      24 B/op	       1 allocs/op
BenchmarkByteToHex/encode_bytes_to_hex_string_with_fmt.Sprintf#03-12                	 5525880	       206.1 ns/op	     136 B/op	       2 allocs/op
BenchmarkByteToHex/encode_bytes_to_hex_string_with_hex.EncodeToString#03-12         	10893450	       108.5 ns/op	     224 B/op	       2 allocs/op
PASS
ok  	github.com/estensen/benchmarks/encoding	10.320s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/fibonacci
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkRecursiveFibonacci_10-12     	  5246632	         225.8 ns/op
BenchmarkRecursiveFibonacci_20-12     	    42224	         27929 ns/op
BenchmarkSequentialFibonacci_10-12    	233767657	         5.008 ns/op
BenchmarkSequentialFibonacci_20-12    	122124894	         9.848 ns/op
PASS
ok  	github.com/estensen/benchmarks/fibonacci	8.091s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/io
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkReaders/ReadAll-12         	14559177	        80.85 ns/op	     512 B/op	       1 allocs/op
BenchmarkReaders/Copy-12            	26736693	        40.67 ns/op	      48 B/op	       1 allocs/op
BenchmarkReaders/ReadFrom-12        	12168366	        95.43 ns/op	     512 B/op	       1 allocs/op
PASS
ok  	github.com/estensen/benchmarks/io	3.901s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/slices
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSliceConversion/bad-12         	  805842	      1473 ns/op	    4080 B/op	       8 allocs/op
BenchmarkSliceConversion/better-12      	 1798948	       637.0 ns/op	    1792 B/op	       1 allocs/op
BenchmarkSliceConversion/best-12        	 1994672	       599.3 ns/op	    1792 B/op	       1 allocs/op
BenchmarkSlicePointers-12               	  931839	      1222 ns/op	     800 B/op	     100 allocs/op
BenchmarkSliceNoPointers-12             	20180094	        58.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceHybrid-12                 	 2612734	       443.3 ns/op	    1792 B/op	       2 allocs/op
BenchmarkAppendSlice/no_pre-alloc-12    	 8299556	       142.4 ns/op	     248 B/op	       5 allocs/op
BenchmarkAppendSlice/pre-alloc-12       	37051797	        30.79 ns/op	      80 B/op	       1 allocs/op
PASS
ok  	github.com/estensen/benchmarks/slices	12.419s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/strings
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkConcatString/concat_string-12         	  113370	     10261 ns/op	   53480 B/op	      99 allocs/op
BenchmarkConcatString/concat_buffer-12         	 1000000	      1083 ns/op	    3008 B/op	       6 allocs/op
BenchmarkConcatString/concat_builder-12        	 1000000	      1040 ns/op	    3312 B/op	       8 allocs/op
BenchmarkConcatString/concat_grown_builder-12  	 1414446	       861.7 ns/op	    2656 B/op	       5 allocs/op
BenchmarkCovertIntToString/convert_fmt-12      	11709692	       102.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkCovertIntToString/convert_strconv-12  	25907197	        44.68 ns/op	      24 B/op	       1 allocs/op
BenchmarkCount/split_strings_and_count-12      	12484929	        95.97 ns/op	      80 B/op	       1 allocs/op
BenchmarkCount/count_newlines-12               	274617675	         4.237 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqualStrings/equal_strings_op-12      	10987564	       106.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkEqualStrings/equal_strings_fold-12    	39737151	        29.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkFileSplit/split_string_to_parse-12    	38446472	        30.36 ns/op	      16 B/op	       1 allocs/op
BenchmarkFileSplit/cut_string_to_parse-12      	37522386	        30.19 ns/op	      16 B/op	       1 allocs/op
BenchmarkConcatFilename/concat_with_print-12   	 8048638	       133.9 ns/op	      48 B/op	       3 allocs/op
BenchmarkConcatFilename/concat-12              	62865579	        19.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkMatchString/match_string-12           	  357582	      3101 ns/op	    3738 B/op	      27 allocs/op
BenchmarkMatchStringCompiled/match_string-12   	 2326338	       498.6 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/strings	21.229s
```
