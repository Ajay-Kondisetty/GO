package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type WorkerPool struct {
	workFunc chan func()

	numberOfWorker int

	wg sync.WaitGroup
}

func (w *WorkerPool) InitWorkerPool(numberOfWorker int) {
	w.numberOfWorker = numberOfWorker

	for i := 0; i < w.numberOfWorker; i++ {
		w.wg.Add(1)
		go w.Worker()
	}

}

func main() {
	fmt.Println("Number of active go routines at the beginning: ", getActiveGoroutinesCount())
	workerPool := new(WorkerPool)

	workerPool.wg = sync.WaitGroup{}
	workerPool.workFunc = make(chan func())
	workerPool.InitWorkerPool(2)
	fmt.Println("Number of active go routines after initializing worker pool: ", getActiveGoroutinesCount())
	for i := 0; i < 10; i++ {
		workerPool.AddWork()
	}

	workerPool.FinishWork()
	fmt.Println("Number of active go routines after closing worker pool: ", getActiveGoroutinesCount())
}

func (w *WorkerPool) Worker() {
	for workFunc := range w.workFunc {
		workFunc()
	}
	w.wg.Done()
}

func (w *WorkerPool) FinishWork() {
	close(w.workFunc)
	w.wg.Wait()
}

func someWork() {
	fmt.Println("Work started.")
	time.Sleep(2 * time.Second)
	fmt.Println("Work complete.")
}

func (w *WorkerPool) AddWork() {
	w.workFunc <- someWork
}

func getActiveGoroutinesCount() int {
	return runtime.NumGoroutine()
}
