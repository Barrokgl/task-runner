// inspired by github.com/robfig/cron

package taskRunner

import (
	"runtime/debug"
	"sort"
	"time"
)

type Runner struct {
	jobs     []*Job
	stop     chan struct{}
	add      chan *Job
	snapshot chan []*Job
	running  bool
	location *time.Location
	Log      Logger
}

type Work interface {
	Run()
	GetID() int64
}

type Job struct {
	Schedule Schedule
	Next     time.Time
	IsError  bool
	Work     Work
	Async bool
}

type byTime []*Job

func (s byTime) Len() int      { return len(s) }
func (s byTime) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byTime) Less(i, j int) bool {
	if s[i].Next.IsZero() {
		return false
	}
	if s[j].Next.IsZero() {
		return true
	}
	return s[i].Next.Before(s[j].Next)
}

func NewRunner(logger Logger) *Runner {
	return NewRunnerWithLocation(time.Now().Location(), logger)
}

func NewRunnerWithLocation(location *time.Location, logger Logger) *Runner {
	return &Runner{
		jobs:     nil,
		add:      make(chan *Job),
		stop:     make(chan struct{}),
		snapshot: make(chan []*Job),
		running:  false,
		location: location,
		Log:      logger,
	}
}

func (r *Runner) AddWork(t time.Time, work Work) {
	r.Schedule(NewScheduled(t), work)
}

func (r *Runner) Schedule(schedule Schedule, work Work) {
	job := &Job{
		Schedule: schedule,
		Work:     work,
	}

	if !r.running {
		r.jobs = append(r.jobs, job)
		return
	}

	r.add <- job
}

func (r *Runner) Jobs() []*Job {
	if r.running {
		r.snapshot <- nil
		jobs := <-r.snapshot
		return jobs
	}
	return r.jobSnapshot()
}

func (r *Runner) Location() *time.Location {
	return r.location
}

func (r *Runner) Start() {
	if r.running {
		return
	}
	r.running = true
	go r.run()
}

func (r *Runner) Run() {
	if r.running {
		return
	}
	r.running = true
	r.run()
}

func (r *Runner) Stop() {
	if !r.running {
		return
	}

	r.stop <- struct{}{}
	r.running = false
}

func (run *Runner) runWithRecovery(work Work) {
	defer func() {
		if r := recover(); r != nil {
			run.Log.Println("panic running job: ", r)
			run.Log.Println(string(debug.Stack()))
		}
	}()
	work.Run()
}

func (r *Runner) run() {
	now := r.now()
	for _, job := range r.jobs {
		job.Next = job.Schedule.Next(now)
	}

	for {

		sort.Sort(byTime(r.jobs))

		var timer *time.Timer

		if len(r.jobs) == 0 || r.jobs[0].Next.IsZero() {
			timer = time.NewTimer(100 * time.Hour)
		} else {
			timer = time.NewTimer(r.jobs[0].Next.Sub(now))
		}

		for {
			select {
			case now = <-timer.C:
				now = now.In(r.location)
				// TODO: job error indicator
				for _, job := range r.jobs {
					if job.Next.After(now) || job.Next.IsZero() {
						break
					}

					if job.Async {
						go r.runWithRecovery(job.Work)
					} else {
						r.runWithRecovery(job.Work)
					}

					r.jobs = r.jobs[1:len(r.jobs)]
				}
			case newJob := <-r.add:
				timer.Stop()
				now = r.now()
				newJob.Next = newJob.Schedule.Next(now)
				r.jobs = append(r.jobs, newJob)

			case <-r.snapshot:
				r.snapshot <- r.jobSnapshot()
				continue

			case <-r.stop:
				timer.Stop()
				return
			}

			break
		}
	}
}

func (r *Runner) jobSnapshot() []*Job {
	jobs := []*Job{}
	for _, j := range r.jobs {
		jobs = append(jobs, &Job{
			Schedule: j.Schedule,
			Next:     j.Next,
			IsError:  j.IsError,
			Work:     j.Work,
			Async:    j.Async,
		})
	}

	return jobs
}

func (r *Runner) now() time.Time {
	return time.Now().In(r.location)
}
