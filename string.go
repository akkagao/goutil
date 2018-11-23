package goutil

import (
	"strings"

	"github.com/jinzhu/gorm"
)

// snake string, XxYy to xx_yy , XxYY to xx_yy
func SnakeString(s string) string {
	return gorm.ToColumnName(s)
}

func SnakeStrings(ss ...string) []string {
	for index := 0; index < len(ss); index++ {
		ss[index] = gorm.ToColumnName(ss[index])
	}
	return ss
}

func UpperFirst(str string) string {
	b := []byte(str)
	first := strings.ToUpper(string(b[:1]))
	return first + string(b[1:])
}
