// @author cold bin
// @date 2022/11/9

package gpool

import (
	"context"
	"fmt"
	"goutil/internal/_log"
	"log"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

type Pool interface {
	SetCap(cap int32)                    // 动态扩容协程池的容量
	Go(f func())                         // 执行f
	CtxGo(ctx context.Context, f func()) // 传递上下文的执行函数
	WorkerCount() int32                  // 获取 worker 数
}

var taskPool = sync.Pool{New: func() any { return &task{} }}

type task struct {
	ctx  context.Context
	f    func()
	next *task
}

func (t *task) zero() {
	t.ctx = nil
	t.f = nil
	t.next = nil
}

// Recycle 回收到 taskPool
func (t *task) Recycle() {
	t.zero()
	taskPool.Put(t)
}

type PanicFunc func(any)

type pool struct {
	cap           int32 // 协程池的最大容量，即允许跑的最大协程数
	fNumPerWorker int32 // 每个 worker 应该执行 task 个数，默认为1

	taskHead  *task      // 任务链表的头
	taskTail  *task      // 任务链表的尾，方便快速添加任务
	taskLock  sync.Mutex // 互斥锁保证链表操作的并发安全
	taskCount int32      // 动态记录当前剩余任务的个数

	workerCount int32 // 动态记录当前活跃的goroutine数目

	panicHandler PanicFunc // 任务函数自定义的panic处理
}

var DefaultPanicFunc = func(a any) {
	log.Println(fmt.Sprintf("goutil: panic in gpool: %v: %s", a, debug.Stack()))
}

// SetCap 原子载入
func (p *pool) SetCap(cap int32) {
	atomic.StoreInt32(&p.cap, cap)
}

func (p *pool) Go(f func()) {
	p.CtxGo(context.Background(), f)
}

func (p *pool) CtxGo(ctx context.Context, f func()) {
	t := taskPool.Get().(*task)
	t.ctx = ctx
	t.f = f

	// 这里会有多个协程对 task 任务链表并发操作
	p.taskLock.Lock()
	if p.taskHead == nil {
		p.taskHead = t
		p.taskTail = t
	} else {
		p.taskTail.next = t
		p.taskTail = t
	}
	p.taskLock.Unlock()

	// 更新状态: 任务+1
	p.incTaskCount()

	// 判断是否需要扩容
	//  1. 任务函数数量大于 p.fNumPerWorker；
	//	2. 且当前正在跑的 worker 数目小于 p.cap
	// 	3. 或者没有 worker
	if p.TaskCount() >= p.fNumPerWorker && p.WorkerCount() < p.Cap() || p.WorkerCount() == 0 {
		p.incWorkerCount()
		w := workerPool.Get().(*worker)
		// 所有的worker都指向一个pool
		w.pool = p
		w.run()
	} else {
		// 超负荷
		_log.GetLog(_log.LogWarnL).Panic("依据负载策略，已达超负荷运行边界")
	}
}

func (p *pool) Cap() int32 {
	return atomic.LoadInt32(&p.cap)
}

func (p *pool) TaskCount() int32 {
	return atomic.LoadInt32(&p.taskCount)
}

func (p *pool) incTaskCount() int32 {
	return atomic.AddInt32(&p.taskCount, 1)
}

func (p *pool) decTaskCount() int32 {
	return atomic.AddInt32(&p.taskCount, -1)
}

// WorkerCount 原子加载 workerCount
func (p *pool) WorkerCount() int32 {
	return atomic.LoadInt32(&p.workerCount)
}

func (p *pool) incWorkerCount() int32 {
	return atomic.AddInt32(&p.workerCount, 1)
}

func (p *pool) decWorkerCount() int32 {
	return atomic.AddInt32(&p.workerCount, -1)
}

// NewPool 配置协程池参数
//
//	cap: 池的容量，大于等于10000
//	fNumPerWorker: 每个 worker 应该执行 task 个数，推荐配为1。当大于2或小于1时，直接置为1
//	panicHandler: 如果为nil将会使用默认的 DefaultPanicFunc 处理
func NewPool(cap, fNumPerWorker int32, panicHandler PanicFunc) Pool {
	if cap <= 0 {
		cap = 10000
	}

	if fNumPerWorker < 1 {
		fNumPerWorker = 1
	}

	if panicHandler == nil {
		panicHandler = DefaultPanicFunc
	}

	return &pool{
		cap:           cap,
		fNumPerWorker: fNumPerWorker,
		panicHandler:  panicHandler,
	}
}
