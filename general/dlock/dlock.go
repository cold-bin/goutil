// @author cold bin
// @date 2023/1/11

package dlock

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"runtime"
	"strconv"
	"time"
)

// DLock 分布式锁接口约定。注意：使用前必须传入有效的redis句柄
type DLock interface {
	// TryLock 会试图给当前线程上锁。需要提供一个键名和锁超时时间
	TryLock(timeout time.Duration) (bool, error)
	// Unlock 会原子的保证对应进程锁的释放
	Unlock() (bool, error)
	// ExpectRedis 分布式锁依赖redis实现，所以需要在此处传入redis句柄
	ExpectRedis(redisC *redis.Client) error
}

type dLock struct {
	rdb      *redis.Client
	key      string
	idPrefix string
}

func newdLock() DLock {
	return &dLock{
		key:      getRandStr(),
		idPrefix: getRandStr(),
	}
}

func (d *dLock) ExpectRedis(redisC *redis.Client) error {
	// 探测redis是否可用
	if err := redisC.Ping(context.Background()).Err(); err != nil {
		return err
	}
	d.rdb = redisC
	return nil
}

func (d *dLock) TryLock(timeout time.Duration) (bool, error) {
	if d.rdb == nil {
		panic("redis client is nil")
	}
	// 获取调用函数的线程标识，拼接id保证随机性
	val := fmt.Sprintf("%s:%s", d.idPrefix, getGID())
	// 尝试获取锁
	return d.rdb.SetNX(context.Background(), d.key, val, timeout).Result()
}

func (d *dLock) Unlock() (bool, error) {
	if d.rdb == nil {
		panic("redis client is nil")
	}
	// 使用lua保证redis多条命令执行的原子性
	luaScript := `
		-- 获取锁中的标示，判断是否与当前线程标示一致
		if (redis.call('GET', KEYS[1]) == ARGV[1]) then
  		-- 一致，则删除锁
  		return redis.call('DEL', KEYS[1])
		end
		-- 不一致，则直接返回
		return 0`

	result, err := d.rdb.Eval(
		context.Background(),
		luaScript,
		[]string{d.key},
		[]string{fmt.Sprintf("%s:%d", d.idPrefix, getGID())},
	).Result()

	if err != nil {
		return false, err
	}

	return result.(int) == 1, nil
}

// 获取调用者的协程ID，用来标识协程
func getGID() uint64 {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func getRandStr() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	return uuid.String()
}
