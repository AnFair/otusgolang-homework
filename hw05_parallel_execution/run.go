package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func startTask(wg *sync.WaitGroup, taskCh chan Task, errorCounter *int32, errorLimit int, ignoreErrors bool) {
	defer wg.Done()

	for {
		if !ignoreErrors && *errorCounter >= int32(errorLimit) {
			fmt.Println("exiting due to exceeding error limit")
			return
		}

		task, opened := <-taskCh
		if !opened {
			return
		}
		err := task()
		if !ignoreErrors && err != nil {
			atomic.AddInt32(errorCounter, 1)
			fmt.Println("errorCounter is ", *errorCounter)
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
// n - кол-во горутин
// m - кол-во допустимых ошибок; если m <= 0 - игнорировать ошибки.
func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}
	errorCount := int32(0)
	taskCh := make(chan Task, len(tasks))
	ignoreErrors := m <= 0

	for _, task := range tasks {
		taskCh <- task
	}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go startTask(&wg, taskCh, &errorCount, m, ignoreErrors)
	}
	close(taskCh)
	wg.Wait()
	if !ignoreErrors && errorCount >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
