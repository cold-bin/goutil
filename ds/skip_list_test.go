package ds

import (
	"testing"
)

func TestAll(t *testing.T) {
	list := NewSkipList[int, string]()

	list.Set(1, "1")
	list.Set(2, "2")
	list.Set(3, "3")
	list.Set(4, "4")

	assert(list.Get(3) == "3", t)
	assert(list.Get(3) == "3", t)
	assert(list.Get(4) == "4", t)
	assert(list.Get(4) == "4", t)
	assert(list.Get(5) != "5", t)
	assert(list.Get(5) != "5", t)

	list.Set(1, "10")
	assert(list.Get(1) == "10", t)
	assert(list.Get(1) != "1", t)

	list.Remove(1)
	assert(list.Get(1) == "", t)
}

func assert(ok bool, t *testing.T) {
	if !ok {
		t.Error("wrong!!!")
	}
}
