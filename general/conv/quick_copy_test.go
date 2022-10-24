// @author cold bin
// @date 2022/10/24

package conv

import (
	"reflect"
	"testing"
)

func TestQuickS2B(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"right 1: ", args{str: "test % .; 你好"}, []byte("test % .; 你好")},
		{"right 2: ", args{str: "测试快速转化为字节切片的正确性"}, []byte("测试快速转化为字节切片的正确性")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickS2B(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickS2B() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickB2S(t *testing.T) {
	type args struct {
		bs []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"right 1: ", args{bs: []byte("test % .; 你好")}, "test % .; 你好"},
		{"right 2: ", args{bs: []byte("测试快速转化为字节切片的正确性")}, "测试快速转化为字节切片的正确性"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickB2S(tt.args.bs); got != tt.want {
				t.Errorf("QuickB2S() = %v, want %v", got, tt.want)
			}
		})
	}
}
