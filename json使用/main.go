package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	boy := person{
		ID:      1,
		Age:     28,
		Name:    "akoncoder",
		Address: "上海市",
	}
	jsonStr, err := json.Marshal(&boy)
	if err != nil {
		fmt.Println("json 序列化失败")
	}

	fmt.Printf("序列化后字符串=%v\n", string(jsonStr))

	var boyStr person
	err1 := json.Unmarshal([]byte(string(jsonStr)), &boyStr)
	if err1 != nil {
		fmt.Println("反序列化失败")
	}

	fmt.Println(boyStr)
}

//使用tag `json:"name"` 指定返回给客户端使用的格式
type person struct {
	ID      int    `json:"id"`
	Age     int    `json:"age"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
