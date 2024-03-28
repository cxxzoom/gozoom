package lru

import "container/list"

type Cache struct {
	maxBytes  int64                         // 允许使用的最大内存
	nBytes    int64                         // 已经使用的内存
	ll        *list.List                    // 双向链表
	cache     map[string]*list.Element      // 缓存，保存双向链表的指针
	onEvicted func(key string, value Value) //回调函数
}

type entry struct {
	key   string
	value Value
}

// Value 包含的是list.List 的Len方法
// 然后list.List里面有这个方法，所以可以调用 .Len()
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
// 获取到值，意思热度上升，因为移除是溢出的.Back的节点
func (c *Cache) Get(key string) (Value, bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		val := ele.Value.(*entry)
		return val.value, true
	}

	return nil, false
}

// RemoveOldest 移除队尾的节点
func (c *Cache) RemoveOldest() {
	// 获取队尾的节点
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.onEvicted != nil {
			c.onEvicted(kv.key, kv.value)
		}
	}
	// 因为Value是interface，使用断言获取值
}

// Add 新增或者修改
func (c *Cache) Add(key string, value Value) {
	// 如果读取到值说明存在，则执行修改
	if ele, ok := c.cache[key]; ok {
		// 这里要和删除对应起来看；删除是获取Back并移除
		// 这里是修改
		c.ll.MoveToFront(ele)
		// 这里还要使用断言获取数据
		val := ele.Value.(*entry)
		// 更新使用了的缓存
		c.nBytes += int64(value.Len()) - int64(val.value.Len())
		// 更新值
		ele.Value = value
	} else {
		// 这玩意儿似乎是个double linked list的一个元素
		// 确实是移动到前面，因为删除是删除的.Back()
		ele := c.ll.PushFront(&entry{key: key, value: value})
		// 把链表的元素放在cache里
		c.cache[key] = ele
		// 更新已经使用的大小
		c.nBytes += int64(len(key)) + int64(value.Len())
	}

	// 如果设置了最大内存，并且使用超出了内存大小
	// 则执行移除逻辑
	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.RemoveOldest()
	}
}

// Len 获取当前缓存的条数
func (c *Cache) Len() int {
	// 链表已经实现了这个功能，直接用就行
	return c.ll.Len()
}
