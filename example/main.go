package main

import (
	"cache"
	"fmt"
	"log"
	"time"
)

/**
 * @Author: bbz
 * @Email: 2632141215@qq.com
 * @File: main
 * @Date:
 * @Desc: ...
 *
 */
func main() {
	// Create a new cache with a maximum size of 5 bytes
	c := cache.NewCache(5)
	err := c.Set("key", cache.String("hello"))
	if err != nil {
		log.Fatal(err)
		return
	}
	v, ok := c.Get("key")
	if ok {
		// Output: hello
		fmt.Println(v.(cache.String))
	}
	val, err := c.GetWithFunc("key", func() (cache.Value, error) {
		return cache.String("world"), nil
	})
	if err != nil {
		return
	}
	// Output: hello because the key was already in the cache no call the function
	fmt.Println(val.(cache.String))

	fmt.Println(c.Used()) // Output: 5
	c.Delete("key")
	fmt.Println(c.Used()) // Output: 0

	// use expire
	c.SetWithExpire("key", &cache.Item{
		Val:        cache.String("hello"),
		ExpireTime: 2000, // 2s
	})
	time.Sleep(time.Second * 3)
	v, ok = c.Get("key")
	fmt.Println(ok) // Output: false because the key was expired

	// use get with func
	v, err = c.GetWithFunc("key", func() (cache.Value, error) {
		return cache.String("world"), nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v.(cache.String)) // Output: 5
	c2 := cache.NewCache(8)
	v2, err := c2.GetWithFunc("key", func() (cache.Value, error) {
		return cache.Int(5), nil
	})
	if err != nil {
		return
	}
	fmt.Println(v2)
	fmt.Println(c2.Used())

}
