// @author cold bin
// @date 2022/11/9

package gpool

import (
	"sync"
)

var workerPool = sync.Pool{New: func() any { return &worker{} }}

// 执行者，绑定一个goroutine运行
// 值得注意的是：多个 worker 也指向统一个 pool
type worker struct {
	pool *pool
}

// 真正开协程的地方，去任务链表里抢任务来做
func (w *worker) run() {
	go func() {
		var t *task

		w.pool.taskLock.Lock()
		// 拿任务
		if w.pool.taskHead != nil {
			t = w.pool.taskHead
			w.pool.taskHead = w.pool.taskHead.next

			w.pool.decTaskCount()
		}

		if t == nil {
			w.close()
			w.pool.taskLock.Unlock()
			w.Recycle()
			return
		}

		w.pool.taskLock.Unlock()

		// 确保捕获的panic一定是任务函数
		func(t *task) {
			defer func() {
				if err := recover(); err != nil {
					w.pool.panicHandler(err)
				}
			}()
			t.f()
		}(t)

		// 回收task
		t.Recycle()
	}()
}

func (w *worker) close() {
	w.pool.decWorkerCount()
}

func (w *worker) zero() {
	w.pool = nil
}

func (w *worker) Recycle() {
	w.zero()
	workerPool.Put(w)
}
