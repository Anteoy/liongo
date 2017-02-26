---
date: 2016-07-18 20:55:00
title: 关于java基础点的一些随笔
categories:
    - java
tags:
---

最近回头来啃一些java的基础知识，在此记录下个人的收获和理解。
1.关于位运算符的异或，简单的说就是在或||运算符的同时为1,1的时候，其结果为0，其他和或运算符的计算结果完全一致，这里写了下利用异或来达到，在不利用第三个变量的基础上，交换两个变量的值。

```
/**
 * Created by yan.chou on 16-7-18.
 * 不使用第三个变量实现两变量互换
 */
public class Java3 {
    public static void main(String[] args) {
        int a = 0;
        int b = 10;
        a = b - a;
        b = b - a;//利用差值计算
        a = b + a;
//        System.err.printf("%d,%d",a,b);
        yh();
    }

    //通过异或运算能够使数据中的某些位翻转，其他位不变。这就意味着任意一个数与任意一个给定的值连续异或两次，值不变。
    public static void yh(){
        int a = 2;
        int b = 10;
        a = a^b;
        b = a^b;//a^b 去b得a
        a = a^b;//a^b 去a得b 这里的b已经是原来的a
        System.err.println(a+" "+b);
    }
}
```
2.对于
Integer a = new Integer(1);
Integer a = Integer.valueOf();
两种方式返回的都是Integer包装类的实例， valueOf()的效率更高，因为它使用了缓冲机制，new Integer()返回的永远是不同的实例，而当数值在-128到127之间的时候 ，使用Integer.valueOf()返回的就是同一对象。demo比较简单，我这里就不上代码了，有兴趣的朋友可以自己尝试。
3.关于string类的几个方法
注：s代表的是字符串的实例
s.toCharArray();//String转换为char数组
s.split("_")；//返回以\_分割的字符串数组String[]
s.subString(0,1);//看api为 inclusion(包含)0位置处的char，而exclusion(排除)1位置处的char，如s = "abc"，则s.subString(0,1)返回a
4.关于java.util.Arrays 和java.util.Collections
4.1对于数组类
Arrays.sort(各种类型的数组);//数组正序排序
Arrays.sort(部分类型的数组，Collections.reverseOrder())//数组倒序，但是不支持基础类型，使用基础类型数组会编译错误（关于第二个参数，实现comparator借口的实例我没有做进一步探索，若有朋友知悉欢迎补充）
4.2对于集合类
Collections.sor(集合);//集合正序排序
Collections.reverse(集合);//集合反序排序
如下面测试代码

```
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

/**
 * Created by root on 16-7-18.
 * 实现倒序排序
 */
public class Java4 {
    public static void main(String[] args) {
        List<Long> S = new ArrayList<Long>();
        S.add(1l);
        S.add(19l);
        S.add(-19l);
        S.add(0l);
        Long[] t = (Long[]) S.toArray(new Long[S.size()]);
        Arrays.sort(t);//对数组进行排序
        Collections.sort(S);//正序
        Collections.reverse(S);//倒序输出
        for (Long o:S
             ) {
            System.out.println(o);
        }
    }
}
```
5.关于类型转换

Arrays.asList(数组);//数组转化为list:
Arrays.toString(字符数组);//数组转化为字符串：
String.valueOf(a);//a为字符串(对象)数组转化为string
此函数api
Parameters:
obj - an Object.
Returns:
if the argument is null, then a string equal to "null"; otherwise, the value of obj.toString() is returned.
注：String.parseString()是没有这个方法的 因为Integer.parseInt等转化的是基本类型Int，而String类包装类型，没有基本类型 Character同样没有parseCharacter()；方法（原因不详） 但有valusOf方法 如下：
![这里写图片描述](http://img.blog.csdn.net/20160718204657757)
Integer.parseInt(）；//String转化为int
api:
Parameters:
s - a String containing the int representation to be parsed
Returns:
the integer value represented by the argument in decimal.
Integer.parseInt();//String转化为包装类int这里不在赘述
Integer.valueOf()；//String转化为包装类Integer这里不在赘述
api:
Parameters:
s - the string to be parsed.
Returns:
an Integer object holding the value represented by the string argument.
其余Boolean,Float,Double等基本相同

数组和List
list.toArray(new Long[size]);list转化为相应类型数组
Arrays.asList(数组对象);//数组转化为数组对象，同样相关有	    Arrays.toString()方法

字符串和字符数组
s.toCharArray();//字符串转化为数组
new String(数组); //数组转化为字符串
Arrays.toString(数组);//数组转化为字符串
String.valueOf(数组)；//这里转化不一定能达到想要的结果，我这里测试时，显示的是对象的空间地址

下面是关于转化的代码：

```
import java.util.Arrays;
import java.util.Collections;

import static java.util.Arrays.asList;

/**
 * Created by root on 16-7-18.
 */
public class Java5  {
    public static void main(String[] args) {
        String s = "AW_09*19\\12_77";
        String[] a = s.split("_");
        Arrays.sort(a);//升序
        Arrays.asList(a);
        Collections.reverse(asList(a));//降序
        char[] c = s.toCharArray();
     //   Arrays.sort(c,Collections.reverseOrder()); 此种用法不能对基本类型数组记性倒序
        for (char o:c
             ) {
            System.out.print(o);
        }

        String.valueOf(a);
        System.out.println(Arrays.toString(a));
        System.out.println(new String(c));//可以用此转换
        System.out.println(String.valueOf(a));//转换出来为地址 不能用次转换
        System.out.println(s.substring(0,1));
    }


}

```


