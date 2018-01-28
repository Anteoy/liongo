---
date: 2017-03-01 16:39:00
title: 自己写ArraryList GoArrayList
categories:
    - golang，数据结构
tags:
    - golang，数据结构,arrayList
---

## 前言：
　　java ArrayList的简易实现，代码中注释比较详尽，通俗易懂。
### 正文：

1. GoArrayList.java
	```
    	package com.anteoy.dataStructuresAndAlgorithm.javav2.my;
    	
    	import java.util.Iterator;
    	import java.util.NoSuchElementException;
    	
    	/**
    	 * Created by zhoudazhuang on 17-3-1.
    	 * Description: 简易Arrarylist实现
    	 */
    	public class GoArrayList<T> implements Iterable<T>
    	{
    	    //定义默认容量
    	    private static final int DEFAULT_CAPACITY=10;
    	    //定义当前存储的T[] 数组
    	    private T[] currentItems;
    	    //当前容量 size
    	    private int currentSize;
    	    //构造函数
    	    public GoArrayList() {
    	        clear();
    	    }
    	
    	//ArraryList扩容函数
    	//    private void grow(int minCapacity) {
    	//        // overflow-conscious code
    	//        int oldCapacity = elementData.length;
    	    //自动扩容1.5倍 2进制右移一位为除以2 左移乘以
    	//        int newCapacity = oldCapacity + (oldCapacity >> 1);
    	//        if (newCapacity - minCapacity < 0)
    	//            newCapacity = minCapacity;
    	//        if (newCapacity - MAX_ARRAY_SIZE > 0)
    	//            newCapacity = hugeCapacity(minCapacity);
    	//        // minCapacity is usually close to size, so this is a win:
    	//        elementData = Arrays.copyOf(elementData, newCapacity);
    	//    }
    	
    	    /**
    	     * 保证容量，自动扩容
    	     * @param newCapacity 新的容器大小
    	     */
    	    public void ensureCapacity(int newCapacity)
    	    {
    	        //指定新容量小如当前实际容量大小
    	        if (newCapacity< currentSize)
    	            return ;
    	        T[] old= currentItems;
    	        currentItems =(T[])new Object[newCapacity];
    	        for (int i = 0; i < size(); i++){
    	            currentItems[i]=old[i];
    	        }
    	    }
    	
    	    /**
    	     * clear和默认构造函数调用
    	     */
    	    public void clear()
    	    {
    	        currentSize =0;
    	        //初始化容量大小
    	        ensureCapacity(DEFAULT_CAPACITY);
    	    }
    	
    	    /**
    	     * 返回当前容量大小
    	     * @return
    	     */
    	    public int size(){
    	        return currentSize;
    	    }
    	
    	    /**
    	     * 判断当前ArrayList是否为空
    	     * @return
    	     */
    	    public boolean isEmpty()
    	    {
    	        return size()==0;
    	    }
    	
    	    /**
    	     *
    	     * @param index 数组下标
    	     * @return
    	     */
    	    public T get(int index)
    	    {
    	        if (index<0 || index>size())
    	            throw new ArrayIndexOutOfBoundsException();
    	        return currentItems[index];
    	    }
    	
    	    /**
    	     * set 方法
    	     * @param index
    	     * @param newVal
    	     * @return currentItems[index]
    	     */
    	    public T set(int index, T newVal)
    	    {
    	        if (index<0 ||index >size())
    	            throw new ArrayIndexOutOfBoundsException();
    	        T old= currentItems[index];
    	        currentItems[index]=newVal  ;
    	        return old;
    	    }
    	
    	
    	    /**
    	     * ArrayList所说没有用的值并不是null，而是ArrayList每次增长会预申请多一点空间
    	     * trimToSize 的作用只是去掉预留元素位置
    	     */
    	    public void trimToSize()
    	    {
    	        ensureCapacity(size());
    	    }
    	
    	    /**
    	     * add 操作
    	     * @param index
    	     * @param x
    	     */
    	    public void add(int index, T x)
    	    {
    	        if (currentItems.length== currentSize)
    	            //扩容1.5倍
    	            ensureCapacity(currentSize *2+1);
    	        //add添加过后，其余元素依次后移
    	        for(int i = currentSize; i>index; i--)
    	            currentItems[i]= currentItems[i-1];
    	        currentItems[index]=x;
    	        currentSize++;
    	    }
    	
    	    /**
    	     * 默认从最后位置添加移动
    	     * @param x
    	     * @return
    	     */
    	    public boolean add(T x)
    	    {
    	        add(currentSize, x);
    	        return true;
    	    }
    	
    	    /**
    	     * 移除
    	     * @param index
    	     * @return removeItem
    	     */
    	    public T remove(int index)
    	    {
    	        if (index<0 || index> currentSize -1)
    	            throw new ArrayIndexOutOfBoundsException();
    	        T removeItem= currentItems[index];
    	
    	        //移除的拿右边的元素填充
    	        for(int i = index; i< currentSize -1; i++)
    	            currentItems[i]= currentItems[i+1];
    	        currentSize--;
    	        return removeItem;
    	    }
    	
    	    //
    	    /**
    	     * 覆写iterator方法，返回一个用于ArrayList的iterator对象
    	     */
    	    @Override
    	    public Iterator<T> iterator() {
    	        // TODO Auto-generated method stub
    	        return new ArrayListIterator();
    	    }
    	
    	    /**
    	     * 内部类，实现arraylist的迭代器
    	     */
    	    private class ArrayListIterator implements Iterator<T>
    	    {
    	        //遍历指针指向的当前位置
    	        private int current=0;
    	        //是否有下一个元素
    	        public boolean hasNext()
    	        {
    	            return current< currentSize;
    	        }
    	        //返回下一个元素
    	        public T next()
    	        {
    	            if (!hasNext())
    	                throw new NoSuchElementException();
    	            //current的值会+1
    	            return currentItems[current++];
    	        }
    	        public void remove()
    	        {
    	            //之前next获取的对象的指针指向位置current++ 所以remove之前--current 否则删除的元素不是当前元素
    	            GoArrayList.this.remove(--current);
    	
    	        }
    	    }
    	}
	```

2. GoArrayListTest.java
	```
	package com.anteoy.dataStructuresAndAlgorithm.javav2.my;
	
	/**
	 * Created by zhoudazhuang on 17-3-1.
	 * Description:
	 */
	public class GoArrayListTest {
	    public static void main(String[] args) {
	        GoArrayList<Integer> goArrayList = new GoArrayList<>();
	        goArrayList.add(1);
	        goArrayList.add(2);
	        goArrayList.add(3);
	        //把数4的位置放在数3的位置上
	        goArrayList.add(2,4);
	        for (Integer goArrayList1:goArrayList){
	            System.out.println(goArrayList1);
	        }
	    }
	}
	```
	输出结果：
    ```
	1
	2
	4
	3
	```
### 后记：
1. 参考文献：数据结构与算法分析

