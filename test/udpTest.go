package main

import "net"

func main(){
	//net.dial 拨号 获取udp连接
	conn, err := net.Dial("tcp", "host:port")
	if err != nil {
		return err
	}
	defer conn.Close()

	//读取到buffer
	buffer := make([]byte, 1024)
	conn.Read(buffer)

	//输出
	conn.Write([]byte("Hello from client"))
}
