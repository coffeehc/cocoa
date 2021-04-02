package foundation

import (
	"math"
	"time"
)

func DateToGoTime(date Date) time.Time {
	secs := date.TimeIntervalSince1970()
	modf, frac := math.Modf(secs)
	return time.Unix(int64(modf), int64(frac*1000000000))
}

func GoTimeToDate(time time.Time) Date {
	secs := float64(time.Second()) + float64(time.Nanosecond())/1000000000
	return DateWithTimeIntervalSince1970(secs)
}
