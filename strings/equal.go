package strings

import "strings"

func EqualOp(first, second string) bool {
	return strings.ToLower(first) == strings.ToLower(second)
}

func EqualFold(first, second string) bool {
	return strings.EqualFold(first, second)
}

