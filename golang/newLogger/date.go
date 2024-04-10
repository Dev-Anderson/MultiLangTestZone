package newlogger

import "time"

func currentDate() string {
	t := time.Now()
	return t.Format("2006-01-02")
}

func now() time.Time {
	return time.Now()
}
