package goutil

import (
	"fmt"
	"testing"
)

func TestJsonMarshal(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	user := User{
		Id:   1,
		Name: "CrazyWolf",
	}
	fmt.Println(JsonMarshal(user))
}
