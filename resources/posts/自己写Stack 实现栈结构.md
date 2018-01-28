---
date: 2017-03-02 14:42:00
title: 自己写Stack 实现栈结构
categories:
    - golang，数据结构
tags:
    - golang，数据结构,arrayList
---

## 前言：
### 栈的应用：
	1. 可计算数学后缀表达式
	2. 把正常中缀表达式转换为后缀表达式
	3. 计算检测编译程序{}等括号符号是否正确，是否存在语法错误
	4. 递归中需要实用栈存储方法信息，计算机中函数调用是通过栈(stack)这种数据结构实现，在递归中调用一层函数，栈就会加一层栈帧，每当函数返回，栈就会减少一层栈帧。
### 正文：

1. java中使用数组实现栈
	```
    	package com.anteoy.dataStructuresAndAlgorithm.javav2;
    	    
    	    import java.util.ArrayList;
    	    
    	    /**
    	     * Created by zhoudazhuang on 17-3-1.
    	     * Description:
    	     */
    	    public class ListGo {
    	    
    	        public static void main(String[] args) {
    	            ArrayList arrayList = new ArrayList();
    	            arrayList.add(1);
    	            arrayList.add(2);
    	            arrayList.add(3);
    	    
    	            //进栈
    	            arrayList.add(arrayList.size(),2);
    	    
    	            //出栈
    	            arrayList.remove(arrayList.size()-1);
    	    
    	            System.out.println(arrayList);
    	        }
    	    }
	```

2. java中LinkedList实现栈
    ```
        package com.anteoy.dataStructuresAndAlgorithm.javav2.my;
            
            import java.util.LinkedList;
            
            /**
            * Created by zhoudazhuang on 17-3-2.
            * Description:
            */
            public class StackByLinkedList {
            
            public static void main(String[] args) {
                useLinkedListAsLIFO();
            }
            /**
             * 将LinkedList当作 LIFO(后进先出)的堆栈
             */
            private static void useLinkedListAsLIFO() {
                // 新建一个LinkedList
                LinkedList stack = new LinkedList();
            
                // 将1,2,3,4添加到堆栈中
                stack.push("1");
                stack.push("2");
                stack.push("3");
                stack.push("4");
                // 打印“栈”
                System.out.println(stack);
            
                // 删除“栈顶元素”
                System.out.println("stack.pop():"+stack.pop());
            
                // 取出“栈顶元素”
                System.out.println("stack.peek():"+stack.peek());
            
                // 打印“栈”
                System.out.println("stack:"+stack);
            }
            }
    ```
### 后记：
1.  尾递归和递归 局部变量栈在递归引用中，不能算尾递归
2. 参考文献：数据结构与算法分析�