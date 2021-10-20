# 800x600 png will be generated
set terminal png

# Allow grid lines to be drawn on the plot
set grid

# Set graphic_file_name to be saved
set output graphic_file_name

# Main title
set title "performance comparison"

# Input is CSV file where separator is ","
set datafile separator ","

# Disable key box
set key off

set ylabel y_label

# Range for values in y axis
set yrange[y_range_min:y_range_max]

# Avoid displaying large numbers in exponential format
set format y "%.0f"

set xtics rotate

set style fill solid
set boxwidth 0.5

plot for [i=0:*] file_path every ::i::i using column_1:column_2:xtic(2) with boxes
