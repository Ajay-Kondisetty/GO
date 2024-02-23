package main

import (
	"fmt"
	"sync"
	"time"
)

var NUM_PROCESSES = 0
var Mutex sync.Mutex

func main() {
	startTime := time.Now()
	fmt.Println("Started working")
	var wg sync.WaitGroup
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go countNumProcesses(&wg)
	}
	wg.Wait()
	endTime := time.Now()
	fmt.Println(NUM_PROCESSES, endTime.Sub(startTime))
}

func countNumProcesses(wg *sync.WaitGroup) {
	defer wg.Done()
	Mutex.Lock()
	NUM_PROCESSES++
	Mutex.Unlock()
}
