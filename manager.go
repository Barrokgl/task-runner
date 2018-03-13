package taskRunner

type TaskManager struct {
	logger Logger
	Runner *Runner
	store  TaskStore
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
		store:  store,
	}
}

type WorkContstructor func(m *TaskManager, task *PersistedTask) *Job
type WorkMap map[string]WorkContstructor

func (m *TaskManager) Initialize(workMap WorkMap) error {
	tasks, err := m.store.Fetch()
	if err != nil {
		return err
	}

	jobs := []*Job{}
	for _, t := range tasks {
		if constructor, ok := workMap[t.WorkType]; ok && constructor != nil {
			jobs = append(jobs, constructor(m, &t))
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