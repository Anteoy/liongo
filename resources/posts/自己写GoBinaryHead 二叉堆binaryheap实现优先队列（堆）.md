---
date: 2017-03-03 18:13:00
title: 自己写GoBinaryHead 二叉堆binaryheap实现优先队列（堆）
categories:
    - golang，数据结构
tags:
    - golang，数据结构,arrayList
---

##前言：

java GoBinaryHead二叉堆binaryheap实现优先队列（堆）
1. 二叉堆是完全二叉树 因为完全二叉数的规律（root始终最小） 用数组实现此数据结构优于链表
2. ，注意在插入和删除时，需要在数组实现的完全二叉树结构代码中，对原有节点数据进行上滤和下滤，插入时，和子树的根节点比较， 只有比子树根节点大才能满足定义， 否则循环交换位置。堆内元素向下移动为 下滤，删除后空余的位置，从上至下找最小儿子节点填充
3. 在printHeap()方法中对数组的遍历使用了去null操作。
4. 代码中已给出比较详尽注释。

###正文：

1. GoBinaryHeap.java
	```
    	package com.anteoy.dataStructuresAndAlgorithm.javav2.my;
    	
    	import java.util.NoSuchElementException;
    	
    	/**
    	 * Created by zhoudazhuang on 17-3-3.
    	 * Description:使用二叉堆binaryheap实现优先队列（堆）
    	 * 二叉堆是完全二叉树 因为完全二叉数的规律（root始终最小） 用数组实现此数据结构优于链表
    	 */
    	public class GoBinaryHeap<AnyType extends Comparable<? super AnyType>> {
    	
    	    private static final int DEFAULT_CAPACITY = 10;// 默认容量
    	    private int currentSize; // 当前堆大小
    	    private AnyType[] array; // 数组
    	
    	    public GoBinaryHeap() {
    	        this(DEFAULT_CAPACITY);
    	    }
    	
    	    public GoBinaryHeap(int capacity) {
    	        currentSize = 0;
    	        array = (AnyType[]) new Comparable[capacity + 1];
    	    }
    	
    	    public GoBinaryHeap(AnyType[] items) {
    	        currentSize = items.length;
    	        array = (AnyType[]) new Comparable[(currentSize + 2) * 11 / 10];
    	        int i = 1;
    	        for (AnyType item : items) {
    	            array[i++] = item;
    	        }
    	        buildHeap();
    	    }
    	
    	    /**
    	     * 从任意排列的项目中建立堆，线性时间运行
    	     */
    	    private void buildHeap() {
    	        for (int i = currentSize / 2; i > 0; i--) {
    	            percolateDown(i);
    	        }
    	    }
    	
    	    /**
    	     * 堆内元素向下移动 下滤 可以用上滤来进行相反的理解
    	     *删除后空余的位置 从上至下 从上至下找最小儿子节点填充
    	     * @param hole 下移的开始下标
    	     */
    	    private void percolateDown(int hole) {
    	        int child;
    	        AnyType tmp = array[hole];
    	        // 类似上滤的循环交换位置
    	        for (; hole * 2 <= currentSize; hole = child) {
    	            child = hole * 2;
    	            if (child != currentSize
    	                    && array[child + 1].compareTo(array[child]) < 0) {
    	                child++;
    	            }
    	            if (array[child].compareTo(tmp) < 0) {
    	                array[hole] = array[child];
    	            } else {
    	                break;
    	            }
    	        }
    	        array[hole] = tmp;
    	    }
    	
    	    /**
    	     * 插入 （上滤）
    	     * 需要满足完全二叉树的堆序性质
    	     * 插入时 和子树的根节点比较 只有比子树根节点大才能满足定义 否则移植循环交换位置
    	     * @param x
    	     */
    	    public void insert(AnyType x) {
    	        if (isFull()) {
    	            enlargeArray(array.length * 2 + 1);
    	        }
    	        //currentSize自增1并赋值给hole
    	        int hole = ++currentSize;
    	        // 插入时 和子树的根节点比较 只有比子树根节点大才能满足定义 否则移植循环交换位置
    	        for (; hole > 1 && x.compareTo(array[hole / 2]) < 0; hole /= 2) {
    	            array[hole] = array[hole / 2];
    	        }
    	        array[hole] = x;
    	    }
    	
    	    /**
    	     * 堆是否满
    	     * @return 是否堆满
    	     */
    	    public boolean isFull() {
    	        return currentSize == array.length - 1;
    	    }
    	
    	    /**
    	     * 堆是否空
    	     * @return 是否堆空
    	     */
    	    public boolean isEmpty() {
    	        return currentSize == 0;
    	    }
    	
    	    /**
    	     * 清空堆
    	     */
    	    public void makeEmpay() {
    	        currentSize = 0;
    	        for (AnyType anyType : array) {
    	            anyType=null;
    	        }
    	    }
    	
    	    /**
    	     * 找到堆中最小元素
    	     * @return 最小元素
    	     */
    	    public AnyType findMin() {
    	        if (isEmpty())
    	            return null;
    	        return array[1];
    	    }
    	
    	    /**
    	     * 删除堆中最小元素
    	     * 根据完全二叉树（堆序性质） 最小的为root节点
    	     * 删除后空余的位置 从上至下 从上至下找最小儿子节点填充
    	     * @return 被删除元素
    	     */
    	    public AnyType deleteMin() {
    	        if (isEmpty()) {
    	            throw new NoSuchElementException();
    	        }
    	        AnyType minItem = findMin();
    	        array[1] = array[currentSize];
    	        array[currentSize--] = null;
    	        percolateDown(1);
    	        return minItem;
    	    }
    	
    	    /**
    	     * 扩大数组容量
    	     * @param newSize 新的容量
    	     */
    	    private void enlargeArray(int newSize) {
    	        AnyType[] old = array;
    	        array = (AnyType[]) new Comparable[newSize];
    	        for (int i = 0; i < old.length; i++) {
    	            array[i] = old[i];
    	        }
    	    }
    	
    	    /**
    	     * 输出数组中的元素
    	     */
    	    public void printHeap() {
    	        for (AnyType anyType : array) {
    	            //不打印数组中的null
    	            if(anyType == null){
    	                continue;
    	            }
    	            System.out.print(anyType + " ");
    	        }
    	    }
    	
    	}
	
	```

2. GoBinaryHeapTest.java
    ```
        package com.anteoy.dataStructuresAndAlgorithm.javav2.my;
        
        /**
         * Created by zhoudazhuang on 17-3-3.
         * Description:二叉堆binaryheap实现优先队列（堆）测试类
         */
        public class GoBinaryHeapTest {
            public static void main(String[] args) {
                GoBinaryHeap<Integer> heap = new GoBinaryHeap<>();
                for (int i = 0; i < 10; i++) {
                    heap.insert(i);
                }
                heap.deleteMin();
                heap.deleteMin();
                heap.deleteMin();
                heap.printHeap();
            }
        }
    
    ```
    输出结果：
    ```
        3 4 5 7 9 8 6 
        Process finished with exit code 0
    ```
### 后记：
1. 参考文献：数据结构与算法分析