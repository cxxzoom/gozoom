package gocache

import (
	"lru"
	"sync"
)

type cache struct {
	mtx   sync.Mutex
	lru   *lru.Cache
	bytes int64
}

func (c *cache) add(key string, value ByteView) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.bytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (ByteView, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if c.lru == nil {
		return ByteView{}, false
	}

	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}

	return ByteView{}, false
}
