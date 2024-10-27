package cronjob

import (
	"fmt"
	"log"
	"sync"

	"github.com/robfig/cron/v3"
)

type Task struct {
	Job  cron.Job
	Cron string
	ID   cron.EntryID
}

type Scheduler struct {
	cron  *cron.Cron
	tasks map[string]*Task
	lock  sync.Mutex
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		cron:  cron.New(cron.WithSeconds()),
		tasks: make(map[string]*Task),
	}
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) AddTask(name string, task *Task) error {
	s.lock.Lock()
	if _, ok := s.tasks[name]; ok {
		s.lock.Unlock()
		return fmt.Errorf("task name exists %s", name)
	}
	s.lock.Unlock()
	id, err := s.cron.AddJob(task.Cron, task.Job)
	if err != nil {
		log.Println("Error scheduling task", err)
		return nil
	}
	task.ID = id
	s.tasks[name] = task
	return nil
}

func (s *Scheduler) RemoveTask(name string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	task, ok := s.tasks[name]
	if ok {
		s.cron.Remove(task.ID)
		delete(s.tasks, name)
	}
}
