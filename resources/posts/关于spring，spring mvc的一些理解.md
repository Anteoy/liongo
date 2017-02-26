---
date: 2016-07-19 23:16:00
title: 关于spring，spring mvc的一些理解
categories:
    - java
tags:
    - spring
---

最近一直在思考，spring，springmvc 到底谁是超集，谁是子集，而后google了一大堆资料，但并没有看到一针见血的，这里说下个人对此的总结理解。
首先，spring的诞生是由于EJB太过于笨重，而开发出的一个轻量级的应用Bean框架，而springmvc是在基于spring的基础上，扩展开发出来（利用spring的IOC，AOP,其他spring预留用户扩展的接口）的一个，类似于strusts2的controller的mvc web应用框架，在spring和spring mvc整合中，spring mvc担任控制层（类似ssh钟struts2的功能职责），spring担任其他Bean的管理作为父容器，springmvc作为子容器，关于spring和springmvc整合部分问题可参考链接：
http://www.imooc.com/article/6331
对于spring，springmvc 我首先想到的是：
spring 提供了AOP和IOC 然后也留下了很多接口提供给用户自己扩展 而spring mvc就是使用了spring的核心的情况下 扩展了mvc的支持 就像struts2的 action controller控制层一样、
而后请教朋友的见解，如下图：
![这里写图片描述](http://img.blog.csdn.net/20160719230745935)
他理解为：
spring mvc也是spring的一个模块，是spring对于mvc实现的一个功能模块，spring由各个模块组成，各司其职，spring-core.jar主要提供ioc实现，spring-aop。jar提供aop实现。一个模块提供一方面功能的实现
在此感谢fudali133的部分纠正，在此之前我进入了一个误区，误以为spring只有核心的那几个模块jar。
总结：
所有的(包括spring mvc模块)的前缀都是spring,可以理解为springmvc为spring的一部分，一个子模块，spring是集合了所有相关功能应用的轻量级应用框架的一个集合。