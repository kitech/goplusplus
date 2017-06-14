package gopp

import "time"

// 中文常用格式
func TimeToFmt1(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// rounded float point part
// origin 8h38m46.115296675s
// now 8h38m46.115s
func Duround(d time.Duration) string {
	return d.String()
}
