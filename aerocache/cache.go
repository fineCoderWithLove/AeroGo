package aerocache

import (
	"aerocache/lru"
	"sync"
)

type cache struct {
	lock       sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

// Add 方法延迟初始化对象
func (c *cache) add(key string, value ByteView) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Query(key); ok {
		return v.(ByteView), ok
	}

	return
}
