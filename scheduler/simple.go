package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {s.workerChan <- r}()
}

func (s *SimpleScheduler) ConfigureMastWorkerChan(c chan engine.Request) {
	s.workerChan = c
}


