package tasks

import "time"

type Repository interface {
	Add(time time.Time, task Task)
	Remove(time time.Time)
	List(time time.Time) []Task
}
