package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMastWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParserResult)
	e.Scheduler.ConfigureMastWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		creatWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	count := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got #%d item: %v", count, item)
			count++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func creatWorker(in chan Request, out chan ParserResult) {
	go func() {
		for {
			request := <-in
			parserResult, err := worker(request)
			if err != nil {
				continue
			}

			out <- parserResult
		}
	}()
}
