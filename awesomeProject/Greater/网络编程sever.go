package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务端监听")
	cons, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("连接失败")
		return
	}

	//	监听成功
	for {

		synacc, err := cons.Accept()
		if err != nil {
			fmt.Println("连接失败")
			return
		}
		fmt.Printf("客户端链接成功=%v，连接的客户信息为%v\n", synacc, synacc.RemoteAddr().String())
		go poccess(synacc)
	}

}
func poccess(con net.Conn) {
	defer con.Close()
	for {
		buf := make([]byte, 512)
		str, err := con.Read(buf)
		if err != nil {
			//fmt.Println("数据读取错误")
			return
		}
		fmt.Println("接收到的信息为", string(buf[0:str]))
	}
}
