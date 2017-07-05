---
date: 2016-07-19 16:00:00
title: 关于java部分知识点的总结（三）
categories:
    - java
tags:
    - java
---

之前一直使用云笔记写自己遇到的一些体会，最近因为工作需要，准备回头梳理下以前自己学习的知识点，想把以前的记录下载博客里，既方便自己以后查阅，同时也能让自己有一个更深刻的记忆。

![dubbo](http://img.blog.csdn.net/20160719154221611)
dubbo构成：
2者（服务提供者【无状态】，服务消费者），2中心（服务注册中心，服务监控中心）。
(1) 连通性：连通流程
(2) 健状性：多数部分宕挂了，其余服务部分仍能正常使用
(3) 伸缩性：主要是伸，动态增加机器部署实例


```
public void ensureCapacity(int minCapacity) {
        int minExpand = (elementData != DEFAULTCAPACITY_EMPTY_ELEMENTDATA)
            // any size if not default element table
            ? 0
            // larger than default for default empty table. It's already
            // supposed to be at default size.
            : DEFAULT_CAPACITY;

        if (minCapacity > minExpand) {
            ensureExplicitCapacity(minCapacity);
        }
    }
```
在java的ArrayList类进行扩容的源代码中，并未使用hash算法，而hashmap在扩容过程中，是新建立一个hashmap后(如上)，使用hash算法重新计算hash，再放入新容器中。ArrayList和hashmap底层都是由数组来实现的，并且都是查询速度快，个人理解，估计ArrayList就是因为没有使用hash算法来进行存储，所以插入和删除比较慢，但因有其索引，所以随机访问的速度快

java.util.stream 与java.io 包里的 InputStream 和 OutputStream 是完全不同的概念。
不同于 StAX 对 XML 解析的 Stream，也不是 Amazon Kinesis 对大数据实时处理的 Stream。
Stream(java.util.stream) 是对集合（Collection）对象功能的增强，它专注于对集合对象进行各种非常便利、高效的聚合操作（aggregate operation），或者大批量数据操作 (bulk data operation)。
java.util.stream 是一个函数式语言+多核时代综合影响的产物。

nginx是横在用户的浏览器和自家的服务器之间。
dubbo是横在自家的服务器和自家的服务器之间。
![dubbo总体构架](http://img.blog.csdn.net/20160719154708519)
引用地址：[这里写链接内容](http://itindex.net/detail/50986-soa-%E6%9C%8D%E5%8A%A1-%E7%BC%96%E7%A8%8B?utm_source=tuicool&utm_medium=referral)

在今天用subversion 更新项目的时候，一直refreshing，后来如何修改idea的配置文件都不能成功，最后喊同时帮忙看看，修改idea配置未果，最后在文件路径使用svn检出发现连接不了svn,最终原因为ip地址冲突，（公司路由器和vpn，不会再ip冲突的情况下，重新非配ip地址） 绑定固定ip地址和dns后，可正常使用


POJO = "Plain Old Java Object"，是MartinFowler等发明的一个术语，用来表示普通的Java对象，不是JavaBean, EntityBean 或者 SessionBean。POJO不但当任何特殊的角色，也不实现任何特殊的Java框架的接口如，EJB， JDBC等等。VO:就是Value Object。它是主要指用户提交的数据,如提交的表单数据,这些数据全部保存一个叫VO的类中。PO：Persistent Object，持久化对象。
引用地址：http://liaojuncai.iteye.com/blog/1297709
注：JavaBean是指一段特殊的Java类，有默然构造方法,只有get,set的方法的java类的对象.

今天因部分原因，需要在centos 7.1 服务器上配置一个postgresql,安装正常执行，到最后一步设置防火墙开放端口允许远程访问就遇到一个比较无语的问题了，以前在centos 6.5，以及centos 7上使用的防火墙命令不一样，鼓捣了一阵，最后使用rpm -qa|grep firewall ，rpm查询所有-qa安装的软件包 交给grep处理 查找有关firewall的  查看本机防火墙版本，发现并不是firewall，而是firewalld，最后在firewall指令加一个d，如关闭防火墙命令：systemctl stop firewalld 禁止开机自启：systemctl disable firewalld 操作成功 远程成功连接数据库

javax.websocket所依赖的jar到底是在java.lang里面还是在服务器（如tomcat）里，因为今天转移项目的时候报错找不到javax.websocket相关jar包，从新导入jdk未果，最后加入tomcat lib才好的，这里证明应该是在tomcat等服务器里的

不大理解filter（过滤器）和interceptor（拦截器）的区别，遂google之。博文中有介绍：

1、拦截器是基于java的反射机制的，而过滤器是基于函数回调 。
2、过滤器依赖与servlet容器，而拦截器不依赖与servlet容器 。
3、拦截器只能对action请求起作用，而过滤器则可以对几乎所有的请求 起作用 。
4、拦截器可以访问action上下文、值栈里的对象，而过滤器不能 。
5、在action的生命周期中，拦截器可以多次被调用，而过滤器只能在容 器初始化时被调用一次 。
引用地址：
http://www.oschina.net/question/565065_86561

数据库优化相关：
1.硬件
2.配置
2.表结构设计
3.SQL以及索引（重要）

