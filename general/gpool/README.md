## gpool

一个强大的协程池

## feature

- [x] 动态扩容
- [x] goroutine的复用
- [x] 自定义panic处理任务函数的panic
- [x] 并发安全

## design

- [x] 使用`sync.Pool`复用大量`worker`和`task`对象，减轻GC的压力
- [x] `sync.Mutex`互斥锁用来保证任务链表代码片段的并发操作安全
- [x] 项目里大量使用原子操作来对全局唯一的`pool`变量做读写等操作，获得更好的并发性能
- [x] goroutine复用策略：
    ```txt
    判断是否需要扩容
	    1. 任务函数数量大于 p.fNumPerWorker；
	    2. 且当前正在跑的 worker 数目小于 p.cap
        3. 或者没有 worker
    ```

## quick start

### one way

类似标准库`go func(){}()`的使用方式，只需要调标准包即可。
值得注意的是：`gpool`只提供了默认的`defaultPool`，
默认系统并发协程数最大是10000

```go
// @author cold bin
// @date 2022/10/24

package main

import (
  "fmt"
  "goutil/general/gpool"
  "sync"
  "sync/atomic"
)

func main() {
  var n int32
  var wg sync.WaitGroup
  for i := 0; i < 1100000; i++ {
    wg.Add(1)
    gpool.Go(func() {
      defer wg.Done()
      //n++
      atomic.AddInt32(&n, 1)
    })
  }
  wg.Wait()
  fmt.Println(n)
}
```

### another way

`gpool`也支持自定义自己的负载策略：支持定制当前系统应该开的协程最大数、
也支持自定义`panic`来处理自己的函数错误、还支持定制每个协程应该完成的任务数（最好为1）
。

```go
// @author cold bin
// @date 2022/10/24

package main

import (
  "fmt"
  "goutil/general/gpool"
  "sync"
  "sync/atomic"
)

func main() {
  p := gpool.NewPool(10, 2, nil)
  var n int32
  var wg sync.WaitGroup
  for i := 0; i < 1100000; i++ {
    wg.Add(1)
    p.Go(func() {
      defer wg.Done()
      atomic.AddInt32(&n, 1)
    })
  }
  wg.Wait()
  fmt.Println(n)
}
```

## attention

`gpool`内部没有使用反射，所以对函数只支持简单的无参也无返回值的函数。
所以如果需要传入参数或传出参数时，需要自己照顾参数的并法安全。
`gpool`只提供可复用的协程池的环境

## reference

来自字节开源的go工具
https://github.com/bytedance/gopkg