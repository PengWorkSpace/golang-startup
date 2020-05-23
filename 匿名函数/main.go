package main

import "fmt"

func main() {

	//匿名直接复制使用
	num := func(a int32, b int32) int32 {
		return a + b
	}(10, 20)

	fmt.Println(num)

	//匿名函数复制给变量，通过变量使用
	newNum := func(a int32, b int32) int32 {
		return a + b
	}
	result := newNum(10, 30)
	fmt.Println(result)
}
