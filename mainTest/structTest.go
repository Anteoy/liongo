//用 new 分配内存 内建函数 new 本质上说跟其他语言中的同名函数功能一样：new(T) 分配了零值填充的 T 类型的内存空间，并且返回其地址，一个 *T 类型的值。用 Go 的术语说，它返回了一个指针，指向新分配的类型 T 的零值。记住这点非常重要。 这意味着使用者可以用 new 创建一个数据结构的实例并且可以直接工作。
//务必记得 make 仅适用于 map，slice 和 channel，并且返回的不是指针。应当用 new获得特定的指针。
package main

import "fmt"

type Vertex struct {
	X, Y float64
}
//func main() {
//	rect1 := new(Vertex)
//	rect2 := &Vertex{1, 2}
//	//分别打印 值，类型，*rect1解除引用的值
//	fmt.Printf("%v  %T  %v \n",  rect1,  rect1,  *rect1)
//	fmt.Printf("%v  %T  %v \n",  rect2,  rect2,  *rect2)
//
//	rect3 := Vertex{X: 5, Y: 6}
//	fmt.Printf("%v  %T\n",  rect3,  rect3)
//
//}
// 输出：
/*
&{0 0}  *main.Vertex  {0 0}
&{1 2}  *main.Vertex  {1 2}
{5 6}  main.Vertex
*/
