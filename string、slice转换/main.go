package main

import "fmt"

func main()  {

	str:= "akoncoder@163.com"
	fmt.Println("转换前string",str)
	sliceStr := []byte(str)
	sliceStr[0]= 'z'
	fmt.Println(sliceStr)
	str = string(sliceStr)
	fmt.Println("转换后string",str)
}