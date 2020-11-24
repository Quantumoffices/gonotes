package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close() // 结束时，关闭连接

	for {
		// 发送数据
		_, err = conn.Write([]byte("Are u ready?"))
		if err != nil {
			fmt.Println("Write err:", err)
			return
		}
		time.Sleep(time.Minute)
	}

}
