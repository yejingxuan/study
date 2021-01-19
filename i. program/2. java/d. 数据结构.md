# 数组学习

1. 数组创建的三种方法
    ```java
    int[] disOrderArry = new int[] {2, 5, 8, 7, 3, 1};
    int[] disOrderArry2 = {2, 5, 8, 7, 3, 1};
    int[] disOrderArry3 = new int[5]
    ```

2. JDK自带的数组工具类 java.util.Arrays

- toString方法

    ```java
    public static String toString(int[] a) {
        if (a == null)
            return "null";
        int iMax = a.length - 1;
        if (iMax == -1)
            return "[]";

        StringBuilder b = new StringBuilder();
        b.append('[');
        for (int i = 0; ; i++) {
            b.append(a[i]);
            if (i == iMax)
                return b.append(']').toString();
            b.append(", ");
        }
    } 
    ```

# java实现单向链表

1. 定义Node节点

        public class Node {	
            public int data;           
            public Node next;

            public Node(int data) {
                super();
                this.data = data;
            }         
        }

2. 定义链表结构

        public class LinkList {
            //头节点
            public Node head;
            //当前节点
            public Node current;
        }


3. 实现链表的节点添加

        public void add(Node node){
            //如果head为null则表示添加的是第一个
            if(head == null){
                head = node;
                current = head;
            }else{
                current.next = node;
                //当前节点向后移一位
                current = current.next;
            }
        }


4. 实现链表的遍历

        public void getAll(Node node){
            if(node == null){
                return ;
            }else{
                current = node;
                while (current != null){
                    System.out.println(current.data);
                    current = current.next;
                }
            }
        }

5. 判断链表是否有环

        public boolean hasCycle(Node head) {
            Node first = head;
            Node second = head;
            
            while (second != null && second.next!=null) {
                first = first.next;
                second = second.next.next;
                if(first == second) {
                    return true;
                }			
            }
            return false;	
        }

# java实现二叉树

### 数据结构参考文章：https://blog.csdn.net/javazejian/article/details/53892797

## 一、 二叉树（查找二叉树）的基本实现

1. 节点定义
    ```java
    public class TreeNode {
        //存储的数据
        public Integer data;
        //左节点的引用
        public TreeNode left;
        //右节点的引用
        public TreeNode right;
        
        public TreeNode() {
            super();
        }
        public TreeNode(Integer data) {
            super();
            this.data = data;
        }
    }
    ```

2. 二叉树定义

    ```java
    public class BinaryTree {
        //根节点
        public TreeNode root;
        //节点数量
        public int size = 0;

        public BinaryTree(TreeNode root) {
            super();
            this.root = root;
        }
        public BinaryTree() {
            super();
        }
    }
    ```

3. 二叉树子节点的添加

    ```java
    public void addNode(TreeNode node) {
        //如果根节点为null
        if(root == null) {
            root = node;
            size++;
            return;
        }else {
            //添加子节点
            int ret = 0;
            TreeNode current = root;
            TreeNode currentFather = root;
            while(current != null) {
                ret = node.data.compareTo(current.data);
                currentFather = current;
                if(ret < 0) {
                    current = current.left;
                }else if(ret > 0) {
                    current = current.right;
                }else if(ret == 0) {
                    current = node;
                }				
            }
                            
            if(ret < 0) {
                currentFather.left = node;
            }else {
                currentFather.right = node;
            }
            size++;		
        }	
    }
    ```


4. 二叉树遍历

    ```java
        //前序遍历
        //访问根结点；遍历左子树；遍历右子树。	
        public void	prevIterator(TreeNode node) {
            if(node != null) {
                System.out.println(node.data);
                midIterator(node.left);		
                midIterator(node.right);
            }
        }

        //后序遍历
        // 遍历左子树；遍历右子树；访问根结点。	
        public void	nextIterator(TreeNode node) {
            if(node != null) {		
                midIterator(node.left);		
                midIterator(node.right);
                System.out.println(node.data);
            }
        }
        
        //中序遍历
        //遍历左子树；访问根结点；遍历右子树。
        public void midIterator(TreeNode node) {
            if(node != null) {
                midIterator(node.left);
                System.out.println(node.data);
                midIterator(node.right);
            }
        }

![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20200629173221.png)


##  二、平衡二叉树