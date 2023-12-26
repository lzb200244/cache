package cache

/**
 * @Author: bbz
 * @Email: 2632141215@qq.com
 * @File: interface
 * @Date:
 * @Desc: ...
 *
 */

type CB func() (Value, error)

type ICache interface {
	// Get value from cache
	Get(key string) (interface{}, bool)
	// GetWithFunc get value from cache, if not exist, call f to get value and set
	GetWithFunc(key string, f CB) (interface{}, error)
	// Set value to cache if exist to update
	Set(key string, value Value) error // return true if set error
	// SetWithExpire set value to cache with expire time if exist to update
	SetWithExpire(key string, item *Item) error
	// Delete value from cache
	Delete(key string) bool
}

type Value interface {
	Len() int
}
