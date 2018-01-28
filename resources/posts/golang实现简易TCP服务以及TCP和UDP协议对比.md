---
date: 2017-03-25 23:57:00
title: golang实现简易TCP服务以及TCP和UDP协议对比
categories:
    - 协议
tags:
    - HTTP,TPP,UDP,GOLANG
---

## 引言
1. ECHO（Echo Protocol，回绕协议，应答协议，如linux中的echo命令），用于查错及测量应答时间（运行在TCP和UDP协议上）本文示例为echo协议，服务器只需把收到的客户端的请求数据发给这个客户端即可，其它什么功能都不做。
tcp/ip是一个协议簇（族），TCP（传输控制协议）和IP（网际协议）是此协议簇的核心。七层OSI模型中，tcp/udp在传输层，而ip在网络层。
2. OSI七层协议由上到下分别是：应用层（http,https），表示层（简单地说不同计算机通信会话进行表示转化，使系统能够识别，把数据转换为能与接收者的系统格式兼容并适合传输的格式。），会话层（设置和维护电脑之间通信连接，ssh），传输层（TCP/UDP），网络层（网际协议IP），数据链路层（表头和表尾被加至数据包形成帧。数据链表头（DLH）是包含了物理地址和错误侦测及改错的方法。数据链表尾（DLT）是一串指示数据包末端的字符串,如以太网），物理层（针脚、电压、线缆规范、集线器、中继器、网卡、主机适配器）;
3. 运行在tcp上的协议有http,https,ftp,pop3(邮局协议，收邮件)，SMTP（Simple Mail Transfer Protocol，简单邮件传输协议，发送邮件）。运行在UDP上的协议有DHCP（Dynamic Host Configuration Protocol，动态主机配置协议，动态配置IP地址）。
4. DNS（Domain Name Service，域名服务），用于完成地址查找，邮件转发等工作（运行在TCP和UDP协议上）。

## TCP和UDP的区别
1. TCP（Transmission Control Protocol）是传输控制协议，UDP（User Datagram Protocol）是用户数据报协议。
2. TCP是面向指定连接的，UDP是面向广播，非连接的,TCP确认数据传输完整性和数据抵达，而UDP不保证。
3. TCP是数据安全的，使用三次握手进行安全确认，使用四次挥手断开连接。第一次客户端申请和服务器握手，第二次服务器确认后申请和客户端握手，第三次客户端发送ACK（Acknowledgement）确认字符申请握手，建立连接;四次挥手可以由任何一端发起，被发起的一端将会握手两次确认已无数据发送，最后由发起第一次关闭握手的一端发起最后一次握手关闭连接。
4. TCP发送的数据量大小，以及发送效率低于UDP。
5. UDP中应用如ping命令，DHCP，DNS,在线视频媒体，电视广播和多人在线游戏;基于TCP的协议有Telnet，FTP以及SMTP，TCP和UDP与生俱来的特性决定了他们各自的应用领域

## TCP客户端和服务器的简易go实现
服务端完整编码：

    ```
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
    ```
    客户端完整编码：
    ```
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
    
    ```