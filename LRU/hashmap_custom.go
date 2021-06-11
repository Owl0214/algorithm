package LRU

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"study/algorithm/util"
)

/**
自定义哈希Map，方便演示，定义hash数组长度为3，这样容易发生碰撞
*/

type CustomHashMap struct {
	arr []*SingleLinkList //存放单链表头结点地址
}

func NewHashMapDefault() CustomHashMap {
	// 初始化数组
	myHash := CustomHashMap{arr: make([]*SingleLinkList, 3)}
	return myHash
}

type SingleLinkList struct {
	head *SingleLinkNode
}

/**
单链表
*/
type SingleLinkNode struct {
	key   interface{}
	value interface{}
	next  *SingleLinkNode
}

func (this *SingleLinkNode) getSingleNode(key interface{}) *SingleLinkNode {
	node := this
	for node != nil {
		if node.key == key {
			return node
		}
		node = node.next
	}
	return node
}

func (this *SingleLinkList) getKV(key interface{}) interface{} {
	nodePtr := this.head.getSingleNode(key)
	if nodePtr != nil {
		return &nodePtr.value
	}
	return nil
}

/**
往单链表中添加元素
*/
func (this *SingleLinkList) putKV(key interface{}, value interface{}) {
	node := this.head
	// 空链表
	if node == nil {
		newNode := SingleLinkNode{
			key:   key,
			value: value,
			next:  nil,
		}
		this.head = &newNode
		return
	}
	// 非空链表
	for node != nil {
		// 链表中有该key，就替换value值
		if node.key == key {
			node.value = value
			return
		}
		// 链表中没有该key，就在链表尾部新加一个节点
		if node.next != nil {
			node = node.next
		} else {
			newNode := SingleLinkNode{
				key:   key,
				value: value,
				next:  nil,
			}
			node.next = &newNode
			return
		}
	}
}

/**
 * hashmap put方法
 */
func (this *CustomHashMap) put(key interface{}, value interface{}) {
	hashcode := this.calHash(key)
	if this.arr[hashcode] == nil {
		newList := SingleLinkList{
			head: nil,
		}
		this.arr[hashcode] = &newList
	}

	this.arr[hashcode].putKV(key, value)
}

/**
 * hashmap get方法
 */
func (this *CustomHashMap) get(key interface{}) interface{} {
	hashcode := this.calHash(key)
	linkList := this.arr[hashcode]
	if linkList != nil {
		node := linkList.head
		for node != nil {
			if node.key == key {
				return node.value
			}
			node = node.next
		}
	}
	return nil
}

/**
 * 自定义哈希算法
 */
func (this *CustomHashMap) calHash(key interface{}) int {
	keyByte, _ := this.encode(key)
	result := util.BytesToInt(keyByte)
	return result % 3
}

func (this *CustomHashMap) encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	// 取前16个字节
	result := buf.Bytes()
	result = result[:16]
	return result, nil
}

func (this *CustomHashMap) printHashMap() {
	fmt.Println("------hashmap start-------")
	for index, singleList := range this.arr {
		fmt.Printf("------hashmap slot %d-------\n", index)
		nodePtr := singleList.head
		for nodePtr != nil {
			fmt.Printf("{key:%v,value:%v}\n", nodePtr.key, nodePtr.value)
			nodePtr = nodePtr.next
		}
	}
	fmt.Println("------hashmap end-------")
}
