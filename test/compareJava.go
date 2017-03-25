//go协程goroutine与Java多线程比较
//个人理解的线程，协程和单，多核线程
//1. 单核CPU上运行的多线程程序, 同一时间只能一个线程在跑, 系统帮你切换线程而已（cpu时间切片）, 系统给每个线程分配时间片来执行, 每个时间片大概10ms左右, 看起来像是同时跑, 但实际上是每个线程跑一点点就换到其它线程继续跑
//效率不会有提高的
//切换线程反倒会增加开销（线程的上下文切换），宏观的可看着并行，单核里面只是并发，真正执行的一个cpu核心只在同一时刻执行一个线程（不是进程）
//2. 多线程的用处在于，做某个耗时的操作时，需要等待返回结果，这时用多线程可以提高程序并发程度。如果一个不需要任何等待并且顺序执行能够完成的任务，用多线程简直是浪费
//3. 个人见解，对于Thread Runable以及ThreadPoolExcutor等建立的线程，线程池内部是使用单核心执行（伪并行的并发多线程），而jdk1.7中提供的fork/join并发框架，使用的是多核心任务切分执行，个人觉得和map/reduce有一定类似之处。
//4. 协程是一种用户态的轻量级线程，协程的调度完全由用户控制。而线程的调度是操作系统内核控制的，通过用户自己控制，可减少上下文频繁切换的系统开销，提高效率。
//对比程序，系统发生1亿次并发，并发为一个无操作的空函数，使用time指令对比性能
//由于多线程的处理机制不同，java的处理时间,cpu负载等明显高与goroutine，所以在此种情况下，goroutine在并发多任务处理能力上有着与生俱来的优势。
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