---
date: 2017-03-02 23:26:00
title: 自己写GoBinarySearchTree 二叉查找树
categories:
    - golang，数据结构
tags:
    - golang，数据结构,BinarySearchTree
---

## 前言：
　　java GoBinarySearchTree的简易实现，代码中注释比较详尽，通俗易懂,注意事项亦在注解中标明。
### 正文：

1. GoBinarySearchTree.java

	```
    	package com.anteoy.dataStructuresAndAlgorithm.javav2.my;
    	
    	/**
    	 * Created by zhoudazhuang on 17-3-2.
    	 * Description: AnyType extends Comparable<? super AnyType>
    	 * 注意这里的extends 接口 不能使用implements AnyType已经是泛型 不能使用
    	 * compareTo这里是多态 root节点在插入过程中是变化的 如在root1 中插入节点root2
    	 */
    	public class GoBinarySearchTree<AnyType extends Comparable<? super AnyType>> {
    	
    	    //根节点
    	    private BinaryNode<AnyType> root;
    	
    	    /**
    	     * 节点存储tuple
    	     * @param <AnyType>
    	     */
    	    private static class BinaryNode<AnyType> {
    	        //节点存储的数据
    	        AnyType element;
    	        //左子节点
    	        BinaryNode<AnyType> left;
    	        //右子节点
    	        BinaryNode<AnyType> right;
    	
    	        BinaryNode(AnyType theElement) {
    	            this(theElement, null, null);
    	        }
    	
    	        //构造注入
    	        BinaryNode(AnyType theElement, BinaryNode<AnyType> lt,
    	                   BinaryNode<AnyType> rt) {
    	            element = theElement;
    	            left = lt;
    	            right = rt;
    	        }
    	    }
    	
    	    public GoBinarySearchTree() {
    	        makeEmpty();
    	    }
    	
    	    /**
    	     * 使树为空树
    	     */
    	    public void makeEmpty() {
    	        root = null;
    	    }
    	
    	    /**
    	     * 该树是否为空树
    	     *
    	     * @return 是否空
    	     */
    	    public boolean isEmpty() {
    	        return root == null;
    	    }
    	
    	    /**
    	     * 该树是否存在含有参数值的节点
    	     *
    	     * @param value
    	     *            元素值
    	     * @return 是否含该元素
    	     */
    	    public boolean contains(AnyType value) {
    	        return contains(value, root);
    	    }
    	
    	    /**
    	     * node节点及其子节点是否有value这个值
    	     * @param value
    	     * @param node
    	     * @return
    	     */
    	    private boolean contains(AnyType value, BinaryNode<AnyType> node) {
    	        //传入节点为空 如果不判断 则可能出现空指针异常 最后如果查找到最后一个节点 没有节点了还是没有找到 则返回false
    	        if (node == null) {
    	            return false;
    	        }
    	
    	        //多态分发compareTo
    	        int compareResult = value.compareTo(node.element);
    	        // 如果比node.element小，根据二叉查找树，递归查找左子树下
    	        if (compareResult < 0) {
    	            return contains(value, node.left);
    	            //和上面相反 这里可以解释为尾递归，但java编译器实则为对尾递归进行优化
    	        } else if (compareResult > 0) {
    	            return contains(value, node.right);
    	        } else {
    	            return true;
    	        }
    	    }
    	
    	    /**
    	     * 查找树的最小元素值
    	     * @return 最小元素值
    	     */
    	    public AnyType findMin() {
    	        if (isEmpty()) {
    	            throw new NullPointerException();
    	        }
    	        //返回最小节点的元素值
    	        return findMin(root).element;
    	    }
    	
    	
    	    /**
    	     * 查找以node为树的最小子节点
    	     * @param node root节点
    	     * @return 最小子节点
    	     */
    	    private BinaryNode<AnyType> findMin(BinaryNode<AnyType> node) {
    	        if (node == null) {
    	            return null;
    	            //最小节点
    	        } else if (node.left == null) {
    	            return node;
    	        }
    	        //尾递归查找最左边的后代节点 直到node.left == null
    	        return findMin(node.left);
    	    }
    	
    	    /**
    	     * 查找该树最大元素值
    	     * @return 最大元素值
    	     */
    	    public AnyType findMax() {
    	        if (isEmpty()) {
    	            throw new NullPointerException();
    	        }
    	        return findMavalue(root).element;
    	    }
    	
    	
    	
    	    /**
    	     * 查找某节点及其子树中的最大元素 这里使用while循环替代递归 递归方法在下面注解中给出
    	     * @param root
    	     * @return
    	     */
    	    private BinaryNode<AnyType> findMavalue(BinaryNode<AnyType> root) {
    	        if(root != null){
    	            while (root.right != null){
    	                root = root.right;
    	            }
    	        }
    	        //如果root为空 会直接返回root即null 所以不用赘余判断
    	        return root;
    	    }
    	    //递归方法
    	//    private BinaryNode<AnyType> findMavalue(BinaryNode<AnyType> root) {
    	//        if (root == null) {
    	//            return null;
    	//        } else if (root.right == null) {
    	//            return root;
    	//        }
    	//        return findMavalue(root.right);
    	//    }
    	
    	    /**
    	     * 插入元素
    	     * @param value
    	     */
    	    public void insert(AnyType value) {
    	        root = insert(value, root);
    	    }
    	
    	    /**
    	     * root node下插入元素value
    	     * @param value
    	     * @param root
    	     * @return 元素插入的节点
    	     */
    	    private BinaryNode<AnyType> insert(AnyType value, BinaryNode<AnyType> root) {
    	        //1. 空树直接返回一个节点 2. 递归终结条件  binaryNode.left 或者binaryNode.right为空了 new一个节点（树叶）并返回 到下面 root.left = insert(v
    	        if (root == null) {
    	            return new BinaryNode<AnyType>(value);
    	        }
    	        //和root节点比较大小
    	        int compareResult = value.compareTo(root.element);
    	        //往左边递归计算 root节点重新分配  这里构造关联关系
    	        if (compareResult < 0) {
    	            root.left = insert(value, root.left);
    	        //和上面相反
    	        } else if (compareResult > 0) {
    	            root.right = insert(value, root.right);
    	        }
    	        return root;
    	    }
    	
    	    /**
    	     * 删除某元素
    	     * @param value
    	     */
    	    public void remove(AnyType value) {
    	        root = remove(value, root);
    	    }
    	
    	    /**
    	     * 在某个节点下删除元素
    	     * @param value
    	     * @param root
    	     * @return 被删除元素的节点
    	     */
    	    private BinaryNode<AnyType> remove(AnyType value, BinaryNode<AnyType> root) {
    	        if (root == null) {
    	            return root;
    	        }
    	        int compareResult = value.compareTo(root.element);
    	        //如果是左子树
    	        if (compareResult < 0) {
    	            //递归到最左边
    	            root.left = remove(value, root.left);
    	        } else if (compareResult > 0) {
    	            //相反
    	            root.right = remove(value, root.right);
    	            //情况一 此节点左右两个儿子都存在
    	        } else if (root.left != null && root.right != null) {//root此时就是要求被删除节点
    	            //查找右边的节点的最小值 和被删除节点交换
    	            root.element = findMin(root.right).element;
    	            //交换完毕 设置右边的值为删除以后的节点的右子节点的最小值
    	            root.right = remove(root.element, root.right);
    	        } else {
    	            //该节点只有左右儿子中的其中一个 直接使用唯一儿子代替
    	            root = (root.left != null) ? root.left : root.right;
    	        }
    	        return root;
    	    }
    	
    	    /**
    	     * 遍历输出树
    	     */
    	    public void printTree() {
    	        if (isEmpty()) {
    	            System.out.println("Empty tree");
    	        } else {
    	            printTree(root);
    	        }
    	    }
    	
    	    /**
    	     * 递归输出打印
    	     * 先输出最小的左子节点 在在每一个递归回归里面 进行递归右边节点的操作
    	     * 达到由小到大输出
    	     * @param root
    	     */
    	    private void printTree(BinaryNode<AnyType> root) {
    	        if (root != null) {
    	            printTree(root.left);
    	            System.out.println(root.element);
    	            printTree(root.right);
    	        }
    	    }
    	
    	}
	```

2. GoBinarySearchTreeTest.java
    ```
    	package com.anteoy.dataStructuresAndAlgorithm.javav2.my;
        
        /**
         * Created by zhoudazhuang on 17-3-2.
         * GoBinarySearchTreeTest
         */
        public class GoBinarySearchTreeTest {
            public static void main(String[] args) {
                GoBinarySearchTree<Integer> tree = new GoBinarySearchTree<Integer>();
                tree.insert(2);
                tree.insert(1);
                tree.insert(5);
                tree.insert(4);
                tree.insert(3);
                tree.printTree();
                System.out.println(" ");
                tree.remove(2);
                tree.remove(4);
                tree.printTree();
            }
        }
    
    ```
	
输出结果：
```
1
2
3
4
5

1
3
5

Process finished with exit code 0
```
### 后记：
1. 参考文献：数据结构与算法分析

