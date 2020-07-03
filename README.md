# go-lru

Go lru cache implement, super faster, less alloc.

## install

```
go get -u github.com/liguangsheng/go-lru
```

## example
```
c := lru.New().Cap(1000).Safe(true).Build()
c.Set("hello", "world")
c.SetWithExpire("hello", "world", time.Now.Add(time.Minute))
c.SetWithDuration("hello", "world", time.Minute)

v, ok := c.Get("hello")
if ok {
	fmt.Println(v) // world
}
```

## benchmark
```
goos: darwin
goarch: amd64
pkg: github.com/liguangsheng/go-lru/_benchmark
Benchmark_golru_Set-8         	12970200	        84.0 ns/op	      15 B/op	       1 allocs/op
Benchmark_golru_Get-8         	26832456	        43.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_golru_UnsafeSet-8   	13437583	        77.6 ns/op	      15 B/op	       1 allocs/op
Benchmark_golru_UnsafeGet-8   	32874222	        35.6 ns/op	       0 B/op	       0 allocs/op
```
