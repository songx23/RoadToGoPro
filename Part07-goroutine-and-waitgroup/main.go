package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// sequential processing
	seqStart := time.Now()
	for i := 0; i < 6; i++ {
		hardCalculation(i)
	}
	seqDur := time.Since(seqStart).Milliseconds()
	fmt.Printf("Total time spent: %d ms\n", seqDur)
	// async processing
	asyncStart := time.Now()
	wg := &sync.WaitGroup{}
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go hardCalculationAsync(wg, i)
	}
	wg.Wait()
	asyncDur := time.Since(asyncStart).Milliseconds()
	fmt.Printf("Total time spent: %d ms\n", asyncDur)
}

func hardCalculation(input int) {
	// pretending I'm working on super hard math problems
	time.Sleep(10 * time.Second)
	fmt.Printf("math problem solved for input: %d.\n", input)
}

func hardCalculationAsync(wg *sync.WaitGroup, input int) {
	// pretending I'm working on super hard math problems
	defer wg.Done()
	time.Sleep(10 * time.Second)
	fmt.Printf("math problem solved for input: %d.\n", input)
}
