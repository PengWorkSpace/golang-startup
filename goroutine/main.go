package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	//go hello()
	//fmt.Println("main goroutine")
	//time.Sleep(5 * time.Second)

	//启动多个goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}

	wg.Wait()
}

//func hello() {
//	time.Sleep(3 * time.Second)
//	fmt.Println("hello goroutine")
//}

func hello(i int)  {
	wg.Done()
	fmt.Println("hello goroutine,", i)
}

