package slices

func AppendInts(numbers []int) []int {
	var result []int
	for i := 0; i < len(numbers); i++ {
		result = append(result, numbers[i])
	}
	return result
}

func AppendIntsPreAlloc(numbers []int) []int {
	result := make([]int, 0, len(numbers))
	for i := 0; i < len(numbers); i++ {
		result = append(result, numbers[i])
	}
	return result
}
