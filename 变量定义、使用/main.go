package main

import (
	"fmt"
	"strconv"
)

// 定义全局变量
var name = "akoncoder"
var number = 100

func main() {
	//
	//innerStr := "inner akonc_coder"
	//innerNumber := 200
	//
	//fmt.Println(name)
	//fmt.Println(number)
	//
	//fmt.Println(innerStr)
	//fmt.Println(innerNumber)

	//test()

	errReturn()
}

func test() {
	var _ int
	scope := "6"
	if code, err := strconv.Atoi(scope); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(code)
	}
}

//err check
//https://www.cnblogs.com/dajianshi/p/4211848.html
func errReturn() {
	var s = "world"
	code, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(code)
	}

	var sq = "hello"
	if code, err = strconv.Atoi(sq); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(code)
	}
}
