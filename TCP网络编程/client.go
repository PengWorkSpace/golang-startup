package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("conn error", err)
	}

	fmt.Printf("connect success,conn:%v,ip=%v", conn, conn.RemoteAddr().String())
}
