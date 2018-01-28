---
date: 2016-07-19 15:07:00
title: 关于java部分知识点的总结（一）
categories:
    - java
tags:
    - java
---

	之前一直使用云笔记写自己遇到的一些体会，最近因为工作需要，准备回头梳理下以前自己学习的知识点，想把以前的记录下载博客里，既方便自己以后查阅，同时也能让自己有一个更深刻的记忆。
	Cannot use this in a static context 在一个static代码块或者是static方法中，不能使用this和supper，因为static在jvm加载时就会初始化，而此时this和super可能并不存在。构造器也是隐式的static方法（关于以前记录的言论，但我现在google并没有找不出相应的资料）
	延伸：关于static关键字：
	static修饰成员变量：静态变量，直接通过类名访问
	static修饰成员方法：静态方法，直接通过类名访问
	static修饰代码块：静态代码块，当JVM加载类时，就会执行该代码
	静态方法：①可直接通过类名访问②静态方法中不能使用this和super关键字③不能直接访问所属类的实例变量和实例方法④可直接访问类的静态变量和静态方法⑥静态方法必须被实现，简单的说就是不能和abstract抽象关键字并存，会报语法错误，如下：


```
/**
 * Created by root on 16-7-19.
 */
public abstract class Constructor {
    public Constructor(){
        System.out.println("static");
    }

    public static void main(String[] args) {

    }

    public abstract static  void test();
}

```


成员变量（全局变量）和局部变量的区别：①作用域不同②初始值不同<1>JAVA会给成员变量赋初始值<2>Java不会给局部变量赋初始值

关于java的访问修饰符大小：
       访问权限   类   包  子类  其他包

    public     ∨   ∨   ∨     ∨

    protect    ∨   ∨   ∨     ×

    default    ∨   ∨   ×     ×

    private    ∨   ×   ×     ×

查看端口占用情况以及杀死 相关端口进程
netstat -ano或者 netstat -ano | findstr 端口
taskkill /pid 端口

websocket中 jsonObject.put(key,value),若不使用else，则isSelf将永远存在，因为第一个if已经改变了
如下：

```
for (ChatServer openSession : set) {
            sysLogger.info(openSession);
            if(openSession.session.equals(session)){
                flag = true;
                System.out.println("推送到消息发送者本身");
                // 添加本条消息是否为当前会话本身发的标志
                jsonObject.put("isSelf", true);
            }else{
                jsonObject.put("isSelf", false);
            }
            // 发送JSON格式的消息
            // openSession.getAsyncRemote().sendText(jsonObject.toString());
            openSession.session.getBasicRemote().sendText(jsonObject.toString());
            sysLogger.info("Opensession:"+openSession);
        }
```

request.getSchema()可以返回当前页面使用的协议，http 或是 https;
request.getServerName()可以返回当前页面所在的服务器的名字;
request.getServerPort()可以返回当前页面所在的服务器使用的端口,就是80;
request.getContextPath()可以返回当前页面所在的应用的名字;

一共有三种设置方式url-pattern /文件夹名/*  /  *.do
servlet匹配 1.精确匹配 2.路径匹配  3.扩展名匹配 4.若前两者都不能匹配，则调用默认的servlet处理
URL 配置格式 三种：
1、完全路径匹配  (以/开始 ) 例如：/hello /init
2、目录匹配 (以/开始) 例如：/*  /abc/*
/ 代表网站根目录

3、扩展名匹配 (不能以/开始) 例如：*.do *.action
典型错误 /*.do
在浏览器中 访问的优先级顺序为：
优先级：完全匹配>目录匹配 > 扩展名匹配
![url配置匹配示例](http://img.blog.csdn.net/20160719144517191)


toLowerCase()字符串转换为小写
toUpperCase()字符串转换为大写

ServletContext,是一个全局的储存信息的空间，一个web应用程序只有唯一的一个，服务器开始，其就存在，服务器关闭，其才释放。request，一个用户可有多个；session，一个用户一个；而servletContext，所有用户共用一个。所以，为了节省空间，提高效率，ServletContext中，要放必须的、重要的、所有用户需要共享的线程又是安全的一些信息

newInstance: 弱类型。低效率。只能调用无参构造。
new: 强类型。相对高效。能调用任何public构造。
newInstance()是实现IOC、反射、面对接口编程 和 依赖倒置 等技术方法的必然选择，new 只能实现具体类的实例化，不适合于接口编程。
里面就是通过这个类的默认构造函数构建了一个对象，如果没有默认构造函数就抛出InstantiationException, 如果没有访问默认构造函数的权限就抛出IllegalAccessException

1.java中所有的关键字都是小写的，
null，ture和false是java语言中的保留字；
而sizeof，在java中不存在sizeof运算符；
implements和instanceof是java预言中的关键字。
2.doget/dopost是在HttpServlet中实现的 不是GenericServlet,service 判断并决定是调用doget还是dopost.   service()是在Servlet接口中定义的
1、实现Servlet接口。
2、继承GenericServlet。
3、继承HttpServlet。

3.AWT比较老，是各操作系统中图形功能的交集，依靠本地方法实现功能。注意：AWT在不同的操作系统中可能显示相同的风格，而swing由于比较先进，自己集成的一套，所以在不同的操作系统中一定显示相同的风格！

对于外部类来说 内部类相当于它的一个属性 内部类中的private也相当于它本身的private属性 所以根据类内可见原则 内部类private是可以被外部类访问的 ,内部类主要作用，使用组合可实现多继承，但个人理解没有使用接口实现多继承好，可以方便的访问外部类的属性方法，即使是private，可以使用组合实现不错的封装作用。


