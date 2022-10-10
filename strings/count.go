package strings

import "strings"

func StringSplit(s, substr string) int {
	return len(strings.Split(s, substr))
}

func StringCount(s, substr string) int {
	return strings.Count(s, substr) + 1
}
