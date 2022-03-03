package main

import (
	"fmt"
	"sync"
)

// sumUp has two writing process which access the same non-thread-safe resource
func sumUp(start, end int) int {
	sum := 0

	wg := sync.WaitGroup{}

	addUp := func(start, end int) {
		defer wg.Done()
		for i := start; i <= end; i++ {
			sum += i
		}
	}
	wg.Add(2)

	middlePoint := end / 2
	go addUp(start, middlePoint)
	go addUp(middlePoint+1, end)

	wg.Wait()
	return sum
}

// sumUpWithInspectors has one writing process and one reading process which access the same non-thread-safe resource
func sumUpWithInspectors(start, end int) int {
	sum := 0

	wg := sync.WaitGroup{}

	addUp := func(start, end int) {
		defer wg.Done()
		for i := start; i <= end; i++ {
			sum += i
		}
	}
	inspect := func(id int) {
		defer wg.Done()
		fmt.Printf("Inspector #%d is checking the sum. Current value is: %d\n", id, sum)
	}
	wg.Add(3)

	go addUp(start, end)
	go inspect(1)
	go inspect(2)

	wg.Wait()
	return sum
}
