package strings

import "regexp"

func MatchString(pattern, s string) bool {
	result, err := regexp.MatchString(pattern, s)
	if err != nil {
		panic(err)
	}
	return result
}

func CompilePattern(pattern string) *regexp.Regexp {
	r, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	return r
}

func MatchStringCompiled(pattern *regexp.Regexp, s string) bool {
	result := pattern.MatchString(s)
	return result
}
