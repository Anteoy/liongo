---
date: 2017-03-24 17:13:00
title: go协程goroutine与Java多线程比较
categories:
    - golang，java
tags:
    - golang,java,多线程,goroutine
---

### 引言：
#### 个人理解的线程，协程和单，多核线程
1. 单核CPU上运行的多线程程序, 同一时间只能一个线程在跑, 系统帮你切换线程而已（cpu时间切片）, 系统给每个线程分配时间片来执行, 每个时间片大概10ms左右, 看起来像是同时跑, 但实际上是每个线程跑一点点就换到其它线程继续跑,效率不会有提高的,切换线程反倒会增加开销（线程的上下文切换），宏观的可看着并行，单核里面只是并发，真正执行的一个cpu核心只在同一时刻执行一个线程（不是进程）。
2. 多线程的用处在于，做某个耗时的操作时，需要等待返回结果，这时用多线程可以提高程序并发程度。如果一个不需要任何等待并且顺序执行能够完成的任务，用多线程是十分浪费的。
3. 个人见解，对于Thread Runable以及ThreadPoolExcutor等建立的线程，线程池内部是使用单核心执行（伪并行的并发多线程），而jdk1.7中提供的fork/join并发框架，使用的是多核心任务切分执行，个人觉得和map/reduce有一定类似之处。
4. 协程是一种用户态的轻量级线程，协程的调度完全由用户控制。而线程的调度是操作系统内核控制的，通过用户自己控制，可减少上下文频繁切换的系统开销，提高效率。

### 环境：
 1. ubuntu 16.04 LTS 
### 测试过程
	对比程序，系统发生1千万次并发，并发为一个无操作的空函数，使用time指令对比性能
java完整版   

    ```
        import java.util.concurrent.ExecutorService;
        import java.util.concurrent.Executors;
        
        /**
         * Created by zhoudazhuang on 14-8-15.
         * jdk1.7以及一下
         * new Runnable() {
        @Override
        public void run() {
        
        }
        jdk1.8可使用lambda
        () -> {}
         * 1. javac Main.java
         * 2. java Main(java二进制字节码,这样就可以运行了 不过这里再准备打包成运行jar包)
         * 3. jar -cvf my.jar *.class
         * 4.  time java -server -jar my.jar my.jar中没有主清单属性 修改压缩jar中MANIFEST.MF添加Main-Class:要空一格Main(主类class名字，有包加包 或者第二种方法直接在参数中指定)
         * 5. 改好后使用 time java -server -jar my.jar 或者直接使用time java -cp ./my.jar Main //主要不要用-jar了 -cp 目录和zip/jar文件的类搜索路径 直接指定就可以了
         * 输出
         * # time java -server -jar my.jar
         * elapsed time: 0.041s
         *  java -server -jar my.jar  0.26s user 0.01s system 215% cpu 0.126 total
         * time ls; time java是linux命令 不用-server我感觉输出也一样
         */
        public class Main {
            private static final int TIMES = 100 * 1000 * 100;
            public static void main(String[] args) throws Exception {
                ExecutorService service = Executors.newFixedThreadPool(Runtime.getRuntime().availableProcessors());
                long t1 = System.currentTimeMillis();
                for (int i=0;i<TIMES;i++) {
                    service.submit(new Runnable() {
                        @Override
                        public void run() {
        
                        }
                    });
                }
                service.shutdown();
                long t2 = System.currentTimeMillis();
                System.out.printf("elapsed time: %.3fs\n", (t2-t1)/1000f);
            }
        }
    ```
    输出
    
    ```
        # time java -cp ./my.jar Main
        elapsed time: 14.589s
        java -cp ./my.jar Main  75.20s user 2.96s system 486% cpu 16.072 total
    ```
    golang完整版
    
    ```
        //对比程序，系统发生1亿次并发，并发为一个无操作的空函数，使用time指令对比性能
        package main
        
        import (
        	"runtime"
        	"fmt"
        	"time"
        )
        
        const (
        	TIMES = 100 * 1000 * 100
        )
        
        func main() {
        	runtime.GOMAXPROCS(runtime.NumCPU())
        	fmt.Println("CPUs:", runtime.NumCPU(), "Goroutines:", runtime.NumGoroutine())
        	t1 := time.Now()
        	for i:=0; i<TIMES; i++ {
        		go func() {}()
        	}
        
        	for runtime.NumGoroutine() > 4 {
        		//fmt.Println("current goroutines:", runtime.NumGoroutine())
        		//time.Sleep(time.Second)
        	}
        	t2 := time.Now()
        	fmt.Printf("elapsed time: %.3fs\n", t2.Sub(t1).Seconds())
        }
    ```
    执行结果：

    ```
        CPUs: 8 Goroutines: 1
        elapsed time: 3.582s
        ./compareJava  13.09s user 1.49s system 405% cpu 3.591 total
    
    ```
由于多线程的处理机制不同，java的处理时间,cpu负载等明显高与goroutine，所以在此种情况下，goroutine在并发多任务处理能力上有着与生俱来的优势。
### 后记
示例程序主要参考自：[https://my.oschina.net/u/209016/blog/301705](https://my.oschina.net/u/209016/blog/301705)




