---
date: 2017-03-01 22:24:00
title: 自己写LinkedList GoLinkedList
categories:
    - golang，数据结构
tags:
    - golang，数据结构,LinkedList
---

##前言：
　　java GoLinkedList的简易实现，代码中注释比较详尽，通俗易懂,注意事项亦在注解中标明。
###正文：

1. GoLinkedList.java
	```
    	package com.anteoy.dataStructuresAndAlgorithm.javav2.my;
    	
    	import java.util.Iterator;
    	
    	/**
    	 * Created by zhoudazhuang on 17-3-1.
    	 * Description: 简易Linkedlist实现
    	 */
    	public class GoLinkedList<T> implements Iterable<T>{
    	  
    	    //当前容量 size
    	    private int currentSize;
    	
    	    //附加的数据域 用来帮助迭代气检测集合中的变化 代表自从构造依赖对链表所做改变的次数 当和迭代器内储存的modCount不一致则抛出异常
    	    private int modCount = 0;
    	
    	    //头节点 注意没有值和前趋节点 只有后继节点（指向下一个节点的链）
    	    private Node<T> beginMarker;
    	    //尾节点 和头节点相反
    	    private Node<T> endMarker;
    	
    	
    	    /**
    	     * 私有静态嵌套类
    	     * 在一个类中，数据成员通常都是私有的
    	     * prev,next都是Node的实例
    	     * @param <AnyType>
    	     */
    	    private static class Node<AnyType> {
    	        //信使携带数据
    	        public AnyType data;
    	        //到前一个节点的链 主要注意这里的链其实也是一个节点Node
    	        public Node<AnyType> prev;
    	        //到下一个节点的链 主要注意这里的链其实也是一个节点Node
    	        public Node<AnyType> next;
    	
    	        //构造函数
    	        public Node( AnyType d, Node<AnyType> p, Node<AnyType> n ) {
    	            data = d; prev = p; next = n;
    	        }
    	    }
    	
    	    public GoLinkedList()
    	    {
    	        clear();
    	    }
    	
    	    /**
    	     * 创建并连接头尾节点，然后设置大小为0
    	     */
    	    private void clear() {
    	        beginMarker = new Node<T>(null,null,null);
    	        endMarker = new Node<T>(null,beginMarker,null);
    	        beginMarker.next = endMarker;
    	
    	        currentSize = 0;
    	        modCount++;
    	    }
    	
    	    public int size(){
    	        return currentSize;
    	    }
    	
    	    /**
    	     * 判断是否为空
    	     * @return
    	     */
    	    public boolean isEmpty(){
    	        return size() == 0;
    	    }
    	
    	    public boolean add(T x){
    	        add(size(),x);
    	        return true;
    	    }
    	
    	    public boolean add(int idx, T x){
    	       return addBefore( getNode(idx), x );
    	    }
    	
    	    public T get(int idx){
    	        return getNode(idx).data;
    	    }
    	
    	    public T set (int idx, T newVal){
    	        Node<T> p = getNode(idx);
    	        T oldVal = p.data;
    	        p.data = newVal;
    	        return oldVal;
    	    }
    	
    	    public T remove(int idx){
    	        return remove(getNode(idx));
    	    }
    	
    	    /**
    	     * 通过获取一个新节点，然后按所只是的顺序改变指针
    	     * 完成一个双向链表中的插入操作
    	     * @param p
    	     * @param x
    	     */
    	    public boolean addBefore(Node<T> p, T x){
    	        //使用信使储存新node 前一个链为p.prev 后一个链为p 前趋后继 表示插入在p和p之前一个节点之间
    	        Node<T> newNode = new Node<>(x, p.prev, p);
    	        //newNode的前节点的后一个链赋为本身
    	        newNode.prev.next = newNode;
    	        //改变p的前节点为newNode
    	        p.prev = newNode;
    	        //容量+1
    	        currentSize++;
    	        //操作次数加1
    	        modCount++;
    	        return true;
    	    }
    	
    	    /**
    	     * 删除p节点
    	     * @param p
    	     * @return 删除的节点
    	     */
    	    private T remove(Node<T> p){
    	        //节点p的后继节点的前趋链被赋予当前p的前趋节点
    	        p.next.prev = p.prev;
    	        //类似上面理解
    	        p.prev.next = p.next;
    	        //大小减1
    	        currentSize--;
    	        //操作次数加1
    	        modCount++;
    	        return p.data;
    	    }
    	
    	    /**
    	     * 这里进行折半分流遍历
    	     * @param idx
    	     * @return
    	     */
    	    private Node<T> getNode(int idx)
    	    {
    	        Node<T> p;
    	
    	        //不符合规则 直接抛异常
    	        if(idx<0 || idx>size())
    	            throw new IndexOutOfBoundsException();
    	        //如果索引表示该表前半部分的一个节点
    	        if(idx<size()/2) {//地址空间不连续 不能想ArrayList那样获取 内部实现是双向链表 而不是数组 所以根据指向节点遍历获取值
    	            //比如这里获取的是索引idx的前一个节点的直接后继
    	            p = beginMarker.next;
    	            for(int i = 0; i<idx; i++)
    	                p = p.next;
    	        }
    	        else {//遍历方式和上面相反
    	            p = endMarker;
    	            for(int i = size(); i>idx; i--)
    	                p = p.prev;
    	        }
    	        return p;
    	    }
    	
    	    public Iterator<T> iterator(){
    	        return new LinkedListIterator();
    	    }
    	
    	    /**
    	     * 抽象了位置概念
    	     */
    	    private class LinkedListIterator implements java.util.Iterator<T> {
    	        //表示包含由next所返回的项的节点
    	        private Node<T> current = beginMarker.next;
    	        //主要为了保证操作是安全的
    	        private int expectedModCount = modCount;
    	        //保留一个当前位置 检测能否remove
    	        private boolean okToRemove = false;
    	
    	        //简单实现 不检查链表的修改
    	        public boolean hasNext() {
    	            return current != endMarker;
    	        }
    	
    	        public T next() {
    	            if (modCount != expectedModCount)
    	                throw new java.util.ConcurrentModificationException();
    	            if (!hasNext())
    	                throw new java.util.NoSuchElementException();
    	
    	            T nextItem = current.data;
    	            current = current.next;
    	            okToRemove = true;
    	            return nextItem;
    	        }
    	
    	        public void remove()
    	        {
    	            if(modCount!=expectedModCount)
    	                throw new java.util.ConcurrentModificationException();
    	            if(!okToRemove)
    	                throw new IllegalStateException();
    	            GoLinkedList.this.remove(current.prev);
    	            expectedModCount++;
    	            okToRemove = false;
    	        }
    	    }
    	
    	}
	
	```

2. GoLinkedListTest.java
	```
    	package com.anteoy.dataStructuresAndAlgorithm.javav2.my;
    	
    	/**
    	 * Created by zhoudazhuang on 17-3-1.
    	 * Description: GoLinkedList测试
    	 */
    	public class GoLinkedListTest {
    	
    	
    	    public static void main(String[] args) {
    	        GoLinkedList<Integer> goLinkedList = new GoLinkedList<>();
    	        goLinkedList.add(1);
    	        goLinkedList.add(2);
    	        goLinkedList.add(3);
    	        goLinkedList.add(2,4);
    	        for (Integer go:goLinkedList){
    	            System.out.println(go);
    	        }
    	        System.out.printf(String.valueOf(goLinkedList.get(1)));
    	    }
    	}
	```
	输出结果：
	```
		1
		2
		4
		3
		2
		Process finished with exit code 0
	```
###后记：
1. 参考文献：数据结构与算法分析

