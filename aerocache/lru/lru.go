package lru

import (
	"container/list"
	"fmt"
)

//LRU缓存，字典和双向链表

type Cache struct {
	maxBytes  int64 //最大使用内存
	nbytes    int64 //当前已经使用的内存
	ll        *list.List
	cache     map[string]*list.Element
	OnEvicted func(key string, value Value) //记录被移除时候的回调函数
}

type entry struct {
	key   string
	value Value
}
type Value interface {
	Len() int
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

// 查找
func (c *Cache) Query(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// 缓存淘汰
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		//更新缓存大小
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// 新增缓存
func (c *Cache) Add(key string, value Value) {
	if element, ok := c.cache[key]; ok {
		//缓存已经存在,移动缓存在队列中的位置
		c.ll.MoveToFront(element)
		//获取val
		kv := element.Value.(*entry)
		//更新当前内存
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		element := c.ll.PushFront(&entry{key, value})
		c.cache[key] = element
		//增加内存keyAndVal
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	//如果内存爆满了，移除最老的元素
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

// 查询元素多少
func (c *Cache) Len() int {
	return c.ll.Len()
}

// 遍历并打印LRU缓存的双向链表 ll
func (c *Cache) PrintList() {
	fmt.Println("LRU Cache List:")
	for element := c.ll.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*entry)
		fmt.Printf("Key: %s, Value: %v\n", entry.key, entry.value)
	}
	fmt.Println()
}

// 遍历并打印LRU缓存的字典 cache
func (c *Cache) PrintCache() {
	fmt.Println("LRU Cache:")
	for key, element := range c.cache {
		entry := element.Value.(*entry)
		fmt.Printf("Key: %s, Value: %v\n", key, entry.value)
	}
	fmt.Println()
}
