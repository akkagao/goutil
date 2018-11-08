package goutil

import (
	"strconv"
	"strings"
)

func IntArrToStrArr(intArray []int) []string {
	s := make([]string, len(intArray))
	for i, v := range intArray {
		s[i] = strconv.Itoa(v)
	}
	return s
}

func IntArrToStrArrAppend(intArray []int) []string {
	s := make([]string, 0, len(intArray))
	for _, v := range intArray {
		s = append(s, strconv.Itoa(v))
	}
	return s
}

func IntArrToString(intArray []int) string {
	return strings.Join(IntArrToStrArr(intArray), "")
}
