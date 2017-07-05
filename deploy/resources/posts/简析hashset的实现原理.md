---
date: 2016-10-01 21:46:00
title: 简析hashset的实现原理
categories:
    - java
tags:
---

hashset底层为hashmap。
源码如下：

```
/**
     * Constructs a new, empty set; the backing <tt>HashMap</tt> instance has
     * default initial capacity (16) and load factor (0.75).
     */
    public HashSet() {
        map = new HashMap<>();
    }
```
默认 initial capacity（hashmap底层数组大小）为16，load factor 为 0.75

add() 方法

```
/**
     * Adds the specified element to this set if it is not already present.
     * More formally, adds the specified element <tt>e</tt> to this set if
     * this set contains no element <tt>e2</tt> such that
     * <tt>(e==null&nbsp;?&nbsp;e2==null&nbsp;:&nbsp;e.equals(e2))</tt>.
     * If this set already contains the element, the call leaves the set
     * unchanged and returns <tt>false</tt>.
     *
     * 将添加到此set中的元素
     * @param e element to be added to this set
     * 如果此set尚未包含指定元素，则返回true。
     * @return <tt>true</tt> if this set did not already contain the specified
     * element
     */
    public boolean add(E e) {
    //HashSet 中的 value 都是PRESENT
        return map.put(e, PRESENT)==null;
    }
```

   放入 HashMap 的 Entry 中 key 与集合中原有 Entry 的 key 相同（hashCode()返回值相等，通过 equals 比较也返回 true），value被覆盖形成链表，但 key 不会有任何改变，这也就满足了 Set 中元素不重复的特性。

该方法如果添加的是在 HashSet 中不存在的，则返回 true；如果添加的元素已经存在，返回 false。因为map.put()如果key不存在，则会存如hashmap中并返回null，即为 return null==null 为true

remove()方法源码：

```
/**
     * Removes the specified element from this set if it is present.
     * More formally, removes an element <tt>e</tt> such that
     * <tt>(o==null&nbsp;?&nbsp;e==null&nbsp;:&nbsp;o.equals(e))</tt>,
     * if this set contains such an element.  Returns <tt>true</tt> if
     * this set contained the element (or equivalently, if this set
     * changed as a result of the call).  (This set will not contain the
     * element once the call returns.)
     *
     * @param o object to be removed from this set, if present
     * @return <tt>true</tt> if the set contained the specified element
     * 如果set包含这个元素，则返回true,否则返回false
     */
    public boolean remove(Object o) {
        return map.remove(o)==PRESENT;
    }
```

