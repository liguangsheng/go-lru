package benchmark
//
//import (
//	"github.com/tidwall/tinylru"
//	"testing"
//)
//
//func Benchmark_tinylru_lru_Set(b *testing.B) {
//	c := tinylru.LRU{}
//	c.Resize(size)
//	for i := 0; i < b.N; i++ {
//		c.Set(key(i), value(i))
//	}
//}
//
//func Benchmark_tinylru_lru_Get(b *testing.B) {
//	c := tinylru.LRU{}
//	c.Resize(size)
//	for i := 0; i < b.N; i++ {
//		if i%2 == 0 {
//			c.Set(key(i), value(i))
//		}
//	}
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		c.Get(key(i))
//	}
//}
