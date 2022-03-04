package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	messageChannel := make(chan int, 10)
	wg := sync.WaitGroup{}
	// consumer
	numberOfWorkers := 10
	w := newWorker(numberOfWorkers, messageChannel, &wg)
	w.start()

	// producer
	maxParam := 100

	for i := 1; i <= maxParam; i++ {
		w.process(i)
	}

	wg.Wait()
}

type worker struct {
	queue       chan int
	waitGroup   *sync.WaitGroup
	workerCount int
}

func newWorker(workerCount int, q chan int, wg *sync.WaitGroup) *worker {
	return &worker{
		queue:       q,
		waitGroup:   wg,
		workerCount: workerCount,
	}
}

func (w *worker) start() {
	for i := 0; i < w.workerCount; i++ {
		go w.startWorker()
	}
}

func (w *worker) startWorker() {
	for {
		select {
		case param, ok := <-w.queue:
			if !ok {
				fmt.Println("Batch worker is stopping...")
				return
			}
			w.hardCalculation(param)
		}
	}
}

func (w *worker) hardCalculation(input int) {
	// pretending I'm working on super hard math problems
	defer w.waitGroup.Done()
	time.Sleep(5 * time.Millisecond)
	//comment out to reduce log noise
	//fmt.Printf("math problem solved for input: %d.\n", input)
}

func (w *worker) process(param int) {
	//comment out to reduce log noise
	//fmt.Printf("Producing with parameter: %d\n", param)
	w.waitGroup.Add(1)
	w.queue <- param
}
