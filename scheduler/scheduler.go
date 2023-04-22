package scheduler

import (
	"sync"

	manager "github.com/bot-games/game-manager"
)

type Scheduler struct {
	cnt uint8
	mtx sync.Mutex
	cb  func()
}

func New() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) Notify(ev manager.SchedulerEvent) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	switch ev {
	case manager.SchedulerEventJoin:
		s.cnt++
	case manager.SchedulerEventLeave:
		if s.cnt == 0 {
			panic("Something went wrong")
		}
		s.cnt--
	default:
		panic("Unknown SchedulerEvent value")
	}

	if s.cnt == 2 {
		s.cnt = 0
		s.cb()
	}
}

func (s *Scheduler) SetOnReady(callback func()) {
	s.cb = callback
}
