package time

import "time"

func NowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TimeFromString(str string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
}

func TimestampToString(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
