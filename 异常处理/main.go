package main

import "fmt"

func main() {

	numDiv()

	fmt.Println("main 执行完毕······")
}

//采用defer + recover 来恢复
//相当于java、C#语言的try····catch····finally
func numDiv() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("程序存在异常·····")
		}

	}()

	//尝试除以0操作
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}
