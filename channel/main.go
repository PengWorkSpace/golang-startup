package main

import (
	"fmt"
)

func main() {

	//var ch = make(chan int, 2)
	//ch <- 4
	//num := <-ch
	//fmt.Println(num)
	//close(ch)
	//fmt.Println("chan is closed")

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	//通道关闭后会退出for range循环
	for i := range ch2 {
		fmt.Println(i)
	}
}
