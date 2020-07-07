package main

import (
	"fmt"
	"strconv"
	"sync"
)

//var (
//	x      float64
//	wg     sync.WaitGroup
//	lock   sync.Mutex
//	rwLock sync.RWMutex
//)

func main() {

	//ch := make(chan int, 1)
	//for i := 0; i < 10; i++ {
	//
	//	select {
	//	case x := <-ch:
	//		fmt.Println(x)
	//	case ch <- i:
	//	}
	//}

	//开启了两个goroutine去累加变量x的值，
	//这两个goroutine在访问和修改x变量的时候就会存在数据竞争，
	//导致最后的结果与期待的不符。

	// 方法一： 使用mutex 互斥锁解决资源竞争
	//wg.Add(2)
	//go add()
	//go add()
	//wg.Wait()
	//fmt.Println(x)

	//start := time.Now()
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go write()
	//}
	//
	//for i := 0; i < 100; i++ {
	//	wg.Add(1)
	//	go read()
	//}
	//
	//wg.Wait()
	//end := time.Now()
	//fmt.Println(end.Sub(start))

	//sync.Map 使用

	var m = sync.Map{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("key is : %v,value is %v \n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//func write() {
//	rwLock.Lock()
//	x = x + 1
//	time.Sleep(10 * time.Millisecond)
//	rwLock.Unlock()
//	wg.Done()
//}
//
//func read() {
//	rwLock.RLock()
//	time.Sleep(1 * time.Millisecond)
//	rwLock.RUnlock()
//	wg.Done()
//}
//
//func add() {
//	for i := 0; i < 500; i++ {
//		lock.Lock()
//		x = x + 1
//		lock.Unlock()
//	}
//	wg.Done()
//}
