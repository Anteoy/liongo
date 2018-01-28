---
date: 2017-02-25 23:23:00
title: docker使用容器ubuntu安装mongodb
categories:
    - 架构,linux
tags:
    - 架构,linux
---

### 前言：
　　最近准备使用docker安装一个mongo，可以使用Docker Hub上的镜像，后来就琢磨着自己用dockerfile来构建，后来在使用dockerfile构建过程中，因为TC网络环境，部分资源始终下载不了，后在容器中使用apt-get ppa依旧如此，最后决定使用mongo官网压缩包到容器里面安装，一切顺利。
### 安装过程：
1. 从docker Hub上拉去ubuntu image
	```
	docker pull ubuntu:16.04
	```

2. 交互式（-i），进入/bin/bash(-t),目录挂在到容器（-v）,宿主机和容器端口映射（-p）创建容器
    ```
    sudo docker run -i -t -v /home/zhoudazhuang/usr/local/:/home/zhoudazhuang -p 10000:27017 ubuntu  /bin/bash
    ```
3. 解压mongodb官网压缩包
    ```
    tar -zxvf mongodb-linux-x86_64-ubuntu1604-3.2.9.tar.gz
    ```

4. 进入mongo的bin目录，命令行下启动mongo：

	```
	nohup ./mongod &
	```

5. 此时在宿主机使用localhost:10000即可连接mongodb
### 后记：
关于容器开启则执行命令
1. 在容器中添加开机启动脚本，编写脚本mongo.sh，写入
	```
	#!/bin/bash
	/home/zhoudazhuang/usr/local/mongodb/mongodb-linux-x86_64-ubuntu1604-3.2.9/bin/mongod
	```
	或者
	```
	#!/bin/bash
	nohup /home/zhoudazhuang/usr/local/mongodb/mongodb-linux-x86_64-ubuntu1604-3.2.9/bin/mongod &
	```
2. 将你的启动脚本复制到 /etc/init.d目录下
	```
	cp mongod /etc/init.d/
	```
3. 执行如下命令将脚本放到启动脚本中去：
	```
	update-rc.d mongod defaults 95
	```