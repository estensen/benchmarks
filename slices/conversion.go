package slices

import "strconv"

func NumbersToStringsBad(numbers []int) []string {
	// Not setting size of slice is bad when you know how big it needs to be
	vals := []string{}
	for _, n := range numbers {
		vals = append(vals, strconv.Itoa(n))
	}
	return vals
}

func NumbersToStringsBetter(numbers []int) []string {
	vals := make([]string, 0, len(numbers))
	for _, n := range numbers {
		vals = append(vals, strconv.Itoa(n))
	}
	return vals
}

func NumbersToStringsBest(numbers []int) []string {
	vals := make([]string, len(numbers))
	for i, n := range numbers {
		vals[i] = strconv.Itoa(n)
	}
	return vals
}
