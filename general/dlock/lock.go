// @author cold bin
// @date 2023/1/11

package dlock

import (
	"github.com/go-redis/redis/v8"
	"time"
)

var mdlock DLock

func init() {
	mdlock = newdLock()
}

// TryLock 必须先使用 SetRedisC 函数置入redis客户端句柄，否则会 panic
func TryLock(timeout time.Duration) (bool, error) {
	return mdlock.TryLock(timeout)
}

// Unlock 必须先使用 SetRedisC 函数置入redis客户端句柄，否则会 panic。释放锁是原子的
func Unlock() (bool, error) {
	return mdlock.Unlock()
}

// SetRedisC 置入有效redis客户端句柄，不可用的redis句柄会panic
func SetRedisC(redisC *redis.Client) {
	if err := mdlock.ExpectRedis(redisC); err != nil {
		panic(err)
	}
}
