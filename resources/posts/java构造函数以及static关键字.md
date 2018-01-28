---
date: 2017-01-03 11:41:00
title: java构造函数以及static关键字
categories:
    - java
tags:
---

###关于java构造器

1. 原本无显示编码构造器，则有一个默认的隐式（隐藏的无参构造器），但是，当显示指定了构造器，则这个默认隐式的构造器将不存在，比如此时无法new无参的构造器（除非显示地编写声明无参的构造函数）。如下：
	```
	 package com.anteoy.coreJava.constructor;
	/**
	 * Created by zhoudazhuang
	 * Date: 17-1-3
	 * Time: 上午10:46
	 * Description :

	 */

	public class TestObj {
	    public TestObj(){

	    }
	    public TestObj(String o,String oo ,String ooo){

	    }
	}
	```
	```
	package com.anteoy.coreJava.constructor;

	/**

	 * Created by zhoudazhuang

	 * Date: 17-1-3

	 * Time: 上午10:46

	 * Description : public TestObj(){}没有无参构造new it则会报错

	 */

	public class NewObj {

	    public static void main(String[] args) {

	        TestObj obj = new TestObj();

	    }

	}
	```
2. 子类构造器若调用父类
如果子类构造器没有显式地调用父类的构造器，则将自动调用父类的默认（没有参数）的构造器。如果父类没有不带参数的构造器，并且在子类的构造器中又没有显式地调用父类的构造器，则j编译器将报语法错误
```
 package com.anteoy.coreJava.constructor;
/**
 * Created by zhoudazhuang
 * Date: 17-1-3
 * Time: 上午10:54
 * Description :
 */
public class SonObj extends TestObj{
    public SonObj(){
        super();//调用父类构造器
    }
}
```
###关于java static关键字
1. static修饰的变量（类变量，与类在jvm属同一时期加载，早于对象加载，jvm加载时加载一起）存在于jvm静态域中。
2. static属于类级别，但个人认为可以抽象地看成此static属于这个类（比如调用时可以显式地加上类名前缀），只不过static修饰的变量或常量，方法等和类在使用时，是属于同一级别（等级的）。
3. 在不同类里面可以定义名称相同的static变量（static final也是可以的）。
	```
	 package com.anteoy.coreJava.others;
	/**
	 * Created by zhoudazhuang
	 * Date: 16-12-28
	 * Time: 下午4:38
	 * Description : Ia,Ib接口有同名变量a，b只有其中一个有
	 */
	public class OoTest implements Ia,Ib{
	    public void oo(){
	        OoTest ooTest = new OoTest();
	//        int a = this.a; //编译报错
	        int a = Ia.a;//编译通过
	        int c = this.b;//编译通过
	    }
	}
	```
4. 关于初始化块和静态初始化块，初始化块{}是构造器的补充，不接受参数，定义一些所有对象共有的属性，方法等，主要可提高可维护性，和初始化块的复用性。主要区别：

    - 初始化顺序 静态初始化块--初始化块--构造方法

    - 静态初始化块只初始化一次，不能初始化普通非static变量
