package deliver

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
	if ttl > 0 {
		go func() {
			time.Sleep(ttl)
			c.mu.Lock()
			defer c.mu.Unlock()
			delete(c.data, key)
		}()
	}
}
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.data[key]
	return value, ok
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}
