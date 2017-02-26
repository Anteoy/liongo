---
date: 2016-07-29 16:05:00
title: 关于近期java研发工程师面试的一些总结
categories:
    - java
tags:
    - java
---

 今天周五，从上个公司离职到现在，忙了3,4天，之前拿到offer的那家公司在往外包公司发展，办公环境也实在不行，于是便有了这几天忙活的事情，这几天投了不少，原本面试已经排到下周二，不过现在算是告一段落了，也拿到了一个初创公司的offer。下面是一些自我总结，方便自己以后查阅，不对也欢迎大家指正和补充。
1.关于oracle的列转行 以及oracle的存储过程
（当时并未回答上）google之主要可使用union all(列转行) ；case when then ，decode（行转列）等
参考http://blog.163.com/magicc_love/blog/static/185853662201371481247696/

2.token失效和定时器的实现。
	线程 setInterval()等
3.I/O InputStream和Reader的区别  OutputStream和Writer的区别
   1.字节流和字符流
   2.InputStreamReader 是字符流Reader的子类，是字节流通向字符流的桥梁。你可以在构造器重指定编码的方式，如果不指定的话将采用底层操作系统的默认编码方式，例如 GBK 等
   参考http://blog.csdn.net/z69183787/article/details/8179889
4.大量if else if else if else 如何进行处理和优化？
 使用java多态 策略模式来实现
5.ajax工作原理，不能使用jquery
	xmlHttpRequest
6.如何实现一个日志系统
spring aop log4j lsf4j



