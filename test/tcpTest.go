package main

import (
	"net"
	"fmt"
)

func main()  {
	//net.dial 拨号 获取tcp连接
	conn, err := net.Dial("tcp", "127.0.0.1:65535")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("获取127.0.0.1：65535的tcp连接成功...")
	defer conn.Close()

	//需要放在read前面，输出到服务端，否则服务端阻塞
	conn.Write([]byte("echo data to server ,then to client!!!"))

	//读取到buffer
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	fmt.Println(string(buffer))

}
