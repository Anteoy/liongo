//test main TODO this
package main

import (
	"fmt"
	"runtime"
)

func test(c chan bool, n int) {

	x := 0
	for i := 0; i < 1000000000; i++ {
		x += i
	}

	println(n, x)
	println(1, 2)

	if n == 9 {
		c <- true
	}
}

func main() {
	runtime.GOMAXPROCS(1) //设置cpu的核的数量，从而实现高并发
	c := make(chan bool)
	go test(c, 1)
	go test(c, 2)
	for i := 0; i < 10; i++ {
		go test(c, i)
	}

	fmt.Println(<-c)
	fmt.Println("main ok")

}
