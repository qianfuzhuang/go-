package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//客户端
	fmt.Println("客户端启动")
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	fmt.Println("链接成功", conn)
	//终端输入数据放入缓冲区
	date := bufio.NewReader(os.Stdin)
	str, err := date.ReadString('\n')
	if err != nil {
		fmt.Println("终端数据获取失败")
	}
	up, err := conn.Write([]byte(str))
	if err != nil {
		fmt.Println("传入服务器失败")
	}
	fmt.Printf("向服务器发送了%d字节数据", up)
}
