package fibonacci

func RecursiveFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return RecursiveFibonacci(n-1) + RecursiveFibonacci(n-2)
}

func SequentialFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	n2 := 0
	n1 := 1
	for i := 2; i < n; i++ {
		n2, n1 = n1, n1+n2
	}
	return n2 + n1
}
