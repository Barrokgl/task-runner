package taskRunner

import "errors"

type TaskManager struct {
	logger Logger
	Runner *Runner
	Store  TaskStore
	jobMap JobMap
}

const (
	STATUS_STARTED = "STARTED"
	STATUS_PENDING = "PENDING"
	STATUS_DONE    = "DONE"
	STATUS_ERROR   = "ERROR"
)

type Logger interface {
	Println(a ...interface{})
}

func NewTaskManager(logger Logger, store TaskStore) *TaskManager {
	return &TaskManager{
		logger: logger,
		Runner: NewRunner(logger),
		Store:  store,
		jobMap: make(JobMap),
	}
}

type JobConstructor func(m *TaskManager, task *PersistedTask) *Job
type JobMap map[string]JobConstructor

func (m *TaskManager) Initialize(workMap JobMap) error {
	tasks, err := m.Store.Fetch()
	if err != nil {
		return err
	}

	jobs := []*Job{}
	for _, t := range tasks {
		if constructor, ok := workMap[t.WorkType]; ok && constructor != nil {
			job := constructor(m, &t)
			if job != nil {
				jobs = append(jobs, job)
			}
		} else {
			m.logger.Println("unsupported task type: ", t.WorkType)
		}
	}

	m.Runner.jobs = jobs
	m.Runner.Start()

	return nil
}

func (m *TaskManager) Stop() {
	m.Runner.Stop()
}

func (m *TaskManager) AddTask(task PersistedTask) (*PersistedTask, error) {
	constructor, ok := m.jobMap[task.WorkType]
	if !ok || constructor == nil {
		return nil, errors.New("Unkown work type")
	}

	persistedTask, err := m.Store.Add(task)
	if err != nil {
		return nil, err
	}

	m.Runner.AddJob(task.TimeOut, constructor(m, persistedTask))
	return &task, nil
}