package gopp

import (
	"fmt"
	"strconv"
	"time"
)

// 中文常用格式
func TimeToFmt1(t time.Time) string { return t.Format("2006-01-02 15:04:05") }

// rounded float point part
// origin 8h38m46.115296675s
// now 8h38m46.115s
func Duround(d time.Duration) string {
	return d.String()
}

const CleanDateFmt = "2006-01-02 15:04:05"
const HttpDateFmt = "Mon, 02 Jan 2006 15:04:05 GMT" // "Sat, 30 Sep 2017 00:10:59 GMT"

var StartTime = time.Now()

func Dur2hum(d time.Duration) string {
	unitMeasures := []time.Duration{365 * 24 * time.Hour, 12 * 24 * time.Hour, 24 * time.Hour, time.Hour, time.Minute, time.Second, time.Millisecond, time.Microsecond, time.Nanosecond}
	// unitWords := []string{"year", "month", "day", "hour", "minute", "second", "millisecond", "microsecond", "nanosecond"}
	unitShorts := []string{"y", "M", "d", "h", "m", "s", "ms", "µs", "ns"}

	r := ""
	for idx, du := range unitMeasures {
		m := d.Nanoseconds() / du.Nanoseconds()
		if m == 0 {
		} else {
			r += fmt.Sprintf("%d%s", m, unitShorts[idx])
		}
		d -= time.Duration(m) * du
		if idx >= 6 {
			break
		}
	}

	return r
}

// offset [0-14)
func SetTimezone(offset int) (olocal *time.Location) {
	olocal = time.Local
	if offset >= 0 && offset < 24 {
		secondsEastOfUTC := int((time.Duration(offset) * time.Hour).Seconds())
		zone := time.FixedZone(fmt.Sprintf("UTC+%d", offset), secondsEastOfUTC)
		time.Local = zone
	}
	return
}
func SetLocal(offset int) *time.Location { return SetTimezone(offset) }

func CondWait(timeoutms int, f func() bool) {
	for {
		time.Sleep(time.Duration(timeoutms) * time.Millisecond)
		if f() {
			break
		}
	}
}

func TimeUnixMS(t time.Time) int64 {
	return t.Unix()*1000 + int64(t.Nanosecond()/1000000)
}
func TimeUnixMSStr(t time.Time) string { return fmt.Sprintf("%d", TimeUnixMS(t)) }
func TimeUnixMSNow() int64 {
	t := time.Now()
	return t.Unix()*1000 + int64(t.Nanosecond()/1000000)
}
func TimeUnixMSStrNow() string { return fmt.Sprintf("%d", TimeUnixMSNow()) }
func TimeFromUnixMS(tsms int64) time.Time {
	return time.Time(time.Unix(tsms/1000, tsms%1000*1000000))
}
func TimeFromUnixMSStr(tsms string) (time.Time, error) {
	ts, err := strconv.ParseInt(tsms, 10, 64)
	return TimeFromUnixMS(ts), err
}
