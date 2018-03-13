package taskRunner

import "time"

type Schedule interface {
	Next(time.Time) time.Time
}

type Scheduled struct {
	actionTime time.Time
}

func (d Scheduled) Next(time.Time) time.Time {
	return d.actionTime
}

func NewScheduled(timeout time.Time) Scheduled {
	return Scheduled{
		actionTime: timeout,
	}
}
