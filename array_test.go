package goutil

import (
	"testing"

	"fmt"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_IntArrToStrArr(t *testing.T) {
	Convey("int slice转换成字符串slice", t, func() {
		Convey("正常数据测试", func() {
			intArray := []int{1, 2}
			result := IntArrToStrArr(intArray)
			fmt.Println("Test_IntArrToStrArr result:", result)
			So(len(result), ShouldEqual, len(intArray))
		})

		Convey("空数组测试", func() {
			intArray := []int{}
			result := IntArrToStrArr(intArray)
			fmt.Println("Test_IntArrToStrArr result:", result)
			So(len(result), ShouldEqual, len(intArray))
		})
	})
}

func Test_IntArrToStrArrAppend(t *testing.T) {

	Convey("int slice转换成字符串slice 使用appen方法实现", t, func() {
		Convey("正常数据测试", func() {
			intArray := []int{1, 2}
			result := IntArrToStrArrAppend(intArray)
			fmt.Println("Test_IntArrToStrArrAppend result:", result)
			So(len(result), ShouldEqual, len(intArray))
		})

		Convey("空数组测试", func() {
			intArray := []int{}
			result := IntArrToStrArrAppend(intArray)
			fmt.Println("Test_IntArrToStrArrAppend result:", result)
			So(len(result), ShouldEqual, len(intArray))
		})
	})
}

func Test_IntArrToString(t *testing.T) {
	Convey("int slice 转换成字符串", t, func() {
		Convey("正常数据测试", func() {
			intArray := []int{1, 2}
			result := IntArrToString(intArray)
			fmt.Println("Test_IntArrToString result:", result)
			So(result, ShouldEqual, "12")
		})
		Convey("空数组测试", func() {
			intArray := []int{}
			result := IntArrToString(intArray)
			fmt.Println("Test_IntArrToString result:", result)
			So(result, ShouldEqual, "")
		})
	})
}
