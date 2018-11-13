package goutil

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func Test_FormatNumber(t *testing.T) {
	Convey("格式化数字,位数不足前面补零", t, func() {
		Convey("不足4位前面补充0", func() {
			So(FormatNumber(1, 4), ShouldEqual, "0001")
		})
		Convey("足够4位不用补充", func() {
			So(FormatNumber(1000, 4), ShouldEqual, "1000")
		})
		Convey("大于四位", func() {
			So(FormatNumber(10000, 4), ShouldEqual, "10000")
		})
		Convey("小于0，负号也占位", func() {
			So(FormatNumber(-11, 4), ShouldEqual, "-011")
		})
		Convey("等于0", func() {
			So(FormatNumber(0, 4), ShouldEqual, "0000")
		})
	})
}
