package goutil

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetMyName(t *testing.T) {
	result := GetMyName()
	fmt.Println(result)
	Convey("获取当前方法的名称", t, func() {
		Convey("获取当前方法的名称", func() {
			So(result, ShouldContainSubstring, "goutil.Test_GetMyName")
			So(result, ShouldEqual, "github.com/akkagao/goutil.Test_GetMyName")
		})
	})
}

func call() string {
	funcName := GetCallerName()
	fmt.Println("=======")
	fmt.Println(funcName)
	fmt.Println("=======")
	return funcName
}
func Test_GetCallerName(t *testing.T) {
	result := call()
	Convey("获取调用者的方法名称", t, func() {
		Convey("获取调用者的方法名称", func() {
			So(result, ShouldContainSubstring, "Test_GetCallerName")
			So(result, ShouldEqual, "github.com/akkagao/goutil.Test_GetCallerName")
		})
	})

}
