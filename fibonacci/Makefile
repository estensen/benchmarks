benchmark:
	go test -bench=. | tee graphic/out.dat ; \
	awk '/Benchmark/{count ++; gsub(/BenchmarkTest/,""); printf("%d,%s,%s,%s\n",count,$$1,$$2,$$3)}' graphic/out.dat > graphic/final.dat ; \
	gnuplot -e "file_path='graphic/final.dat'" -e "graphic_file_name='graphic/operations.png'" -e "y_label='number of operations'" -e "y_range_min='000000000''" -e "y_range_max='400000000'" -e "column_1=1" -e "column_2=3" performance.gnuplot ; \
	gnuplot -e "file_path='graphic/final.dat'" -e "graphic_file_name='graphic/time_operations.png'" -e "y_label='each operation in nanoseconds'" -e "y_range_min='000''" -e "y_range_max='45000'" -e "column_1=1" -e "column_2=4" performance.gnuplot ; \
	rm -f graphic/out.dat graphic/final.dat ; \
	echo "'graphic/operations.png' and 'graphic/time_operations.png' graphics were generated."
