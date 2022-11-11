// @author cold bin
// @date 2022/11/11

package gpool

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestPoolPanic(t *testing.T) {
	p := NewPool(100, 1, nil)
	p.Go(func() {
		panic("test panic")
	})
}

// TestPool 测试协程复用
func TestPool(t *testing.T) {
	p := NewPool(100, 1, nil)
	var n int32
	var wg sync.WaitGroup
	for i := 0; i < 200000; i++ {
		wg.Add(1)
		p.Go(func() {
			defer wg.Done()
			atomic.AddInt32(&n, 1)
		})
	}
	wg.Wait()
	if n != 200000 {
		t.Error(n)
	}
}
