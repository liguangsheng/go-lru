package benchmark

import (
	"github.com/liguangsheng/go-lru"
	"testing"
)

func Benchmark_golru_Set(b *testing.B) {
	c := lru.New().Cap(size).Safe(true).Build()
	for i := 0; i < b.N; i++ {
		c.Set(key(i), value(i))
	}

}

func Benchmark_golru_Get(b *testing.B) {
	c := lru.New().Cap(size).Safe(true).Build()
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			c.Set(key(i), value(i))
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value, ok := c.Get(string(i))
		if ok {
			_ = value
		}
	}
}

func Benchmark_golru_UnsafeSet(b *testing.B) {
	c := lru.New().Cap(size).Safe(false).Build()
	for i := 0; i < b.N; i++ {
		c.Set(key(i), value(i))
	}

}

func Benchmark_golru_UnsafeGet(b *testing.B) {
	c := lru.New().Cap(size).Safe(false).Build()
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			c.Set(key(i), value(i))
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value, ok := c.Get(string(i))
		if ok {
			_ = value
		}
	}
}
