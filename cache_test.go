package cache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
 * @Author: bbz
 * @Email: 2632141215@qq.com
 * @File: cache_test
 * @Date:
 * @Desc: ...
 *
 */

func TestNewCache(t *testing.T) {
	c := NewCache(10)
	err := c.Set("key", String("hello"))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(c.Get("key"))
	err2 := c.Set("key2", String("world"))
	if err2 != nil {
		t.Error(err2)
		return
	}
	fmt.Println(c.Get("key2"))
	err3 := c.Set("key3", String("world"))
	if err3 == nil {
		t.Error(err3)
		return
	}
	if err3.Error() != "cache over limit" {
		t.Error(err3)
		return
	}

}
func TestCache_GetWithFunc(t *testing.T) {
	c := NewCache(5)

	v, err := c.GetWithFunc("key", func() (Value, error) {
		fmt.Println("called")
		return String("hello"), nil
	})
	if err != nil {
		t.Error(err)
		return
	}
	if v.(String) != "hello" {
		t.Error("error")
		return
	}
	val, ok := c.Get("key")
	if !ok {
		t.Error("error")
		return
	}
	if val.(String) != "hello" {
		t.Error("error")
		return
	}
	if c.Used() != 5 {
		t.Error("error")
		return
	}
	fmt.Println(c.Used())
}

func TestCache_SetWithExpire(t *testing.T) {

	c := NewCache(5)
	err := c.SetWithExpire("key", &Item{Val: String("hello"), ExpireTime: 10})
	if err != nil {
		t.Error(err)
		return
	}
	if c.Used() != 5 {
		t.Error("error")
		return
	}
	time.Sleep(5 * time.Second)
	fmt.Println(c.Get("key"))
	fmt.Println(c.Used())
}
func TestCache_Singleflight(t *testing.T) {
	c := NewCache(5)
	var wg sync.WaitGroup
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			v, err := c.GetWithFunc("key", func() (Value, error) {
				fmt.Println("called") // called 1 time because of singleflight
				time.Sleep(time.Millisecond * 100)
				return String("hello"), nil
			})
			if err != nil {
				t.Error(err)
				return
			}
			fmt.Println(v)
		}()
	}
	wg.Wait()

}
