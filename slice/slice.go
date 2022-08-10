package slice

import (
	"errors"
	"github.com/908638237/util/common"
	"math/rand"
	"reflect"
	"time"
)

// ShuffleSlice 打乱slice
func ShuffleSlice(src interface{}) (err error) {
	srcValue := reflect.ValueOf(src)
	if err = checkInterfaceVal(srcValue, reflect.Ptr, reflect.Slice); err != nil {
		return
	}
	srcValue = srcValue.Elem()
	l := srcValue.Len()
	if l <= 0 {
		return
	}
	rand.Seed(time.Now().UnixNano() + 11)
	temp := make([]reflect.Value, 0)
	for i := 0; i < l; i++ {
		temp = append(temp, srcValue.Index(i))
	}
	for i := len(temp) - 1; i > 0; i-- {
		j := rand.Intn(i)
		temp[i], temp[j] = temp[j], temp[i]
	}
	srcValue.SetLen(0)
	srcValue.SetCap(0)
	arr := reflect.Append(srcValue, temp...)
	srcValue.Set(arr)
	return
}

func SliceToMap(src, desPointer interface{}, fn func(elem interface{}) (k, v interface{})) (err error) {
	srcValue := reflect.ValueOf(src)
	desValue := reflect.ValueOf(desPointer)
	if err = checkInterfaceVal(srcValue, reflect.Slice); err != nil {
		return
	}
	if err = checkInterfaceVal(desValue, reflect.Ptr, reflect.Map); err != nil {
		return
	}
	l := srcValue.Len()
	desValue = desValue.Elem()
	if l <= 0 {
		return errors.New("src长度为空")
	}
	for i := 0; i < l; i++ {
		key, val := fn(srcValue.Index(i).Interface())
		if i == 0 && reflect.ValueOf(val).Type() != desValue.Type().Elem() {
			return errors.New("value类型不正确,限制: " + desValue.Type().Elem().String())
		}
		desValue.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
	}
	return
}

func checkInterfaceVal(srcValue reflect.Value, kind ...reflect.Kind) error {
	srcKind := srcValue.Kind()
	ptrString := ""
	startIndex := 0
	if len(kind) > 0 && kind[0] == reflect.Ptr {
		if srcKind != kind[0] {
			return errors.New("参数必须是ptr类型, 当前: " + srcKind.String())
		}
		srcKind = srcValue.Elem().Kind()
		ptrString = ".ptr"
		startIndex = 1
	}
	for i := startIndex; i < len(kind); i++ {
		if srcKind != kind[i] {
			return errors.New("参数必须是" + kind[i].String() + ptrString + "类型,当前:" + srcKind.String())
		}
	}
	return nil
}

// RemoveRepeatElementSlice 切片去重
func RemoveRepeatElementSlice[D common.DataType](list []D) []D {
	temp := make(map[D]bool)
	index := 0
	for _, v := range list {
		_, ok := temp[v]
		if ok {
			list = append(list[:index], list[index+1:]...)
			index--
		} else {
			temp[v] = true
		}
		index++
	}
	return list
}

// ContainsSlice 判断slice是否包含某值
func ContainsSlice(src interface{}, val interface{}) bool {
	srcValue := reflect.ValueOf(src)
	value := reflect.ValueOf(val)
	if srcValue.Kind() != reflect.Slice || srcValue.Len() < 1 {
		return false
	}
	for i := 0; i < srcValue.Len(); i++ {
		sv := srcValue.Index(i)
		if value.Kind() == sv.Kind() && value.Interface() == sv.Interface() {
			return true
		}
	}
	return false
}
