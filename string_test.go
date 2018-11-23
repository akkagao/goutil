package goutil

import (
	"fmt"
	"testing"
)

func TestSnakeStrings(t *testing.T) {
	fmt.Println(SnakeStrings("FirstName", "SSO", "num_of_Dogs"))
}

func TestUpperFirst(t *testing.T) {
	fmt.Println(UpperFirst("aa"))
}
