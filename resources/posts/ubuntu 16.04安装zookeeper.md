---
date: 2016-11-24 17:01:00
title: ubuntu 16.04安装zookeeper
categories:
    - linux
tags:
---

##搭建环境

 1. ubuntu 16.04 LTS
 2. zookeeper-3.5.1-alpha
 3. dubbo 调用

##搭建过程

**使用官方源码包进行安装**

1. 资源准备

    - zookeeper-3.5.1-alpha.tar.gz 源码包

    - 官方稳定版下载地址[http://apache.fayea.com/zookeeper/](http://apache.fayea.com/zookeeper/)

2. 开始安装

    1. 解压压缩包到指定目录

        ```tar -zxvf zookeeper-3.5.1-alpha.tar.gz -C /home/zhoudazhuang/local/```

    2. 使用第一步解压的源码包路径，cd /home/zhoudazhuang/local/zookeeper-3.5.1-alpha/conf 拷贝一份zoo_sample.cfg ，并命名为zoo.cfg

    3. 编辑zoo.cfg

        ```vim zoo.cfg```

        主要修改这几个地方：

            - 增加dataDir和dataLogDir节点（目录自己创建并指定，作为数据存储目录和日志文件目录，很有用），增加配置如下：

            ```dataDir=/home/strong/usr/local/zookeeper-3.5.1-alpha/data                dataLogDir=/home/strong/usr/local/zookeeper-3.5.1-alpha/logs```

            - 指定server地址

                ```server.1=172.16.64.125:2888:3888```

                *注：server.id=hostname:port:port。第一个端口用于集合体中的 follower 以侦听 leader；第二个端口用于 Leader 选举。第一个hostname即为本服务器地址*

            完整配置如下：
		```
		# The number of milliseconds of each tick
		tickTime=2000
		# The number of ticks that the initial
		# synchronization phase can take
		initLimit=10
		# The number of ticks that can pass between
		# sending a request and getting an acknowledgement
		syncLimit=5
		# the directory where the snapshot is stored.
		# do not use /tmp for storage, /tmp here is just
		# example sakes.
		dataDir=/home/strong/usr/local/zookeeper-3.5.1-alpha/data
		dataLogDir=/home/strong/usr/local/zookeeper-3.5.1-alpha/logs
		# the port at which the clients will connect
		clientPort=2181

		server.1=172.16.64.125:2888:3888

		# the maximum number of client connections.
		# increase this if you need to handle more clients
		#maxClientCnxns=60
		#
		# Be sure to read the maintenance section of the
		# administrator guide before turning on autopurge.
		#
		# http://zookeeper.apache.org/doc/current/zookeeperAdmin.html#sc_maintenance
		#
		# The number of snapshots to retain in dataDir
		#autopurge.snapRetainCount=3
		# Purge task interval in hours
		# Set to "0" to disable auto purge feature
		#autopurge.purgeInterval=1

		```
    4. 在上一步设置的data目录中新建文件mydata,并填入上一步中配置的server.1中的数字，我这里是1,这里相当于是编了号的服务器，有多台可参考此配置。

    5. 在系统变量path中加入zookeeper （此处可省略）
    ```export ZOOKEEPER_HOME=/home/hadooptest/zookeeper-3.4.3
    PATH=$ZOOKEEPER_HOME/bin:$PATH```
    6. zookeeper服务的启动与客户端的连接（调用命令在zookeeper的bin文件夹）

        - zkServer.sh start

        - zkCli.sh -server 172.16.64.125:2181(此处我在172.16.64.111中客户端对服务端进行调用)

 ##参考
    1. zookeeper官方参考文档1[http://www.cloudera.com/content/www/zh-CN/documentation/enterprise/5-3-x/topics/cdh_ig_zookeeper_package_install.html](http://www.cloudera.com/content/www/zh-CN/documentation/enterprise/5-3-x/topics/cdh_ig_zookeeper_package_install.html)
    2. zookeeper官方参考文档2[http://zookeeper.majunwei.com/document/3.4.6/GettingStarted.html](http://zookeeper.majunwei.com/document/3.4.6/GettingStarted.html)
   3. [https://my.oschina.net/phoebus789/blog/730787](https://my.oschina.net/phoebus789/blog/730787)
