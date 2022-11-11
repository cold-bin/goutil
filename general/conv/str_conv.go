// @author cold bin
// @date 2022/11/11

package conv

import (
	"strconv"
	"time"
)

// StringPtrUint 忽略错误string转uint
func StringPtrUint(x string) (y uint) {
	re, _ := strconv.ParseUint(x, 10, 64)
	y = uint(re)
	return
}

// UintPtrString uint转string
func UintPtrString(x uint) (y string) {
	y = strconv.FormatUint(uint64(x), 10)
	return
}

// StringPtrInt 忽略错误string转int
func StringPtrInt(x string) (i int) {
	i, _ = strconv.Atoi(x)
	return
}

// IntPtrString int转string
func IntPtrString(i int) (x string) {
	x = strconv.Itoa(i)
	return x
}

// Int64PtrString int64转string
func Int64PtrString(x int64) (y string) {
	y = strconv.FormatInt(x, 10)
	return
}

// StringPtrInt64 忽略错误string转int64
func StringPtrInt64(x string) (y int64) {
	y, _ = strconv.ParseInt(x, 10, 64)
	return
}

// UnixTimePtrTime UnixTime转时间
func UnixTimePtrTime(x int64) time.Time {
	return time.Unix(x, 0)
}

// UnixTimePtrString UnixTime转string
func UnixTimePtrString(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	return t.Format("2006-01-02 15:04:05")
}
