package strings

import (
	"fmt"
	"strings"
)

func FileSplit(name string) string {
	parts := strings.Split(name, "/")
	return parts[0]
}

func FileCut(name string) string {
	filename, _, _ := strings.Cut(name, ".")
	return filename
}

func FileConcatPrint(dir, filename string) string {
	return fmt.Sprintf("%s/%s", dir, filename)
}

func FileConcat(dir, filename string) string {
	return dir + "/" + filename
}
