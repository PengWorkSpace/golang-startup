package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter interface {
	Inc()
	Load() int64
}

// 普通版
type CommonCounter struct {
	counter int64
}

func (c CommonCounter) Inc() {
	c.counter++
}

func (c CommonCounter) Load() int64 {
	return c.counter
}

type MutexCounter struct {
	counter int64
	lock    sync.Mutex
}

//func Inc(m *MutexCounter) {
//	mutex.lock.Lock()
//	defer mutex.lock.Unlock()
//	mutex.counter ++
//}
//
//func Load(m *MutexCounter) int64  {
//
//	m.lock.Lock()
//	defer m.lock.Unlock()
//	return m.counter
//}

type AtomicCounter struct {
}

func test(c Counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}

func main() {

	c1 := CommonCounter{ counter: 0}
	test(c1)

}
