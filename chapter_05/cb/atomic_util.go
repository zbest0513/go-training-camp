package cb

import "sync/atomic"

//减法
func less(a *int64, n ...int64) int64 {
	if len(n) == 0 {
		i := int64(-1)
		return atomic.AddInt64(a, i)
	}
	var result int64
	for _, i := range n {
		var x = int64(0 - i)
		result = atomic.AddInt64(a, x)
	}
	return result
}

//加法
func add(a *int64, n ...int64) int64 {
	if len(n) == 0 {
		return atomic.AddInt64(a, int64(1))
	}
	var result int64
	for _, i := range n {
		result = atomic.AddInt64(a, i)
	}
	return result
}

//查询
func load(a *int64) int64 {
	return atomic.LoadInt64(a)
}
