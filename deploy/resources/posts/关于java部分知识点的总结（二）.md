---
date: 2016-07-19 15:26:00
title: 关于java部分知识点的总结（二）
categories:
    - java
tags:
    - java
---
    之前一直使用云笔记写自己遇到的一些体会，最近因为工作需要，准备回头梳理下以前自己学习的知识点，想把以前的记录下载博客里，既方便自己以后查阅，同时也能让自己有一个更深刻的记忆。

关于反射：
	super.getClass() 得到的依然是runtime当前类，若要得到真正的父类，需要用super.getClass().getSuperclass()
虽然这里写的是super，但其实用this也一样

Integer i01=59 的时候，会调用 Integer 的 valueOf 方法，
1
2
3
4
5
  publicstaticInteger valueOf(inti) {
     assertIntegerCache.high>= 127;
     if(i >= IntegerCache.low&& i <= IntegerCache.high)
     returnIntegerCache.cache[i+ (-IntegerCache.low)];
     returnnewInteger(i); }
这个方法就是返回一个 Integer 对象，只是在返回之前，看作了一个判断，判断当前 i 的值是否在 [-128,127] 区别，且 IntegerCache 中是否存在此对象，如果存在，则直接返回引用，否则，创建一个新的对象。
在这里的话，因为程序初次运行，没有 59 ，所以，直接创建了一个新的对象。

int i02=59 ，这是一个基本类型，存储在栈中。

Integer i03 =Integer.valueOf(59); 因为 IntegerCache 中已经存在此对象，所以，直接返回引用。

Integer i04 = new Integer(59) ；直接创建一个新的对象。

System. out .println(i01== i02); i01 是 Integer 对象， i02 是 int ，这里比较的不是地址，而是值。 Integer 会自动拆箱成 int ，然后进行值的比较。所以，为真。

System. out .println(i01== i03); 因为 i03 返回的是 i01 的引用，所以，为真。

System. out .println(i03==i04); 因为 i04 是重新创建的对象，所以 i03,i04 是指向不同的对象，因此比较结果为假。

System. out .println(i02== i04); 因为 i02 是基本类型，所以此时 i04 会自动拆箱，进行值比较，所以，结果为真。

GenericServlet类的实现接口中包括了ServletConfig接口，但是它自身的init(ServletConfig config)方法又需要外界给它传递一个实现ServletConfig的对象，就是说GenericServlet和ServletConfig的依赖关系既是继承关系，也是一种关联关系。

Webservice是跨平台，跨语言的远程调用技术;
它的通信机制实质就是xml数据交换;
它采用了soap协议（简单对象协议）进行通信

java,exe是java虚拟机
javadoc.exe用来制作java文档
jdb.exe是java的调试器
javaprof,exe是剖析工具

jsp 页面有isErrorPage属性且值为false，不可以使用 exception 对象

年轻代：对象被创建时（new）的对象通常被放在Young（除了一些占据内存比较大的对象）,经过一定的Minor GC（针对年轻代的内存回收）还活着的对象会被移动到年老代（一些具体的移动细节省略）。
年老代：就是上述年轻代移动过来的和一些比较大的对象。Minor GC(FullGC)是针对年老代的回收
永久代：存储的是final常量，static变量，常量池。

HashMap 把 Hashtable 的 contains 方法去掉了 ，改成 containsvalue 和 containsKey 。因为 contains 方法容易让人引起误解。

notify()就是对对象锁的唤醒操作。但有一点需要注意的是notify()调用后，并不是马上就释放对象锁的，而是在相应的synchronized(){}语句块执行结束，自动释放锁后，JVM会在wait()对象锁的线程中随机选取一线程，赋予其对象锁，唤醒线程，继续执行。这样就提供了在线程间同步、唤醒的操作。

1.寄存器：最快的存储区, 由编译器根据需求进行分配,我们在程序中无法控制.
2. 栈：存放基本类型的变量数据和对象的引用，但对象本身不存放在栈中，而是存放在堆（new 出来的对象）或者常量池中（字符串常量对象存放在常量池中。）
3. 堆：存放所有new出来的对象。
4. 静态域：存放静态成员（static定义的）
5. 常量池：存放字符串常量和基本类型常量（public static final）。
6. 非RAM存储：硬盘等永久存储空间

Thread.sleep() 和 Object.wait(),都可以抛出 InterruptedException。这个异常是不能忽略的,因为它是一个检查异常(checked exception)

面向对象的五大基本原则

单一职责原则（SRP）
开放封闭原则（OCP）
里氏替换原则（LSP）
依赖倒置原则（DIP）
接口隔离原则（ISP）

StringBuffer内部数组的长度和字符串长度是不相同的，默认长度为16，你第一个append操作‘大家好’的字符长度未超出16，所以直接添加，而第二个操作的字符长度超过了16则调用扩展方法，将大小扩充到原来的两倍，操作为：(16+1)*2=34

以下是java concurrent包下的4个类，选出差别最大的一个

A、Semaphore：类，控制某个资源可被同时访问的个数;
B、 Future：接口，表示异步计算的结果；
C、ReentrantLock：类，具有与使用synchronized方法和语句所访问的隐式监视器锁相同的一些基本行为和语义，但功能更强大；
D、 CountDownLatch： 类，可以用来在一个线程中等待多个线程完成任务的类。答案 C

concurrence

（线程）并发，同事发生
volatile（线程并发中）易失，非易失，容易挥发的
用volatile修饰的变量，线程在每次使用变量的时候，都会读取变量修改后的最的值。volatile很容易被误用，用来进行原子性操作。
序列化中使用瞬态关键字transient（短暂，瞬态）让程序忽略此关键字所修饰的这一部分

ThreadLocal是解决线程安全问题一个很好的思路，它通过为每个线程提供一个独立的变量副本解决了变量并发访问的冲突问题。在很多情况下，ThreadLocal比直接使用synchronized同步机制解决线程安全问题更简单，更方便，且结果程序拥有更高的并发性。
ThreadLocalMap是ThreadLocal类的一个静态内部类，它实现了键值对的设置和获取（对比Map对象来理解），每个线程中都有一个独立的ThreadLocalMap副本，它所存储的值，只能被当前线程读取和修改。
总结一句话就是一个是锁机制进行时间换空间，一个是存储拷贝进行空间换时间。

sleep和wait的区别有：
  1，这两个方法来自不同的类分别是Thread和Object
  2，最主要是sleep方法没有释放锁，而wait方法释放了锁，使得敏感词线程可以使用同步控制块或者方法。
  3，wait，notify和notifyAll只能在同步控制方法或者同步控制块里面使用，而sleep可以在
    任何地方使用
   synchronized(x){
      x.notify()
     //或者wait()
   }
   4,sleep必须捕获异常，而wait，notify和notifyAll不需要捕获异常


非线程共享的那三个区域的生命周期与所属线程相同，而线程共享的区域与JAVA程序运行的生命周期相同，所以这也是系统垃圾回收的场所只发生在线程共享的区域（实际上对大部分虚拟机来说知发生在Heap上）的原因。


Python是解释执行的，其他语言都需要先编译

Java标识符由数字，字母和下划线（_），美元符号（$）组成。在Java中是区分大小写的，而且还要求首位不能是数字。最重要的是，Java关键字 不能当作Java标识符。
其实 普通的类方法是可以和类名同名的，和构造方法唯一的区分就是，构造方法没有返回值。

suspend() 和 resume() 方法：两个方法配套使用，suspend()使得线程进入阻塞状态，并且不会自动恢复，必须其对应的 resume() 被调用，才能使得线程重新进入可执行状态




