package tasks

import "time"

func KeyByDay(time time.Time) string {
	return time.Format("02-01-2006")
}
