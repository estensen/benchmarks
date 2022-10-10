package strings

import "strings"

func FileSplit(name string) string {
	parts := strings.Split(name, "/")
	return parts[0]
}

func FileCut(name string) string {
	filename, _, _ := strings.Cut(name, ".")
	return filename
}
