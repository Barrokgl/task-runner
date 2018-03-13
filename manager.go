package taskRunner

import (
	"bitbucket.org/GromKri/go-qiwi-service/model"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type TaskManager struct {
	db     *gorm.DB
	logger *logrus.Logger
	Runner *Runner
	store  TaskStore
}

const (
	STATUS_STARTED = "STARTED"
	STATUS_PENDING = "PENDING"
	STATUS_DONE    = "DONE"
	STATUS_ERROR   = "ERROR"
)

func NewTaskManager(db *gorm.DB, logger *logrus.Logger, store TaskStore) *TaskManager {
	return &TaskManager{
		db:     db,
		logger: logger,
		Runner: NewRunner(logger),
		store:  store,
	}
}

type WorkContstructor func(m *TaskManager, task *model.SystemTask) *Job
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
			m.logger.Warningln("unsupported task type: ", t.WorkType)
		}
	}

	m.Runner.jobs = jobs
	m.Runner.Start()

	return nil
}

func (m *TaskManager) Stop() {
	m.Runner.Stop()
}

//func (m *TaskManager) AddPaymentCheck(paymentId int64, timeout time.Time) error {
//	task, err := m.store.Add(model.SystemTask{
//		TimeOut:    timeout,
//		WorkType:   TASK_PAYMENT_CHECK,
//		Status:     STATUS_STARTED,
//		PayloadInt: paymentId,
//	})
//	if err != nil {
//		return err
//	}
//
//	m.Runner.AddWork(timeout, &PaymentCheckTask{
//		ID:        int64(task.ID),
//		db:        m.db,
//		logger:    m.logger,
//		paymentId: paymentId,
//		startTime: time.Now(),
//		timeout:   timeout,
//	})
//
//	util.PrettyPrint(m.Runner.Jobs(), m.logger)
//
//	return nil
//}
