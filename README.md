[数据结构与算法学习](https://owl0214.github.io/)
# LRU
## 双向链表+hash实现说明

以数组+单链表的形式，自定义了hashmap

在实现中，自定义了散列函数，目的是为了容易产生碰撞，模拟哈希冲突

hashmap中，key值为LRU的key值，value值存放了双向链表节点的地址

