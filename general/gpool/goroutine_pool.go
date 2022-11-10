// @author cold bin
// @date 2022/10/24

// Package gpool
// 本包提供协程的池化
package gpool

import (
	"context"
	"log"
)

var defaultPool Pool

func init() {
	defaultPool = NewPool(10000, 1, DefaultPanicFunc)
	log.SetFlags(log.Llongfile)
}

// SetCap 动态扩容
func SetCap(cap int32) {
	defaultPool.SetCap(cap)
}

func Go(f func()) {
	defaultPool.Go(f)
}

func CtxGo(ctx context.Context, f func()) {
	defaultPool.CtxGo(ctx, f)
}

func WorkerCount() int32 {
	return defaultPool.WorkerCount()
}
