package main

import (
	"fmt"
	"sync"
)

func sumUpWithLock(start, end int) int {
	sum := 0

	wg := sync.WaitGroup{}
	mux := sync.Mutex{}

	addUp := func(start, end int) {
		defer wg.Done()
		for i := start; i <= end; i++ {
			mux.Lock()
			sum += i
			mux.Unlock()
		}
	}
	wg.Add(2)

	middlePoint := end / 2
	go addUp(start, middlePoint)
	go addUp(middlePoint+1, end)

	wg.Wait()
	return sum
}

func sumUpWithRWLock(start, end int) int {
	sum := 0

	wg := sync.WaitGroup{}
	rwMux := sync.RWMutex{}

	addUp := func(start, end int) {
		defer wg.Done()
		for i := start; i <= end; i++ {
			rwMux.Lock()
			sum += i
			rwMux.Unlock()
		}
	}

	inspect := func(id int) {
		defer wg.Done()
		rwMux.RLock()
		fmt.Printf("Inspector #%d is checking the sum. Current value is: %d\n", id, sum)
		rwMux.RUnlock()
	}

	wg.Add(4)

	middlePoint := end / 2
	go addUp(start, middlePoint)
	go addUp(middlePoint+1, end)
	go inspect(1)
	go inspect(2)

	wg.Wait()
	return sum
}
