package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func startTask(wg *sync.WaitGroup, taskCh chan Task, errorCounter *atomic.Int32, errorLimit int, ignoreErrors bool) {
	defer wg.Done()

	for {
		if !ignoreErrors && errorCounter.Load() >= int32(errorLimit) {
			return
		}

		task, opened := <-taskCh
		if !opened {
			return
		}
		err := task()
		if !ignoreErrors && err != nil {
			errorCounter.Add(1)
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
// n - кол-во горутин
// m - кол-во допустимых ошибок; если m <= 0 - игнорировать ошибки.
func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}
	errorCount := atomic.Int32{}
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
	if !ignoreErrors && errorCount.Load() >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
