## gpool

一个强大的协程池

## feature

-[x] 动态扩容
-[x] goroutine的复用
-[x] 自定义panic处理任务函数的panic
-[x] 并发安全

## design

-[x] 使用`sync.Pool`复用大量`worker`和`task`对象，减轻GC的压力
-[x] `sync.Mutex`互斥锁用来保证任务链表代码片段的并发操作安全
-[x] 项目里大量使用原子操作来对全局唯一的`pool`变量做读写等操作，获得更好的并发性能

## quick start

one way:

```go

```

another way:

```go

```