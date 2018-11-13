package goutil

import (
	"fmt"
	"strconv"
)

/**
格式化数字number，如果位数小于len则前面补零
*/
func FormatNumber(number, len int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(len)+"d", number)
}
