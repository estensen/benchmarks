package slices

import "strconv"

func numbersToStringsBad(numbers []int) []string {
	// Not setting size of slice is bad when you know how big it needs to be
	vals := []string{}
	for _, n := range numbers {
		vals = append(vals, strconv.Itoa(n))
	}
	return vals
}

func numbersToStringsBetter(numbers []int) []string {
	vals := make([]string, 0, len(numbers))
	for _, n := range numbers {
		vals = append(vals, strconv.Itoa(n))
	}
	return vals
}

func numbersToStringsBest(numbers []int) []string {
	vals := make([]string, len(numbers))
	for i, n := range numbers {
		vals[i] = strconv.Itoa(n)
	}
	return vals
}
