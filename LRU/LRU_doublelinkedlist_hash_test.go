package LRU

import (
	"fmt"
	"testing"
)

/**
    LRU算法，链表+哈希表解决方法：
	借助hash表，把LRU缓存淘汰算法的时间复杂度，降低为O(1)
    解决方案：使用双向链表存储数据，双向链表新增一个hNext指针，指向哈希表中因冲突后产生的单链表中，单链表存放该节点的地址
	模拟哈希：自定义哈希表结构为哈希表+单链表，自定义哈希函数使得哈希碰撞频率加大，模拟产生冲突
*/

type LRUDllHashMap struct {
	myMap     CustomHashMap
	myDllList DoubleLinkedList
}

/*
双向链表
*/
type DoubleLinkedList struct {
	head     *DllNode
	tail     *DllNode
	length   int // 双向链表中当前存放多少个节点
	capacity int // 双向链表最多有多少个节点
}

func Constructor(capacity int) DoubleLinkedList {
	// 初始化链表，默认容量6，初始没有数据存放
	dll := DoubleLinkedList{
		capacity: capacity,
		length:   0,
	}
	return dll
}

/*
双向链表节点
*/
type DllNode struct {
	key   string
	value interface{}
	prev  *DllNode
	next  *DllNode
}

func (this *LRUDllHashMap) put(key string, value interface{}) {
	ptr := this.myMap.get(key) //如果缓存中有对应数据，返回值为双向链表中的node指针
	if ptr != nil {
		nodePtr := ptr.(*DllNode) // 将hashmap的value断言为双节点的指针
		if nodePtr != nil {
			// 断开node的前驱和后继，将其放到双向链表头部
			nodePtr.prev.next = nodePtr.next
			nodePtr.next.prev = nodePtr.prev
			// 重新设置头结点的前驱，尾结点的后继
			this.myDllList.head.prev = nodePtr
			this.myDllList.tail.next = nodePtr
			// 将节点设置为链表的头
			nodePtr.next = this.myDllList.head
			nodePtr.prev = this.myDllList.tail

			// 设置链表的头和尾
			this.myDllList.head = nodePtr
			this.myDllList.tail = nodePtr.prev
		} else if nodePtr == nil {
			newNodePtr := this.myDllList.newNode(key, value)
			this.myMap.put(key, newNodePtr)
		}
	} else {
		//如果hashmap中没有数据，说明是新缓存
		newNodePtr := this.myDllList.newNode(key, value)
		this.myMap.put(key, newNodePtr)
	}
}

func (this *DoubleLinkedList) newNode(key string, value interface{}) *DllNode {
	// 哈希表中没有该元素
	newNode := DllNode{
		key:   key,
		value: value,
	}
	if this.length == this.capacity {
		// 双向链表已满
		newNode.prev = this.tail.prev
		this.tail.prev.next = &newNode
		this.head.prev = &newNode
		newNode.next = this.head
		this.head = &newNode
		this.tail = newNode.prev

	} else if this.length == 0 {
		this.head = &newNode
		this.tail = &newNode
		this.length++
	} else {
		//双向链表未满
		newNode.prev = this.tail
		newNode.next = this.head
		this.tail.next = &newNode
		this.head.prev = &newNode
		this.head = &newNode
		this.length++
	}
	return &newNode
}

func TestCustomHashmap(t *testing.T) {
	var temp LRUDllHashMap = LRUDllHashMap{
		myMap:     NewHashMapDefault(),
		myDllList: Constructor(4),
	}
	temp.put("1", "cache1")
	temp.put("2", "cache2")
	temp.put("3", "cache3")
	temp.put("1", "cache111")
	//temp.print()
	fmt.Println("-------分隔线-------")
	temp.put("4", "cache4")
	temp.print()
	fmt.Println("-------分隔线-------")
	temp.put("5", "cache4")
	temp.put("6", "cache4")
	temp.put("7", "cache4")
	temp.print()
}

func (this *LRUDllHashMap) print() {
	this.myMap.printHashMap()
	this.myDllList.toString()
}

func (this *DoubleLinkedList) toString() {
	node := this.head
	fmt.Println("---doublelinkedlist start---")
	for node != nil {
		fmt.Printf("{key:%v,value:%v}\n", node.key, node.value)
		node = node.next
		if node == this.head {
			return
		}
	}
	fmt.Println("---doublelinkedlist end---")
}
