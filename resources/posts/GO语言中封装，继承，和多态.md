---
date: 2017-01-10 00:04:00
title: GO语言中封装，继承，和多态
categories:
    - golang
tags:
    - golang
---


##封装
go中的封装和java的不太一样，在go里面是没有java中的class，不过可以把struct看成一个类，封装可以简单地看作对struct的封装，如下
```
type obj1 struct {
	valte1 string
}
type obj2 struct {
	valte2 string
}
```
##继承
把struct看做类，struct中可以包含其他的struct，继承内部struct的方法和变量，同时可以重写，代码如下
```
package main

import "fmt"

type oo struct {
	inner
	ss1 string
	ss2 int
	ss3 bool
}

type inner struct {
	ss4 string
}

func (i *inner) testMethod(){
	fmt.Println("testMethod is called!!!")
}

func main()  {
	oo1 := new(oo)
	fmt.Println("ss4无值："+oo1.ss4)
	oo1.ss4 = "abc"
	fmt.Println("ss4已赋值"+oo1.ss4)
	oo1.testMethod()//继承调用
	oo1.inner.testMethod()//继承调用 这里也可以重写
}
```
##多态
go中的多台比java的隐匿得多，严格上说没有多态，但可以利用接口进行，对于都实现了同一接口的两种对象，可以进行类似地向上转型，并且在此时可以对方法进行多态路由分发，完全代码如下
```
package main

import "fmt"

type interfacetest interface {
	//testMothod1() string
	//testMothod()//这种会报语法错误 在go里面是不允许的
	iMethod() //加上int则会报错 说明go的方法判断有返回值，而java没有
}

type obj1 struct {
	valte1 string
}

type obj2 struct {
	valte2 string
}

//从属不同对象的testMothod 返回值不同的接口实现
func ( obj11 *obj1)iMethod(){
	fmt.Println("testMothod go obj1")
}

//从属不同对象的testMothod 返回值不同的接口实现
func ( obj11 *obj2)iMethod() {
	fmt.Println("testMothod go obj2")
}


func gorun(ii interfacetest){
	fmt.Println(ii.iMethod)
}

func main(){
	var i interfacetest
	//interfacetest_ := new(interfacetest)//这种方式进行多台路由转发会报错 GO需先声明 如 var i interfacetest
	obj1_ := new(obj1)
	//赋obj1
	i = obj1_
	i.iMethod()//正确打印
	gorun(i)
	gorun(obj1_)
	//interfacetest_.testMethod() //这种在java中允许，在go中是不允许的
	//赋obj2
	obj2_ := new(obj2)
	i = obj2_
	i.iMethod()//正确打印
	gorun(i)
	gorun(obj2_)
	list := [2]interfacetest{obj1_,obj2_}

	slice := []interfacetest{}
	slice = append(slice, obj1_)
	slice = append(slice, obj2_)
	for index,value := range slice {
		fmt.Println(index)
		fmt.Println(value)
	}
	fmt.Println(len(slice))

	fmt.Println(len(list))
}
```