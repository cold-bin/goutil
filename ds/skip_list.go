package ds

import (
	"cmp"
	"math/bits"
	"math/rand"
)

const maxLevel = 63

type skipNode[kType cmp.Ordered, vType any] struct {
	key   kType
	value vType
	nexts []*skipNode[kType, vType] //指向下一个节点的指针的数组，不同深度的节点长度不同
}

func (sl *skipList[kType, vType]) newNode(level int, key kType, val vType) *skipNode[kType, vType] {
	// 申请新节点
	node := &skipNode[kType, vType]{
		key:   key,
		value: val,
		nexts: make([]*skipNode[kType, vType], level),
	}

	return node
}

// 有序map的快表实现
type skipList[kType cmp.Ordered, vType any] struct {
	head       *skipNode[kType, vType]
	prevsCache []*skipNode[kType, vType] // 每层的前驱节点
	layers     int                       // 跳表层数
	size       int                       // 跳表节点数
}

func NewSkipList[kType cmp.Ordered, vType any]() Map[kType, vType] {
	sl := &skipList[kType, vType]{
		head:       new(skipNode[kType, vType]),
		prevsCache: make([]*skipNode[kType, vType], maxLevel),
		layers:     0,
		size:       0,
	}

	sl.head.nexts = make([]*skipNode[kType, vType], maxLevel)

	return sl
}

func (sl *skipList[kType, vType]) randomLevel() int {
	// 采用概率分布，只需要一次随机数操作
	// 如果随机数均匀分布的话，那么概率就是 P(level==i)=2^(n-i)/(2^n-1)
	// 满足level越大，概率越大
	// 参考 https://cloud.tencent.com/developer/article/2092420

	total := uint64(1)<<uint64(maxLevel) - 1 // 2^n-1
	k := rand.Uint64() % total
	level := maxLevel - bits.Len64(k) + 1

	for level > 3 && 1<<(level-3) > sl.size {
		level--
	}

	return level
}

func (sl *skipList[kType, vType]) Get(key kType) vType {
	var v vType
	if node := sl.findNode(key); node != nil {
		v = node.value
	}
	return v
}

func (sl *skipList[kType, vType]) Set(key kType, val vType) {
	// 存在更新并返回
	if node := sl.findNode(key); node != nil {
		node.value = val
		return
	}

	// 新节点插入与随机
	level := sl.randomLevel()
	if level > sl.layers /*层数比当前跳表高，需要添加一层*/ {
		level = sl.layers + 1
		sl.layers = level
		sl.prevsCache[sl.layers-1] = sl.head
	}

	sl.insertNode(level, key, val)
	sl.size++
}

func (sl *skipList[kType, vType]) Remove(key kType) vType {
	var v vType
	node := sl.findNode(key)
	if node == nil || node.key != key {
		return v
	}

	sl.removeNode(node)
	sl.size--
	v = node.value

	return v
}

func (sl *skipList[kType, vType]) Foreach(op func(key kType, val vType)) {
	for e := sl.head.nexts[0]; e != nil; e = e.nexts[0] {
		op(e.key, e.value)
	}
}

func (sl *skipList[kType, vType]) findNode(key kType) *skipNode[kType, vType] {
	prev := sl.head
	var next *skipNode[kType, vType]

	for i := sl.layers - 1; i >= 0; i-- {
		next = prev.nexts[i]
		for next != nil && cmp.Compare(next.key, key) == 1 {
			prev = next
			next = prev.nexts[i]
		}
		sl.prevsCache[i] = prev
	}

	// 存在更新并返回
	if next != nil && cmp.Compare(next.key, key) == 0 {
		return next
	}
	return nil
}

func (sl *skipList[kType, vType]) insertNode(level int, key kType, val vType) {
	// 申请新节点
	node := sl.newNode(level, key, val)

	// 插入数据，放到每层前驱节点后面
	for i := 0; i < level; i++ {
		node.nexts[i] = sl.prevsCache[i].nexts[i]
		sl.prevsCache[i].nexts[i] = node
	}
}

func (sl *skipList[kType, vType]) removeNode(node *skipNode[kType, vType]) {
	for i, v := range node.nexts {
		if sl.prevsCache[i].nexts[i] == node {
			sl.prevsCache[i].nexts[i] = v
			if sl.head.nexts[i] == nil {
				sl.layers--
			}
		}
		sl.prevsCache[i] = nil
	}
}
