package LRU

import (
	"fmt"
	"testing"
)

/**
越靠近链表尾部的结点是越早之前访问的。越先删除
*/

func TestLRU(t *testing.T) {
	cache := newLRUCacheByCapacity(6)
	cache.Put(1, 1)
	cache.printLinkedList()
	fmt.Println("-------------------")
	cache.Put(2, "2abc")
	cache.printLinkedList()
	fmt.Println("-------------------")
	cache.Put(3, "3中文")
	cache.printLinkedList()
	fmt.Println("-------------------")
	cache.Put(4, 4)
	cache.printLinkedList()
	fmt.Println("-------------------")
	cache.Put(1, "更新了1")
	cache.Put(5, "第五个")
	cache.Put(6, "第6个")
	cache.Put(6, "第6个更新")
	cache.Put(2, "2更新")
	cache.Put(12, "12")
	cache.printLinkedList()
}

type LRUCache struct {
	length   int // 链表当前存放多少个节点
	capacity int // 链表最多有多少个节点
	head     *linkNode
}

/**
单链表节点
*/
type linkNode struct {
	key   int
	value interface{}
	next  *linkNode
}

func newLRUCacheByCapacity(capacity int) LRUCache {
	// 初始化链表，默认容量6，初始没有数据存放
	cache := LRUCache{
		capacity: capacity,
		length:   0,
	}
	return cache
}

/**
单链表访问O(n)
*/
func (this *LRUCache) Get(key int) *linkNode {
	if this.length <= 0 {
		return nil
	}
	node := this.head
	i := 0
	for i < this.length {
		if node.key == key {
			return node
		} else if node.next != nil {
			node = node.next
			i++
		} else {
			return nil
		}
	}
	return nil
}

func (this *LRUCache) getPrev(key int) *linkNode {
	if this.length > 0 {
		node := this.head
		for node.next != nil {
			if node.next.key == key {
				return node
			}
			node = node.next
		}
	}
	return nil
}

/**
获取链表最后一个节点的指针
*/
func (this *LRUCache) getLast() *linkNode {
	if this.length > 0 {
		node := this.head
		for node.next != nil {
			node = node.next
		}
		return node
	}
	return nil
}

func (this *LRUCache) getFirst() *linkNode {
	if this.length > 0 {
		return this.head
	}
	return nil
}

func (this *LRUCache) Put(key int, value interface{}) {
	// 查询链表中是否有当前元素，如果有，则将该元素放到当前链表头部
	node := this.Get(key)
	if node != nil {
		// 链表中有该key的话，将该元素挪至链表头部
		if this.head.key == key {
			this.head.value = value
			return
		}
		// 如果不是第一个元素，将该元素挪至链表头部
		// 获取前驱，断开当前节点,将前驱的后继，指向后继
		prevPtr := this.getPrev(key)
		prevPtr.next = node.next
		// 将元素放到链表头部
		firstPtr := this.getFirst()
		node.value = value
		node.next = firstPtr
		this.head = node
	} else {
		// 数据转换失败，数据无效，认为节点中没有该数据，将元素放到链表头部
		this.newHead(key, value)
	}
}

/**
新增元素时，如果链表未满，则直接将尾节点的后继指向新节点。
如果链表满了，需要删除尾结点，然后将新节点插入到链表尾部(即将尾结点的前驱的后继指向新节点，旧的尾部节点等待垃圾回收即可)
*/
func (this *LRUCache) newHead(key int, value interface{}) {
	newNode := linkNode{key: key, value: value, next: nil}
	if this.length == 0 {
		this.head = &newNode
		this.length++
		return
	}
	// 1. 判断链表是否已满
	if this.length == this.capacity {
		// 满了删除尾结点，然后将元素放到头结点位置
		lastNode := this.getLast()
		lastPrev := this.getPrev(lastNode.key)
		lastPrev.next = nil

		// 新加头结点
		firstNode := this.getFirst()
		newNode.next = firstNode
		this.head = &newNode
	} else {
		// 链表没满
		firstNode := this.getFirst()
		newNode.next = firstNode
		this.head = &newNode
		this.length++
	}
}

func (list LRUCache) printLinkedList() {
	node := list.head
	for node != nil {
		fmt.Printf("linked list node:%v\n", node)
		node = node.next
	}
}
