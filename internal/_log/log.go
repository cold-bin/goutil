// @author cold bin
// @date 2022/10/24

// Package _log
// 主要提供给goutil内部的日志记录
package _log

import (
	"log"
	"os"
	"sync"
)

const (
	LogDebugL = 1 + iota
	LogInfoL
	LogWarnL
	LogErrL
)

// goutil 内部使用的logger，分级别
var (
	logErrPool = sync.Pool{New: func() any {
		return log.New(os.Stderr, "[goutil err log] ", log.LstdFlags|log.Llongfile)
	}}

	logWarnPool = sync.Pool{New: func() any {
		return log.New(os.Stderr, "[goutil warn log] ", log.LstdFlags|log.Llongfile)
	}}

	logInfoPool = sync.Pool{New: func() any {
		return log.New(os.Stderr, "[goutil info log] ", log.LstdFlags|log.Llongfile)
	}}

	logDebugPool = sync.Pool{New: func() any {
		return log.New(os.Stderr, "[goutil debug log] ", log.LstdFlags|log.Llongfile)
	}}
)

// GetLog 获取指定level的logger
func GetLog(level int) *log.Logger {
	switch level {
	case LogDebugL:
		if logger := logDebugPool.Get().(*log.Logger); logger != nil {
			return logger
		}
		return log.New(os.Stderr, "[goutil debug log] ", log.LstdFlags|log.Llongfile)

	case LogInfoL:
		if logger := logInfoPool.Get().(*log.Logger); logger != nil {
			return logger
		}
		return log.New(os.Stderr, "[goutil info log] ", log.LstdFlags|log.Llongfile)

	case LogWarnL:
		if logger := logWarnPool.Get().(*log.Logger); logger != nil {
			return logger
		}
		return log.New(os.Stderr, "[goutil warn log] ", log.LstdFlags|log.Llongfile)

	case LogErrL:
		if logger := logErrPool.Get().(*log.Logger); logger != nil {
			return logger
		}
		return log.New(os.Stderr, "[goutil err log] ", log.LstdFlags|log.Llongfile)
	default:
		return log.New(os.Stderr, "[goutil log] ", log.LstdFlags|log.Llongfile)
	}
}
