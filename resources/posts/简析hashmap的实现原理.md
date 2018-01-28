---
date: 2016-10-01 21:31:00
title: 简析hashmap的实现原理
categories:
    - java
tags:
    - java
    - code
---

  提一下哈希表，看下百科：
		  散列表（Hash table，也叫哈希表），是根据关键码值(Key value)而直接进行访问的数据结构。也就是说，它通过把关键码值映射到表中一个位置来访问记录，以加快查找的速度。这个映射函数叫做散列函数，存放记录的数组叫做散列表。
给定表M，存在函数f(key)，对任意给定的关键字值key，代入函数后若能得到包含该关键字的记录在表中的地址，则称表M为哈希(Hash）表，函数f(key)为哈希(Hash) 函数。
 简单理解：1.通过某种算法（使用key的hash算法）,计算出key的磁盘散列值，优点为速度和易用。
                     2.hashmap底层实现仍为数组（HashMap 底层就是一个数组结构，数组中的每一项又是一个链表。数组每个元素里存的是链表的表头信息，有了表头就可以遍历整个链表），当底层需要扩容，它会自动x2重新计算散列值，并把指针指向新的hashmap，而key相同的情况下，插入的value会形成一个链表
  简析实现：
  部分源码：


```
/**
     * Inflates the table.
     */
    private void inflateTable(int toSize) {
        // Find a power of 2 >= toSize
        int capacity = roundUpToPowerOf2(toSize);

        threshold = (int) Math.min(capacity * loadFactor, MAXIMUM_CAPACITY + 1);
        //transient Entry<K,V>[] table = (Entry<K,V>[]) EMPTY_TABLE table为Entry数组
        table = new Entry[capacity];
        initHashSeedAsNeeded(capacity);
    }
```
  初始化化hashmap时初始化一个Entry数组


put()方法：
```
/**
     * Associates the specified value with the specified key in this map.
     * If the map previously contained a mapping for the key, the old
     * value is replaced.
     *
     * @param key key with which the specified value is to be associated
     * @param value value to be associated with the specified key
     * @return the previous value associated with <tt>key</tt>, or
     *         <tt>null</tt> if there was no mapping for <tt>key</tt>.
     *         (A <tt>null</tt> return can also indicate that the map
     *         previously associated <tt>null</tt> with <tt>key</tt>.)
     */
    public V put(K key, V value) {
        if (table == EMPTY_TABLE) {
            inflateTable(threshold);
        }
        //其允许存放null的key和null的value，当其key为null时，调用putForNullKey方法，放入到table[0]的这个位置
        if (key == null)
            return putForNullKey(value);
            //通过调用hash方法对key进行哈希，得到哈希之后的数值。该方法实现可以通过看源码，其目的是为了尽可能的让键值对可以分不到不同的桶中，个人理解为Entry
        int hash = hash(key);
        //根据indexFor计算出在数组中的位置
        int i = indexFor(hash, table.length);
        //如果i处的Entry不为null，则通过其next指针不断遍历e元素的下一个元素。
        for (Entry<K,V> e = table[i]; e != null; e = e.next) {
            Object k;
            //key.equals(k) 当完全匹配key值时
            if (e.hash == hash && ((k = e.key) == key || key.equals(k))) {
                V oldValue = e.value;
                //替换当前新的value值，并实用链表结果进行数据储存，新加入的放在链头，最先加入的放在链尾
                e.value = value;
                e.recordAccess(this);
                return oldValue;
            }
        }
		//如果hashmap中传入的key值不存在，则进行存储并返回null
        modCount++;
        addEntry(hash, key, value, i);
        return null;
    }
```

  get()


```
/**
     * Returns the value to which the specified key is mapped,
     * or {@code null} if this map contains no mapping for the key.
     *
     * <p>More formally, if this map contains a mapping from a key
     * {@code k} to a value {@code v} such that {@code (key==null ? k==null :
     * key.equals(k))}, then this method returns {@code v}; otherwise
     * it returns {@code null}.  (There can be at most one such mapping.)
     *
     * <p>A return value of {@code null} does not <i>necessarily</i>
     * indicate that the map contains no mapping for the key; it's also
     * possible that the map explicitly maps the key to {@code null}.
     * The {@link #containsKey containsKey} operation may be used to
     * distinguish these two cases.
     *
     * @see #put(Object, Object)
     */
    public V get(Object key) {
        if (key == null)
            return getForNullKey();
        Entry<K,V> entry = getEntry(key);

        return null == entry ? null : entry.getValue();
    }

/**
     * Offloaded version of get() to look up null keys.  Null keys map
     * to index 0.  This null case is split out into separate methods
     * for the sake of performance in the two most commonly used
     * operations (get and put), but incorporated with conditionals in
     * others.
     */
    private V getForNullKey() {
	    //如果大小为0，则返回null
        if (size == 0) {
            return null;
        }
        //遍历获取null的值
        for (Entry<K,V> e = table[0]; e != null; e = e.next) {
            if (e.key == null)
                return e.value;
        }
        return null;
    }
```
		注：hashmap既允许key为null，同时也允许value为null,而hashtable是禁止的。



