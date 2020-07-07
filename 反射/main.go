package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {

	v := reflect.TypeOf(x)
	fmt.Printf("type is: %v\n", v)
}

func main() {
	name := 3.14
	reflectType(name)

	str := "akoncoder"
	reflectType(str)
}
