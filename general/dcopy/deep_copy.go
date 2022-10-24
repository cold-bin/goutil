// @author cold bin
// @date 2022/10/24

package dcopy

import (
	"reflect"
	"time"
)

func DeepCopy(src any) any {
	if src == nil {
		return nil
	}

	original := reflect.ValueOf(src)
	cpy := reflect.New(original.Type()).Elem()
	copyRecursive(original, cpy)

	return cpy.Interface()
}

// 递归深拷贝，直至拷贝到基础类型为止，基础类型的赋值一定是
func copyRecursive(src, dst reflect.Value) {

	switch src.Kind() {
	case reflect.Ptr:
		// 这里主要是取出指针的实际数据，然后根据数据在反射创建一个新的指针变量，
		// 更改这个指针变量的元素值，从而实现指针的深拷贝
		originalValue := src.Elem()

		if !originalValue.IsValid() {
			return
		}

		dst.Set(reflect.New(originalValue.Type()))
		copyRecursive(originalValue, dst.Elem())

	case reflect.Interface:
		// 先判断指针是否为nil，为nil就没必要继续了，直接返回到上层递归
		if src.IsNil() {
			return
		}
		originalValue := src.Elem()
		copyValue := reflect.New(originalValue.Type()).Elem()
		copyRecursive(originalValue, copyValue)

		dst.Set(copyValue)

	case reflect.Struct:
		// 使用时间的程序通常应该将它们作为值，而不是指针来存储和传递。不能深度拷贝时间，貌似拷贝出来的时间都是utc
		t, ok := src.Interface().(time.Time)
		if ok {
			dst.Set(reflect.ValueOf(t))
			return
		}
		for i := 0; i < src.NumField(); i++ {
			// TODO 跳过未导出字段，后期迭代考虑迁移到可配置的选项里
			if src.Type().Field(i).PkgPath != "" {
				continue
			}
			// 递归深拷贝直到未配置类型的非基础类型
			copyRecursive(src.Field(i), dst.Field(i))
		}

	case reflect.Slice:
		if src.IsNil() {
			return
		}
		// 反射创建切片，从而初始化
		dst.Set(reflect.MakeSlice(src.Type(), src.Len(), src.Cap()))
		// 递归深拷贝切片的每个元素
		for i := 0; i < src.Len(); i++ {
			copyRecursive(src.Index(i), dst.Index(i))
		}
	case reflect.Array:
		dst.Set(reflect.New(reflect.ArrayOf(src.Len(), src.Type())).Elem())
		for i := 0; i < src.Len(); i++ {
			copyRecursive(src.Index(i), dst.Index(i))
		}
	case reflect.Map:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.MakeMap(src.Type()))
		for _, originalKey := range src.MapKeys() {
			// 取值
			originalValue := src.MapIndex(originalKey)
			// 复制
			copyValue := reflect.New(originalValue.Type()).Elem()
			// 首先对map的值递归，直到基础类型
			copyRecursive(originalValue, copyValue)
			// 然后再递归map的键递归，直到基础类型
			//copyKey := Copy(key.Interface())
			copyKey := reflect.New(originalKey.Type()).Elem()
			copyRecursive(originalKey, copyKey)
			dst.SetMapIndex(copyKey, copyValue)
		}
	//	TODO: 函数貌似深拷贝也没啥用
	//case reflect.Func:
	//	if src.IsNil() {
	//		return
	//	}
	//	reflect.MakeFunc(src.Type(), func(in []reflect.Value) (out []reflect.Value) {
	//
	//	})
	// TODO: 反射提供的api不足以深拷贝一个 chan 好像，不知道为啥... 那暂时默认使用浅拷贝吧
	//case reflect.Chan:
	//if src.IsNil() {
	//	return
	//}
	//newChanV := reflect.MakeChan(src.Type(), src.Cap())
	//dst.Set(newChanV)
	default:
		// 递归结束条件，直到未配置的类型才会被直接赋值
		dst.Set(src)
	}
}
