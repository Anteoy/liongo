---
date: 2016-11-30 11:45:00
title: frp配置本地服务端口到服务器80端口
categories:
    - 架构
tags:
    - 架构
    - golang
---

##搭建环境：

 1. ubuntu 16.04 LTS （本地服务计算机） ubuntu 14.04 LTS(阿里云服务器)

 2. apache tomcat 7

 3. java 7

 4. frp 0.8.1 linux

##搭建过程：
1. 资源准备
	- frp 0.8.1 linux 二进制包
	- tomcat
	- 任意版本jvm
2. 开始安装

    1. 分别在服务端ubuntu和客户端ubuntu解压安装包（jdk以及tomcat这里不再赘述）

        ```tar -zxvf frp_0.8.1_linux.tar.gz```

    2. 配置本地ubuntu 16.04 LTS 中frpc.ini为：
	```
	#frpc.ini
	[common]
	server_addr = #阿里云服务器ip地址
	server_port = 7000
	auth_token = 123

	[web]
	type = http
	local_port = 8889 #本地端口地址
	```
	3. 修改阿里云服务器ubuntu 14.04 LTS版本frps.ini配置为：

        ```
		# frps.ini
		[common]
		bind_port = 7000
		vhost_http_port = 80

		[web]#不同标签名对应client中相同的标签
		type = http
		custom_domains = #阿里云服务器地址(即服务访问地址，这里最好这么配置)
		auth_token = 123

		[web-home]
		type = http
		custom_domains = #同上一标签
		auth_token = 123
        ```
	4. 启动服务
	```
	./frps -c ./frps.ini
	```
    5. 启动tomcat项目以及客户端
       ``` ./frpc -c ./frpc.ini```
	6. 此时访问custom_domains标签设置的地址:80即可经过frp代理进入到本地服务调试。

