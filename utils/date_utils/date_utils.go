package date_utils

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	dateFormat    = "2006-01-02 15:04:05"
)

func DateTimeNow() time.Time {
	return time.Now()
}

func GetNowString() string {
	return DateTimeNow().Format(dateFormat)
}

func ApiDateLayout() string {
	return DateTimeNow().Format(apiDateLayout)
}
