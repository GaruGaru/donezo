package tasks

import (
	"encoding/json"
	"time"
)

type Task struct {
	Time        time.Time
	Description string
	Effort      *int
	EffortType  *string
}

func (t Task) toJson() []byte {
	jsn, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return jsn
}

