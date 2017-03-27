package main

import (
	"net"
	"log"
)

func main()  {
	// 接受udp请求
	pc, err := net.ListenPacket("udp", "host:port")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	//simple read
	buffer := make([]byte, 1024)
	pc.ReadFrom(buffer)

	//simple write
	pc.WriteTo([]byte("Hello from client"), 5874)
}
