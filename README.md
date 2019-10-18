# benchmarks
Measure Go performance

## Run
```bash
$ go test -bench=. ./... 
goos: darwin
goarch: amd64
BenchmarkSliceConversion/bad-12         	  703904	      1585 ns/op	    4080 B/op	       8 allocs/op
BenchmarkSliceConversion/better-12      	 1452730	       804 ns/op	    1792 B/op	       1 allocs/op
BenchmarkSliceConversion/best-12        	 1631972	       730 ns/op	    1792 B/op	       1 allocs/op
PASS
ok  	_/Users/havard.anda.estensen/Developer/benchmarks/slices	5.338s
goos: darwin
goarch: amd64
BenchmarkConcatString/concat_string-12         	  109778	     10480 ns/op	   53488 B/op	      99 allocs/op
BenchmarkConcatString/concat_buffer-12         	  795718	      1502 ns/op	    3440 B/op	       6 allocs/op
BenchmarkConcatString/concat_builder-12        	 1378024	       882 ns/op	    2032 B/op	       7 allocs/op
PASS
ok  	_/Users/havard.anda.estensen/Developer/benchmarks/strings	4.749s
```
