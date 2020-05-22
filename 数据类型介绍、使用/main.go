package main

import (
	"fmt"
	"strconv"
)

func main() {
	//bool类型
	b := false
	fmt.Println(b)

	//数值类型
	num := 32
	fmt.Println(num)

	//字符串类型
	str := "akoncoder"
	fmt.Println(str)

	//数据类型
	arrays := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrays)

	//结构体使用
	p := person{age: 23, name: "akoncoder"}
	fmt.Printf("age:%v\n"+"name:%v\n", p.age, p.name)

	//切片
	slice := []int{3, 4, 5, 5, 5, 6, 6}
	sliceNew := slice[1:4]
	fmt.Println(slice)
	fmt.Println(sliceNew)

	sliceArrays := make([]int, 5, 10)
	sliceArrays[1] = 10
	fmt.Println(sliceArrays)

	//Map
	dict := make(map[string]int)
	dict["zhangsan"] = 23
	dict["lisi"] = 30
	fmt.Println(dict)

	//删除Map元素
	delete(dict, "lisi")
	fmt.Println(dict)

	//方法使用
	p.showinfo()

	//函数使用
	display("akoncoder", 28)

	//interface 使用

	sayStr := p.Say("hi")
	walkStr := p.Walk("walk")
	fmt.Println(sayStr)
	fmt.Println(walkStr)
}

//结构体类型
type person struct {
	age  int
	name string
}

// 方法，归属于person
func (p person) showinfo() {
	fmt.Println("name is " + p.name + "; age is " + strconv.Itoa(p.age))
}

//函数
func display(name string, age int) {

	fmt.Println("name is " + name + ";" + "age is " + strconv.Itoa(age))
}

type rabbitMachine interface {
	Say(s string) string
	Walk(w string) string
}

func (p person) Say(s string) string {

	return "i am " + p.name + "; i can say " + s
}

func (p person) Walk(w string) string {
	return "i am " + p.name + "; i can " + w
}
