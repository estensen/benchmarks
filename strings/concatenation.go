package strings

import (
	"bytes"
	"strings"
)

func ConcatString(words []string) string {
	var str string

	for _, word := range words {
		str += word
	}
	return str
}

func ConcatBuffer(words []string) string {
	var buffer bytes.Buffer

	for _, word := range words {
		buffer.WriteString(word)
	}
	return buffer.String()
}

func ConcatBuilder(words []string) string {
	var builder strings.Builder

	for _, word := range words {
		builder.WriteString(word)
	}

	return builder.String()
}

func ConcatGrownBuilder(words []string) string {
	var builder strings.Builder
	builder.Grow(len(words))

	for _, word := range words {
		builder.WriteString(word)
	}

	return builder.String()
}
