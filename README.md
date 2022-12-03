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
BenchmarkHashString/SHA256-12         	 5197537	       221.6 ns/op	      32 B/op	       1 allocs/op
BenchmarkHashString/SHA256_directly-12         	 6216489	       191.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/BLAKE2b-12                 	 8519658	       139.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/MD5-12                     	10803200	       109.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashString/xxHash-12                  	72143413	        16.11 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/crypto	6.874s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/encoding
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkByteToHex/encode_bytes_to_hex_string_with_fmt.Sprintf_for_1_byte-12         	12689306	        89.64 ns/op	      26 B/op	       2 allocs/op
BenchmarkByteToHex/encode_bytes_to_hex_string_with_hex.EncodeToString_for_1_byte-12  	80529291	        14.74 ns/op	       2 B/op	       1 allocs/op
BenchmarkByteToHex/encode_bytes_to_hex_string_with_fmt.Sprintf_for_5_bytes-12        	10836607	       108.2 ns/op	      40 B/op	       2 allocs/op
BenchmarkByteToHex/encode_bytes_to_hex_string_with_hex.EncodeToString_for_5_bytes-12 	42891154	        28.82 ns/op	      16 B/op	       1 allocs/op
BenchmarkByteToHex/encode_bytes_to_hex_string_with_fmt.Sprintf_for_10_bytes-12       	 9706875	       119.8 ns/op	      48 B/op	       2 allocs/op
BenchmarkByteToHex/encode_bytes_to_hex_string_with_hex.EncodeToString_for_10_bytes-12         	31689613	        33.88 ns/op	      24 B/op	       1 allocs/op
BenchmarkByteToHex/encode_bytes_to_hex_string_with_fmt.Sprintf_for_50_bytes-12                	 5202006	       210.4 ns/op	     136 B/op	       2 allocs/op
BenchmarkByteToHex/encode_bytes_to_hex_string_with_hex.EncodeToString_for_50_bytes-12         	10616301	       109.5 ns/op	     224 B/op	       2 allocs/op
BenchmarkByteToHex/encode_bytes_to_hex_string_with_fmt.Sprintf_for_100_bytes-12               	 5614594	       210.9 ns/op	     136 B/op	       2 allocs/op
BenchmarkByteToHex/encode_bytes_to_hex_string_with_hex.EncodeToString_for_100_bytes-12        	10426776	       108.6 ns/op	     224 B/op	       2 allocs/op
PASS
ok  	github.com/estensen/benchmarks/encoding	12.838s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/fibonacci
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkRecursiveFibonacci_10-12     	 5078794	       235.1 ns/op
BenchmarkRecursiveFibonacci_20-12     	   41163	     29002 ns/op
BenchmarkSequentialFibonacci_10-12    	227056536	         5.274 ns/op
BenchmarkSequentialFibonacci_20-12    	100000000	        10.45 ns/op
PASS
ok  	github.com/estensen/benchmarks/fibonacci	7.045s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/io
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkReaders/ReadAll-12         	13980483	        80.81 ns/op	     512 B/op	       1 allocs/op
BenchmarkReaders/Copy-12            	27849270	        40.83 ns/op	      48 B/op	       1 allocs/op
BenchmarkReaders/ReadFrom-12        	11368351	        97.12 ns/op	     512 B/op	       1 allocs/op
PASS
ok  	github.com/estensen/benchmarks/io	3.815s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/slices
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSliceConversion/bad-12         	  814022	      1464 ns/op	    4080 B/op	       8 allocs/op
BenchmarkSliceConversion/better-12      	 1859724	       634.5 ns/op	    1792 B/op	       1 allocs/op
BenchmarkSliceConversion/best-12        	 2039358	       575.6 ns/op	    1792 B/op	       1 allocs/op
BenchmarkSlicePointers-12               	  920289	      1236 ns/op	     800 B/op	     100 allocs/op
BenchmarkSliceNoPointers-12             	19781014	        60.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceHybrid-12                 	 2700028	       436.1 ns/op	    1792 B/op	       2 allocs/op
BenchmarkAppendSlice/no_pre-alloc-12    	 7955655	       143.2 ns/op	     248 B/op	       5 allocs/op
BenchmarkAppendSlice/pre-alloc-12       	33313550	        30.94 ns/op	      80 B/op	       1 allocs/op
PASS
ok  	github.com/estensen/benchmarks/slices	12.330s
goos: darwin
goarch: amd64
pkg: github.com/estensen/benchmarks/strings
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkConcatString/concat_string_for_len_0-12         	1000000000	         0.3796 ns/op	       0 B/op	       0 allocs/op
BenchmarkConcatString/concat_buffer_for_len_0-12         	324184522	         3.672 ns/op	       0 B/op	       0 allocs/op
BenchmarkConcatString/concat_builder_for_len_0-12        	1000000000	         0.5025 ns/op	       0 B/op	       0 allocs/op
BenchmarkConcatString/concat_grown_builder_for_len_0-12   	478043023	         2.512 ns/op	       0 B/op	       0 allocs/op
BenchmarkConcatString/concat_string_for_len_1-12         	153706264	         7.776 ns/op	       0 B/op	       0 allocs/op
BenchmarkConcatString/concat_buffer_for_len_1-12         	21049174	        52.29 ns/op	      80 B/op	       2 allocs/op
BenchmarkConcatString/concat_builder_for_len_1-12        	42350704	        25.97 ns/op	      16 B/op	       1 allocs/op
BenchmarkConcatString/concat_grown_builder_for_len_1-12   	29984712	        37.39 ns/op	      17 B/op	       2 allocs/op
BenchmarkConcatString/concat_string_for_len_5-12         	 7606724	       154.8 ns/op	     168 B/op	       4 allocs/op
BenchmarkConcatString/concat_buffer_for_len_5-12         	14776785	        77.75 ns/op	     128 B/op	       2 allocs/op
BenchmarkConcatString/concat_builder_for_len_5-12        	12368157	        95.51 ns/op	     112 B/op	       3 allocs/op
BenchmarkConcatString/concat_grown_builder_for_len_5-12   	10364071	       112.5 ns/op	     117 B/op	       4 allocs/op
BenchmarkConcatString/concat_string_for_len_10-12        	 3149253	       376.2 ns/op	     600 B/op	       9 allocs/op
BenchmarkConcatString/concat_buffer_for_len_10-12        	 7333538	       160.5 ns/op	     304 B/op	       3 allocs/op
BenchmarkConcatString/concat_builder_for_len_10-12       	 7739191	       151.5 ns/op	     240 B/op	       4 allocs/op
BenchmarkConcatString/concat_grown_builder_forlen_10-12  	 6322935	       187.2 ns/op	     376 B/op	       5 allocs/op
BenchmarkConcatString/concat_string_for_len_50-12        	  373458	      3308 ns/op	   13288 B/op	      49 allocs/op
BenchmarkConcatString/concat_buffer_for_len_50-12        	 2008327	       597.5 ns/op	    1472 B/op	       5 allocs/op
BenchmarkConcatString/concat_builder_for_len_50-12       	 2470419	       470.7 ns/op	    1008 B/op	       6 allocs/op
BenchmarkConcatString/concat_grown_builder_for_len_50-12  	 2357031	       494.3 ns/op	    1616 B/op	       5 allocs/op
BenchmarkConcatString/concat_string_for_len_1000-12      	    1450	    699030 ns/op	 5320825 B/op	     999 allocs/op
BenchmarkConcatString/concat_buffer_for_len_1000-12      	  106977	     10629 ns/op	   42944 B/op	      10 allocs/op
BenchmarkConcatString/concat_builder_for_len_1000-12     	  114169	      9555 ns/op	   46576 B/op	      15 allocs/op
BenchmarkConcatString/concat_grown_builder_for_len_1000-12         	  120100	      9075 ns/op	   46080 B/op	       9 allocs/op
BenchmarkCovertIntToString/convert_fmt-12                         	11330158	       101.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkCovertIntToString/convert_strconv-12                     	24478507	        45.36 ns/op	      24 B/op	       1 allocs/op
BenchmarkCount/split_strings_and_count-12                         	11802714	        97.25 ns/op	      80 B/op	       1 allocs/op
BenchmarkCount/count_newlines-12                                  	282724567	         4.275 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqualStrings/equal_strings_op-12                         	10794523	       109.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkEqualStrings/equal_strings_fold-12                       	38553870	        30.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkFileSplit/split_string_to_parse-12                       	35614780	        30.65 ns/op	      16 B/op	       1 allocs/op
BenchmarkFileSplit/cut_string_to_parse-12                         	36255652	        30.72 ns/op	      16 B/op	       1 allocs/op
BenchmarkConcatFilename/concat_with_print-12                      	 8532556	       137.4 ns/op	      48 B/op	       3 allocs/op
BenchmarkConcatFilename/concat-12                                 	56896839	        20.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkMatchString/match_string-12                              	  327580	      3096 ns/op	    3744 B/op	      27 allocs/op
BenchmarkMatchStringCompiled/match_string-12                      	 2313631	       519.9 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/estensen/benchmarks/strings	47.934s
```
