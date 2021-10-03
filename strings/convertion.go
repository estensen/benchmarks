package strings

import (
	"fmt"
	"strconv"
)

func ConvertFmt(number int) string {
	return fmt.Sprint(number)
}

func ConvertStrconv(number int) string {
	return strconv.Itoa(number)
}
