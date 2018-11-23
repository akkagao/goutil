package goutil

import (
	"strings"
)

func UpperFirst(str string) string {
	b := []byte(str)
	first := strings.ToUpper(string(b[:1]))
	return first + string(b[1:])
}
