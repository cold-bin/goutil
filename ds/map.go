package ds

import "cmp"

type Map[kType cmp.Ordered, vType any] interface {
	Get(key kType) vType
	Set(kType kType, val vType)
	Remove(key kType) vType
	Foreach(op func(key kType, val vType))
}
