---
date: 2017-03-05 00:17:00
title: ubuntu 安装本地版storm并运行WordCount
categories:
    - ubuntu，storm
tags:
    - ubuntu，storm
---

## 前言:
1. 开始从apache下载的最新版1.0.3，本地安装没有问题，但是当执行example-wordcount的时候报错找不到主类，后来解压jar包发现，1.03里面没有wordcount，有一些其他的类，于是第二次去下载安装了一个早期版本0.9.6，成功执行。
2. 关于1.0.3和0.9.6的配置异同，可参考官方文档地址[http://storm.apache.org/releases/1.0.3/Setting-up-a-Storm-cluster.html](http://storm.apache.org/releases/1.0.3/Setting-up-a-Storm-cluster.html) 和[http://storm.apache.org/releases/0.9.6/Setting-up-a-Storm-cluster.html](http://storm.apache.org/releases/0.9.6/Setting-up-a-Storm-cluster.html) 
3. 个人感觉在安装使用过程中，应尽量从官方文档以及FAQ等获取有用信息，否则自己容易进入一些误区。
4. 这是一个简易的本地版，供浏览学习之用，如有不妥之处，欢迎指正。

## 正文：
1. 安装一个或多个可用的zookeeper(Set up a Zookeeper cluster )
	可参考之前博客[ubuntu 16.04安装zookeeper](http://blog.csdn.net/yan_chou/article/details/53322429)
2. 本地版或集群版都需要如下依赖( Install dependencies on Nimbus and worker machines )
	1.0.3版本
	Java 7
	Python 2.6.6
	0.9.6版本
	Java 6
	Python 2.6.6
	注：Python在ubuntu以及其他Linux中默认已安装,jdk的安装这里不再赘述
3. 从apache官网下载storm并解压到任意位置（Download and extract a Storm release to Nimbus and worker machines），并加入path环境变量
	官方下载地址[http://storm.apache.org/downloads.html](http://storm.apache.org/downloads.html)
4. 打开解压后的storm，编辑conf下的storm.yaml
	
    1.0.3版本
	```
	storm.zookeeper.servers:
	     - "127.0.0.1"
	nimbus.seeds: ["127.0.0.1"]
	```
	0.9.6版本
	```
	storm.zookeeper.servers:
	     - "127.0.0.1"
	nimbus.host: "127.0.0.1"
	```
	配置zookeeper和主控节点host
5. 运行主控节点nimbus和工作节点Supervisor
	
    执行bin目录下可storm,Nimbus: Run the command "bin/storm nimbus" under supervision on the master machine.
	Supervisor: Run the command "bin/storm supervisor" under supervision on each worker machine. The supervisor daemon is responsible for starting and stopping worker processes on that machine.
	```
	bin/storm nimbus
	```
	正确响应如下
	
	```
	...ib/ring-jetty-adapter-0.3.11.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/minlog-1.2.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/jline-2.11.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/core.incubator-0.1.0.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/snakeyaml-1.11.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/log4j-over-slf4j-1.6.6.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/commons-io-2.4.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/asm-4.0.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/jetty-util-6.1.26.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/ring-servlet-0.3.11.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/objenesis-1.2.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/ring-core-1.1.5.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/conf -Xmx1024m -Dlogfile.name=nimbus.log -Dlogback.configurationFile=/usr/dev_sda_7/local/apache-storm-0.9.6/logback/cluster.xml backtype.storm.daemon.nimbus
	```
	```
	bin/storm supervisor
	```
	正确响应如下
	
	```
    	...sda_7/local/apache-storm-0.9.6/lib/core.incubator-0.1.0.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/snakeyaml-1.11.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/log4j-over-slf4j-1.6.6.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/commons-io-2.4.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/asm-4.0.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/jetty-util-6.1.26.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/ring-servlet-0.3.11.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/objenesis-1.2.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/lib/ring-core-1.1.5.jar:/usr/dev_sda_7/local/apache-storm-0.9.6/conf -Xmx256m -Dlogfile.name=supervisor.log -Dlogback.configurationFile=/usr/dev_sda_7/local/apache-storm-0.9.6/logback/cluster.xml backtype.storm.daemon.supervisor
	```
6. 执行0.9.6下的wordcount示例

	```
	$ /usr/dev_sda_7/local/apache-storm-0.9.6/bin/storm jar /usr/dev_sda_7/local/apache-storm-0.9.6/examples/storm-starter/storm-starter-topologies-0.9.6.jar storm.starter.WordCountTopology | grep '[*,*]'
	```
	| grep '[*,*]'  用来过滤日志
	在正确响应如下：
	
	```
	13390 [Thread-11-count] INFO  backtype.storm.daemon.executor - Processing received message source: split:6, stream: default, id: {}, ["four"]
	13390 [Thread-11-count] INFO  backtype.storm.daemon.task - Emitting: count default [four, 57]
	13390 [Thread-13-count] INFO  backtype.storm.daemon.executor - Processing received message source: split:6, stream: default, id: {}, ["and"]
	13390 [Thread-13-count] INFO  backtype.storm.daemon.task - Emitting: count default [and, 100]
	```
	成功执行计数
## 后记：
参考文献 apache storm官方文档:

[http://storm.apache.org/releases/0.9.6/Setting-up-a-Storm-cluster.html](http://storm.apache.org/releases/0.9.6/Setting-up-a-Storm-cluster.html)/releases/0.9.6/Setting-up-a-Storm-cluster.html))