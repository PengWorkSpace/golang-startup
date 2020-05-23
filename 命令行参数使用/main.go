package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("命令行参数个数：", len(os.Args))
	for index, value := range os.Args {
		fmt.Printf("args[%v]=%v", index, value)

	}
}
