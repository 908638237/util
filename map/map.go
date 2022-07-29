package _map

import (
	"github.com/908638237/util/common"
	"math/rand"
	"strconv"
	"time"
)

func IntToStrMap(data map[string]int) map[string]string {
	rs := make(map[string]string)
	for k, v := range data {
		rs[k] = strconv.Itoa(v)
	}
	return rs
}

func StrToIntMap(data map[string]string) map[string]int {
	rs := make(map[string]int)
	for k, v := range data {
		intVal, _ := strconv.Atoi(v)
		rs[k] = intVal
	}
	return rs
}

func MergeMap[K comparable, V comparable](maps ...map[K]V) map[K]V {
	data := make(map[K]V)
	if len(maps) < 1 {
		return nil
	}
	if len(maps) < 2 {
		return maps[0]
	}
	for _, mapVal := range maps {
		for key, val := range mapVal {
			data[key] = val
		}
	}
	return data
}

// Shuffle 打乱slice顺序
func Shuffle(slice []any) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

func SliceIn(arr []string, val string) bool {
	if len(arr) < 1 {
		return false
	}
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func InArray[K comparable, D common.DataType](val D, arr map[K]D) bool {
	if len(arr) < 1 {
		return false
	}
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
