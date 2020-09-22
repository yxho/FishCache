package lru

import "container/list"

//LRU(Least Recently Used)
// 最近最少访问(Least Recently Used, LRU)
//最近最少使用，相对于仅考虑时间因素的 FIFO 和仅考虑访问频率的 LFU，LRU 算法可以认为是相对平衡的一种淘汰算法。
//LRU 认为，如果数据最近被访问过，那么将来被访问的概率也会更高。LRU 算法的实现非常简单，维护一个队列，如果某条记录被访问了，则移动到队尾，
//那么队首则是最近最少访问的数据，淘汰该条记录即可。

// LRU缓存
type Cache struct {
	maxBytes  int64      //允许使用的最大内存
	nowbytes  int64      //当前已用的内存
	ll        *list.List //双向链表
	cache     map[string]*list.Element
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

// Len() int，用于返回值所占用的内存大小
type Value interface {
	Len() int
}

func (c *Cache) Len() int {
	return c.ll.Len()
}

// 实例化Cache
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

//查找功能。
//主要有 2 个步骤，第一步是从字典中找到对应的节点，第二步，将该节点移动到队尾。
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToBack(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

//删除
//删除最近最少访问节点（队首节点）
func (c *Cache) RemoveOldest() {
	ele := c.ll.Front()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nowbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

//新增
//如果键存在，则更新对应节点的值，并将该节点移到队尾。
//不存在则新增该节点
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToBack(ele)
		kv := ele.Value.(*entry)
		c.nowbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushBack(&entry{key, value})
		c.cache[key] = ele
		c.nowbytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nowbytes {
		c.RemoveOldest()
	}
}

