package goutil

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_FormatTime(t *testing.T) {
	Convey("格式化时间", t, func() {
		now := time.Now()
		timeStr := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
		Convey("测试正常情况", func() {
			So(FormatTime(now), ShouldEqual, timeStr)
		})
	})
}

func Test_FormatDate(t *testing.T) {
	Convey("格式化日期", t, func() {
		now := time.Now()
		dateStr := fmt.Sprintf("%d-%02d-%02d", now.Year(), int(now.Month()), now.Day())
		Convey("格式化日期", func() {
			So(FormatDate(now), ShouldEqual, dateStr)
		})
	})
}
