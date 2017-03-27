package main
import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)
var host = flag.String("host", "", "host")
var port = flag.String("port", "3333", "port")
func main() {
	flag.Parse()
	var l net.Listener
	var err error
	l, err = net.Listen("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + *host + ":" + *port)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}
		//logs an incoming message
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}
func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		// Copy 从 src 中复制数据到 dst 中，直到所有数据都复制完毕，返回复制的字节数和
		// 遇到的错误。如果复制过程成功结束，则 err 返回 nil，而不是 EOF，因为 Copy 的
		// 定义为“直到所有数据都复制完毕”，所以不会将 EOF 视为错误返回。
		// 如果 src 实现了 WriteTo 方法，则调用 src.WriteTo(dst) 复制数据，否则
		// 如果 dst 实现了 ReadeFrom 方法，则调用 dst.ReadeFrom(src) 复制数据
		io.Copy(conn, conn)
	}
}