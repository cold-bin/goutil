// @author cold bin
// @date 2022/10/24

package conv

import (
	"unsafe"
)

// QuickS2B 快速地将string类型转化为[]byte类型
func QuickS2B(str string) []byte {
	base := *(*[2]uintptr)(unsafe.Pointer(&str))
	return *(*[]byte)(unsafe.Pointer(&[3]uintptr{base[0], base[1], base[1]}))
}

// QuickB2S 快速地将[]byte转化为string类型
func QuickB2S(bs []byte) string {
	base := (*[3]uintptr)(unsafe.Pointer(&bs))
	return *(*string)(unsafe.Pointer(&[2]uintptr{base[0], base[1]}))
}
