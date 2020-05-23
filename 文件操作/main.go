package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// 1、文件的打开关闭
	file, err := os.Open("/Users/akon_coder/go_projects/com.akoncoder.init/文件操作/input.txt")
	if err != nil {
		fmt.Println("open file error")
	}

	fmt.Printf("file=%v", file)

	closeErr := file.Close()
	if closeErr != nil {
		fmt.Println("关闭文件异常")
	}

	// 2、bufio使用
	file, err1 := os.Open("/Users/akon_coder/go_projects/com.akoncoder.init/文件操作/input.txt")
	if err1 != nil {
		fmt.Println("open file error")
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		//io.EoF 表示读到文件末尾
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}

	//3.使用ioutil读取
	filePath := "/Users/akon_coder/go_projects/com.akoncoder.init/文件操作/input.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败")
	}

	// 默认显示的是[]byte, 转换成string
	output := string(content)
	fmt.Println(output)
}
