// @author cold bin
// @date 2022/10/24

package dcopy

import (
	"reflect"
	"testing"
)

func TestCopy(t *testing.T) {

	type Struct struct {
		Int int
		Map map[int]string
	}

	s := Struct{
		Int: 1,
		Map: map[int]string{1: "1"},
	}
	newS := DeepCopy(s)
	s.Map[1] = "2"

	if reflect.DeepEqual(s, newS) {
		t.Errorf("TestCopy()=相等, want=不相等")
	}
}
