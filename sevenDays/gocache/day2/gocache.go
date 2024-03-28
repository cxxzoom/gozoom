package gocache

import (
	"fmt"
	"log"
	"sync"
)

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (g GetterFunc) Get(key string) ([]byte, error) {
	return g(key)
}

// Group 这是一个缓存命名空间，加载数据并分布在其中
// Is a namespace and associated data loaded spread over
type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

var (
	mtx    sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, bytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mtx.Lock()
	defer mtx.Unlock()

	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{bytes: bytes},
	}
	groups[name] = g

	return g
}

// 为什么这里要读锁：
// 因为是并发的，写go要有并发思维
// 多个goroutine的读写操作可能会到值数据不一致；并且有可能崩溃
func GetGroup(name string) *Group {
	mtx.RLock()
	g := groups[name]
	mtx.RUnlock()
	return g
}

func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}

	if v, ok := g.mainCache.get(key); ok {
		log.Println("[GoCache] Hit!")
		return v, nil
	}

	return g.load(key)
}

// 加载数据，盲猜是从Getter加
func (g *Group) load(key string) (ByteView, error) {
	return g.getLocally(key)
}

// 先看NewGroup，这是实例化的时候，传入了一个getter Getter
// 然后这里没命中数据的时候，通过调用这个getter获取数据
// 然后把数据存到链表中，即populateCache
func (g *Group) getLocally(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, nil
	}

	val := ByteView{b: CloneBytes(bytes)}
	g.populateCache(key, val)
	return val, nil
}

// 把数据缓存到链表里
func (g *Group) populateCache(key string, val ByteView) {
	g.mainCache.add(key, val)
}
