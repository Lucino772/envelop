package wrapper

import "sync"

type Pool struct {
	wg         sync.WaitGroup
	resultChan chan *PoolResult
	done       chan struct{}
}

type PoolResult struct {
	Key   any
	Error error
}

func NewPool() *Pool {
	return &Pool{
		resultChan: make(chan *PoolResult, 10),
		done:       make(chan struct{}),
	}
}

func (pool *Pool) WaitAsync() {
	go func() {
		defer close(pool.done)
		defer close(pool.resultChan)
		pool.wg.Wait()
	}()
}

func (pool *Pool) Go(key any, function func() error) {
	pool.wg.Add(1)
	go func(wg *sync.WaitGroup, resultChan chan<- *PoolResult, taskKey any) {
		defer wg.Done()
		err := function()
		resultChan <- &PoolResult{Key: taskKey, Error: err}
	}(&pool.wg, pool.resultChan, key)
}

func (pool *Pool) Result() <-chan *PoolResult {
	return pool.resultChan
}

func (pool *Pool) Done() <-chan struct{} {
	return pool.done
}
