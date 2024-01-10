package utils

import (
	"strings"
	"time"
)

// The timestamp now. In seconds. Local.
func TimestampNow() int {
	return int(time.Now().Unix())
}

// Deprecated.
func GetTimeRangeByType(t string, location string) (start int64, end int64) {
	switch t {
	case "day":
		start = GetTodayStart(location).Unix()
		end = start + 24*3600
	case "week":
		start = GetWeekStart(location).Unix()
		end = start + 24*3600*8
	case "month":
		start = GetMonthStart(location).Unix()
		end = start + 24*3600*30
	case "daybefore":
		start = GetTodayStart(location).Unix() - 24*3600
		end = start + 24*3600
	case "year":
		start = GetYearStart(location).Unix() - 24*3600
		end = start + 12*24*3600*30
	default:
		start = GetTodayStart(location).Unix()
		end = start + 24*3600
	}

	return
}

// Deprecated.
func TimeFormat(timestamp int64, location string) string {
	locationTime, _ := time.LoadLocation(location)
	timeStr := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	res, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, locationTime)
	ar := strings.Split(res.String(), " ")
	return ar[0] + " " + ar[1]
}

// Deprecated.
func TimeFormatTime(timestamp int64, location string) time.Time {
	locationTime, _ := time.LoadLocation(location)
	timeStr := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	res, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, locationTime)
	return res
}

// Return the local time.
func TimeNowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Deprecated.
func GetTodayStart(location string) time.Time {
	locationTime, _ := time.LoadLocation(location)
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, locationTime)

	return t
}

// Deprecated.
func GetWeekStart(location string) time.Time {
	locationTime, _ := time.LoadLocation(location)
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, locationTime)

	return time.Unix(t.Unix()-7*24*3600, 0)
}

// Deprecated.
func GetMonthStart(location string) time.Time {
	locationTime, _ := time.LoadLocation(location)
	timeStr := time.Now().Format("2006-01")
	t, _ := time.ParseInLocation("2006-01", timeStr, locationTime)

	return t
}

// Deprecated.
func GetYearStart(location string) time.Time {
	locationTime, _ := time.LoadLocation(location)
	timeStr := time.Now().Format("2006")
	t, _ := time.ParseInLocation("2006", timeStr, locationTime)

	return t
}

// Deprecated.
func GetTimeByStr(str string, location string) time.Time {
	locationTime, _ := time.LoadLocation(location)
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", str, locationTime)
	return t
}
