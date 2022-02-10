package main

import (
	"sync"
	"testing"
)

func BenchmarkWorker_ProcessSetupTest_1(b *testing.B)    { benchmarkProcess(1, b) }
func BenchmarkWorker_ProcessSetupTest_10(b *testing.B)   { benchmarkProcess(10, b) }
func BenchmarkWorker_ProcessSetupTest_100(b *testing.B)  { benchmarkProcess(100, b) }
func BenchmarkWorker_ProcessSetupTest_1000(b *testing.B) { benchmarkProcess(1000, b) }

func benchmarkProcess(workerCount int, b *testing.B) {
	wg := &sync.WaitGroup{}
	worker := newWorker(workerCount, make(chan int, workerCount), wg)
	worker.start()

	for i := 0; i < b.N; i++ {
		worker.process(i)
	}

	// Stop the worker
	wg.Wait()
}
