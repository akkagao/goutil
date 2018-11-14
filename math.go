package goutil

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/**
格式化数字number，如果位数小于len则前面补零
*/
func FormatNumber(number, len int) string {
	// 实际使用%0N 格式化数字，位数不足补零
	// fmt.Sprintf("%0Xd", number)
	return fmt.Sprintf("%0"+strconv.Itoa(len)+"d", number)
}

func RandomNumber(len int) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(10000)
	return r
}
