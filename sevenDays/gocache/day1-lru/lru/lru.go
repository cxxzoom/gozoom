package lru

import "C"
import "container/list"

type Cache struct {
	maxBytes  int64      // 允许使用的最大内存
	nBytes    int64      // 已经使用的内存
	ll        *list.List // 双向链表
	cache     map[string]*list.Element
	onEvicted func(key string, value Value) //
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

// New 实例化cache
func New(maxBytes int64, fn func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		onEvicted: fn,
	}
}

// Get 获取key的value
func (c *Cache) Get(key string) (Value, bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		val := ele.Value.(*entry)
		return val.value, true
	}

	return nil, false
}

// RemoveOld 移除队尾的节点
func (c *Cache) RemoveOld() {
	// 获取队尾的节点
	ele := c.ll.Back()
	// 因为Value是interface，使用断言获取值
	l := ele.Value.(*entry)
	// 移除这个节点
	delete(c.cache, l.key)
	// 更新缓存已使用的容量
	c.nBytes -= int64(len(l.key)) + int64(l.value.Len())
	// 如果回调函数不为空，则调用：
	//TODO 暂时还没明白这里的回调函数的使用和实现
	if c.onEvicted != nil {
		c.onEvicted(l.key, l.value)
	}
}

// Add 新增或者修改
func (c *Cache) Add(key string, value Value) {
	// 如果读取到值说明存在，则执行修改
	if ele, ok := c.cache[key]; ok {
		// 这里要和删除对应起来看；删除是获取Back并移除
		c.ll.MoveToFront(ele)
		// 这里还要使用断言获取数据
		val := ele.Value.(*entry)
		// 更新使用了的缓存
		c.nBytes += int64(value.Len()) + int64(val.value.Len())
		// 更新值
		ele.Value = value
	} else {
		// 这玩意儿似乎是个double linked list的一个元素
		ele := c.ll.PushFront(&entry{key: key, value: value})
		// 把链表的元素放在cache里
		c.cache[key] = ele
		// 更新已经使用的大小
		c.nBytes += int64(len(key)) + int64(value.Len())
	}

	// 如果设置了最大内存，并且使用超出了内存大小
	// 则执行移除逻辑
	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.RemoveOld()
	}
}

// Len 获取当前缓存的条数
func (c *Cache) Len() int {
	// 链表已经实现了这个功能，直接用就行
	return c.ll.Len()
}
