package swork

import (
	"sync"
)

type Processor interface {
	Process(string)
}

type WorkerGroup struct {
	sync.WaitGroup
	StringC chan string
}

func NewWorkerGroup(procs int, p Processor) *WorkerGroup {
	w := &WorkerGroup{
		StringC: make(chan string),
	}

	for i := 0; i < procs; i++ {
		w.Add(1)

		go func() {
			for line := range w.StringC {
				p.Process(line)
			}

			w.Done()
		}()
	}
	return w
}
