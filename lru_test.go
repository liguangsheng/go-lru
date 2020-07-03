package lru

import (
	"sync"
	"testing"
	"time"
)

func equal(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if expected != actual {
		t.Error("not equal")
		return false
	}
	return true
}

func Test_LRU(t *testing.T) {
	c := New().Build()
	c.Set("hello", "world")
	ret, ok := c.Get("hello")
	equal(t, true, ok)
	equal(t, "world", ret)
	c.Remove("hello")
	_, ok = c.Get("hello")
	equal(t, false, ok)

	for i := 0; i < 100000; i++ {
		c.Set(string(i), string(i))
	}
}

func Test_LRUOverflow(t *testing.T) {
	c := New().Cap(3).Build()
	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)
	c.Set("d", 4)
	_, ok := c.Get("a")
	equal(t, false, ok)
	c.Get("b")
	c.Set("c", 1)
	_, ok = c.Get("c")
	equal(t, true, ok)
}

func Test_LRUExpire(t *testing.T) {
	c := New().Build()
	c.SetWithExpire("hello", "world", time.Now().Add(time.Second))
	time.Sleep(time.Second * 2)
	_, ok := c.Get("hello")
	equal(t, false, ok)
}

func Test_LRURace(t *testing.T) {
	c := New().Build()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			c.Set("hello", "world")
		}()

		go func() {
			defer wg.Done()
			c.Get("hello")
		}()
	}
	wg.Wait()
}
