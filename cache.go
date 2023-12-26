package cache

import (
	"errors"
	"sync"
	"time"
)

/**
 * @Author: bbz
 * @Email: 2632141215@qq.com
 * @File: cache
 * @Date:
 * @Desc: ...
 *
 */

type Item struct {
	Val        Value
	ExpireTime int64
}

// IsExpired check if item is expired
func (i *Item) IsExpired() bool {
	if i.ExpireTime == 0 {
		return false
	}
	return time.Now().UnixMilli() > i.ExpireTime
}

type Cache struct {
	data   map[string]*Item
	used   int64
	max    int64
	mu     sync.RWMutex
	loader *Group
}

func (c *Cache) Used() int64 {
	return c.used
}

func NewCache(max int64) *Cache {
	return &Cache{
		data: make(map[string]*Item),
		max:  max,
		loader: &Group{
			m: make(map[string]*call),
		},
	}

}
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	v, ok := c.data[key]
	if !ok {
		c.mu.RUnlock()
		return nil, false
	}
	if v.IsExpired() {
		// lazy delete
		c.mu.RUnlock()
		c.Delete(key)
		return nil, false
	}
	c.mu.RUnlock()
	return v.Val, true
}

func (c *Cache) GetWithFunc(key string, f CB) (interface{}, error) {
	if v, ok := c.Get(key); ok {
		return v, nil
	}
	if f == nil {
		return nil, errors.New("f is nil")
	}
	// use singleflight to avoid cache breakdown
	v, err := c.loader.Do(key, f)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	err = c.Set(key, v)
	if err != nil {
		return nil, err
	}
	return v, nil

}

func (c *Cache) Set(key string, value Value) error {
	return c.SetWithExpire(key, &Item{Val: value})

}

func (c *Cache) SetWithExpire(key string, item *Item) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	//check if exist then update
	if v, ok := c.data[key]; ok {
		v.Val = item.Val
		v.ExpireTime = item.ExpireTime
		// valid overflow
		if c.used+int64(item.Val.Len()-v.Val.Len()) > c.max {
			return errors.New("cache over limit")
		}
		c.used = c.used + int64(item.Val.Len()-v.Val.Len())
		c.data[key] = v
		return nil
	}
	if c.used+int64(item.Val.Len()) > c.max {
		return errors.New("cache over limit")
	}
	c.used = c.used + int64(item.Val.Len())
	item.ExpireTime = time.Now().UnixMilli() + item.ExpireTime

	c.data[key] = item
	return nil
}

func (c *Cache) Delete(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.data[key]; ok {
		delete(c.data, key)
		c.used = c.used - int64(v.Val.Len())
		return true
	}
	return false
}
