// @author cold bin
// @date 2022/10/22

package _err

import (
	"fmt"
	"runtime"
)

// Err goutil 内部错误信息，针对nil错误使用 Err{} 空对象来表示
type Err struct {
	Where string
	Err   error
}

// 调用此方法前先使用 WrapErr 函数将 error 类型转变为 goutil 自身定义的错误类型
func (e Err) Error() string {
	return fmt.Sprintf("[goutil] %s发生error: %s", e.Where, e.Err)
}

func (e Err) IsNil() bool {
	return e.Err == nil
}

//// ErrNil goutil 提供一个类似 nil 的错误空值
//func ErrNil() Err {
//	return Err{}
//}

func WrapErr(err error) Err {
	// 错误发生的函数位置
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return Err{}
	}

	return Err{
		Where: fmt.Sprintf("%s 第%d行", file, line),
		Err:   err,
	}
}
