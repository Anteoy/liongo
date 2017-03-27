package main

import (
	"net"
	"fmt"
)

func main(){
	// tcp 监听并接受端口
	l, err := net.Listen("tcp", "127.0.0.1:65535")
	if err != nil {
		fmt.Println(err)
		return
	}
	//最后关闭
	defer l.Close()
	fmt.Println("tcp服务端开始监听65535端口...")
	// 使用循环一直接受连接
	for {
		//Listener.Accept() 接受连接
		c, err := l.Accept()
		if err!= nil {
			return
		}
		//处理tcp请求
		handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	//一些代码逻辑...
	fmt.Println("tcp服务端开始处理请求...")
	//读取
	buffer := make([]byte, 1024)
	//如果客户端无数据则会阻塞
	c.Read(buffer)

	//输出buffer
	c.Write(buffer)
	fmt.Println("tcp服务端开始处理请求完毕...")
}
