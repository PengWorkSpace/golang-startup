package main

import "fmt"

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

}

//结构体类型
type person struct {
	age  int
	name string
}
