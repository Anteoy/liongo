---
date: 2017-02-17 10:18:00
title: java递归简析
categories:
    - java
tags:
    - java
---

**引言：**
 1. 给定一个整数，依次打印其没每位上的数字

**想法：**
    1. 把数字转为String,再转为char，最后放入char[]，逐个打印
    2. 递归

**代码:**

```
package com.anteoy.dataStructuresAndAlgorithm.javav2;

/**
 * Created by zhoudazhuang on 17-2-16.
 * Description:
 */
public class PrintString {
    public static void main(String [] args)
    {
        printStrs(123456789);
        printByte(123456);
        printString(123456);
    }
    /**
     * 逐个字符打印所给整数
     * @param n 这里系统标准输出流每次都只打印一个字符
     */
    private static void printStrs(int n) {
        //递归临界条件
        if(n>=10){
            //前递归顺序执行
            printStrs(n/10);
        }
        //后递归执行（反前递归顺序）
        System.out.println(n%10);
    }

    private static void printByte(int n){
        //转为String
        String s = String.valueOf(n);
        byte[] bytes = s.getBytes();
        for (byte b: bytes) {
            System.out.println(String.valueOf(b));
        }
    }

    private static void printString(int n){
        //转为String
        String s = String.valueOf(n);
        //转为char[]
        char[] chars = new char[s.length()];
        s.getChars(0,s.length(),chars,0);
        //for each输出
        for (char b: chars) {
            System.out.println(String.valueOf(b));
        }
    }
}

```

**后记：**

&nbsp; &nbsp;&nbsp; &nbsp; 这里主要注意，递归的前递归，递归临界条件，后后递归。比如这里的System.out.println(n%10)语句是达到临界后的后递归，执行顺序和临界前的前递归顺序是相反的，就像爬山一样，一层一层爬上去，而下山的时候，走下山的梯子和上去的是反方向的。


