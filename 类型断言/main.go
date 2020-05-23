package main

import "fmt"

func main() {
	var a interface{}
	var num1 float32 = 12.22
	a = num1
	//类型断言成功
	num2 := a.(float32)
	fmt.Println(num2)

	//类型断言失败
	// num3 := a.(float64)
	// fmt.Println(num3)

	//类型断言带检测机制
	num4, ok := a.(float64)
	if ok {
		fmt.Println("类型转换成功", num4)
	} else {
		fmt.Println("类型转换失败")
	}

}
