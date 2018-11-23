package goutil

import "time"

func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
