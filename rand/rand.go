package rand

import (
	"errors"
	"fmt"
	"github.com/908638237/util"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

func Max[D util.DataType](a, b D) D {
	if a > b {
		return a
	}
	return b
}

func Min[D util.DataType](a, b D) D {
	if a > b {
		return b
	}
	return a
}

func Dice(arr map[int]int) (key int, err error) {
	start := int32(0)
	max := int32(1)
	key = -1
	data := make(map[int][2]int32)
	for k, v := range arr {
		end := start + int32(v)
		data[k] = [2]int32{start, int32(end)}
		start = end + 1
		max = end
	}
	rand := RandInterval(0, int32(math.Max(float64(max), 1)))
	for k, v := range data {
		if rand >= v[0] && rand <= v[1] {
			key = k
			break
		}
	}
	if key == -1 {
		return key, errors.New("Dice Not Result ")
	}
	return key, nil
}

func Dices(arr map[int]int, num int, isUnique bool) (ret []int) {
	if isUnique {
		num = int(math.Min(float64(len(arr)), float64(num)))
	}
	for i := 0; i < num; i++ {
		id, err := Dice(arr)
		if err != nil {
			continue
		}
		if isUnique {
			delete(arr, id)
		}
		ret = append(ret, id)
	}
	return ret
}

func DicesInt32(arr map[int]int, num int, isUnique bool) (ret []int32) {
	if isUnique {
		num = int(math.Min(float64(len(arr)), float64(num)))
	}
	for {
		if num == 0 {
			break
		}
		id, err := Dice(arr)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		ret = append(ret, int32(id))
		if isUnique {
			delete(arr, id)
		}
		num--
	}
	return
}

func DoorDices(arr map[int]int, num, leng int, isUnique bool) (ret []int) {
	if leng < 1 {
		leng = 1
	}
	if isUnique {
		num = int(math.Min(float64(len(arr)), float64(num)))
	}
	for i := 0; i < num; i++ {
		id, err := Dice(arr)
		if err != nil {
			break
		}
		if isUnique {
			arr = DelRelation(arr, id, leng)
		}
		ret = append(ret, id)
	}
	return ret
}

func DelRelation(arr map[int]int, num, leng int) map[int]int {
	data := make(map[int]int)
	key1 := strconv.Itoa(num)
	key1 = string([]byte(key1)[:len(key1)-leng])
	for key, val := range arr {
		key2 := strconv.Itoa(key)
		key2 = string([]byte(key2)[:len(key2)-leng])
		if key == num || key1 == key2 {
			continue
		}
		data[key] = val
	}
	return data
}

func BuildWeight(arr map[int]map[string]interface{}, key string) (data map[int]int) {
	data = make(map[int]int)
	for k, v := range arr {
		data[k] = v[key].(int)
	}
	return data
}

func RandBySlice(list []int) int {
	if len(list) == 1 {
		return list[0]
	}
	key := RandInterval(int32(0), int32(len(list)-1))
	return list[key]
}

func RandInterval(b1, b2 int32) int32 {
	if b1 == b2 {
		return b1
	}
	min, max := int64(b1), int64(b2)
	if min > max {
		min, max = max, min
	}
	return int32(rand.Int63n(max-min+1) + min)
}

func RandFloatval(strList []string, per bool) float64 {
	num := float64(0)
	if len(strList) > 2 {
		return float64(0)
	}
	if len(strList) == 1 {
		min, _ := strconv.ParseFloat(strList[0], 10)
		return min
	}
	min, _ := strconv.ParseInt(strList[0], 10, 64)
	max, _ := strconv.ParseInt(strList[1], 10, 64)
	if min == max {
		return float64(min)
	}
	if per {
		min *= 100
		max *= 100
		num = float64(rand.Int63n(max-min-1)+min) / 100
	} else {
		num = float64(rand.Int63n(max-min-1) + min)
	}
	return num
}

func RandExceptSelf(start, end, self int) int {
	if self >= start && end >= self {
		if RandInterval(1, 100) > 50 {
			return int(RandInterval(int32(start), int32(self-1)))
		} else {
			return int(RandInterval(int32(self+1), int32(end)))
		}
	}
	return int(RandInterval(int32(start), int32(end)))
}

func RandAarryUnique(arr []int, num int) map[int]int {
	rs := make(map[int]int, num)
	if len(arr) <= num {
		for _, v := range arr {
			rs[v] = 1
		}
		return rs
	}
	for {
		if len(rs) >= num {
			break
		}
		randKey := RandInterval(int32(0), int32(len(arr))-1)
		rs[arr[randKey]] = 1
		arr = append(arr[:randKey], arr[randKey+1:]...)
	}
	return rs
}

func RandArrayStringUnique(arr []string, num int) string {
	rs := make([]string, 0)
	if len(arr) <= num {
		return strings.Join(arr, ",")
	}
	for {
		if len(rs) >= num {
			return strings.Join(rs, ",")
		}
		randKey := RandInterval(int32(0), int32(len(arr))-1)
		rs = append(rs, arr[randKey])
		arr = append(arr[:randKey], arr[randKey+1:]...)
	}
}
