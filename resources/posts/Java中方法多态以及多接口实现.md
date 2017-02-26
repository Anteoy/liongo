---
date: 2017-01-09 22:24:00
title: Java中方法多态以及多接口实现
categories:
    - java
tags:
---

###关于java中方法多态

1. 通过多子类重写父类或接口实现。
2. 通过同类同方法（方法名相同，参数数量或者类型不同）实现，这里主要注意Java中判断同一方法的标准是方法名和参数，与返回值无关，如下，可简要看着yy(int a)
	```
	package com.anteoy.coreJava.polymorphism;
    /**
     * Created by zhoudazhuang
     * Date: 17-1-9
     * Time: 下午9:07
     * Description :java识别是否为重复冲突函数，依靠的是函数名和参数，与返回值无关，如yy(int a)
     */
    public class Polymorphism {
        //注释代码编译不通过,即使参数都为空 不允许仅仅只有返回值不同的同名函数
        /*String yy(int a){

        }
        int yy(int a){
            return 1;
        }*/
        String yy(int a,String b){
            return null;
        }
        int yy(int b,int c){
            return 1;
        }
        int yy (int a,boolean b){
            return 1;
        }
    }
	```
###关于java中实现多接口有同名参数冲突
1. 不同类中可以有public的同名变量
2. 当实现的接口中有冲突的public static final的变量时，如果需要在实现类中引用，则需带上接口名，如下：
interface Ia
```
package com.anteoy.coreJava.others;
/**
 * Created by zhoudazhuang
 * Date: 16-12-28
 * Time: 下午4:13
 * Description :
 */
public interface Ia {
    public static final int a = 2;
    int b = 3;

}
```
interface Ib
```
package com.anteoy.coreJava.others;
/**
 * Created by zhoudazhuang
 * Date: 16-12-28
 * Time: 下午4:13
 * Description :
 */
public interface Ib {
  public static int a = 1;
//    Ia.a;

}
```
interface OoTest
```
package com.anteoy.coreJava.others;

/**
 * Created by zhoudazhuang
 * Date: 16-12-28
 * Time: 下午4:38
 * Description : Ia,Ib接口有同名变量a，b只有其中一个有
 */
public class OoTest implements Ia,Ib{

    public int c = 0;

    {
        String sex = "ada";
    }

    public void oo(){
        OoTest ooTest = new OoTest();
//        int a = this.a; //编译报错
        int a = Ia.a;//编译通过
        int c = this.b;//编译通过
    }
}
```
interface OoTest2
```
package com.anteoy.coreJava.others;

/**
 * Created by zhoudazhuang
 * Date: 17-1-9
 * Time: 下午10:17
 * Description :
 */
public class OoTest2 {
    public int c = 0;
}

```
