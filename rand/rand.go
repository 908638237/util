package rand

import (
	"errors"
	"fmt"
	"github.com/908638237/util/log"
	"github.com/golang-module/carbon/v2"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

type Prize struct {
	Id     int
	Weight int
	Args   any
}

type (
	Slice interface {
		[]int | []int32 | []int64 | []string | []float32 | []float64 | []uint8 | []uint16 | []uint32 | []uint64
	}
	Number interface {
		int | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
	}
)

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Dice(arr map[int]int) (key int, err error) {
	start := int32(0)
	max := int32(1)
	key = -1
	data := make(map[int][2]int32)
	for k, v := range arr {
		if v < 1 {
			continue
		}
		end := start + int32(v)
		data[k] = [2]int32{start, end}
		start = end + 1
		max = end
	}
	randomNum := RandInterval(0, int32(math.Max(float64(max), 1)))
	for k, v := range data {
		if randomNum >= v[0] && randomNum <= v[1] {
			key = k
			break
		}
	}
	if key == -1 {
		return key, errors.New("Dice Not Result ")
	}
	return key, nil
}

func DiceString(arr map[string]int) (key string, err error) {
	start := int32(0)
	max := int32(1)
	key = ""
	data := make(map[string][2]int32)
	for k, v := range arr {
		if v < 1 {
			continue
		}
		end := start + int32(v)
		data[k] = [2]int32{start, end}
		start = end + 1
		max = end
	}
	randomNum := RandInterval(0, int32(math.Max(float64(max), 1)))
	for k, v := range data {
		if randomNum >= v[0] && randomNum <= v[1] {
			key = k
			break
		}
	}
	if key == "" {
		return key, errors.New("Dice Not Result ")
	}
	return key, nil
}

func RandomDraw(count int, arr []Prize, ignore []int) int {
	start, end := 0, 0
	randomNum := rand.Intn(count)
	if randomNum < 1 {
		randomNum = 1
	}

	if ignore != nil && len(ignore) > 0 {
		for _, id := range ignore {
			for k, v := range arr {
				if id == v.Id {
					arr = append(arr[:k], arr[k+1:]...)
				}
			}
		}
	}
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i].Weight < arr[j].Weight
	})
	//quality := arr[len(arr)-1].Id

	for _, item := range arr {
		//if randomNum <= item.Weight {
		//	return item.Id
		//}
		end += item.Weight
		if start <= randomNum && end >= randomNum {
			return item.Id
		}
		start = end
	}
	//return quality
	return -1
}

func RandomActDraw(count int, arr []Prize, ignore []int) (id int, limit bool) {
	start, end := 0, 0
	randomNum := rand.Intn(count)
	if randomNum < 1 {
		randomNum = 1
	}
	if ignore != nil && len(ignore) > 0 {
		for _, id := range ignore {
			for k, v := range arr {
				if id == v.Id {
					arr = append(arr[:k], arr[k+1:]...)
				}
			}
		}
	}
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i].Weight < arr[j].Weight
	})
	//quality := arr[len(arr)-1].Id

	for _, item := range arr {
		//if randomNum <= item.Weight {
		//	return item.Id
		//}
		end += item.Weight
		if start <= randomNum && end >= randomNum {
			id = item.Id
			if item.Args.(int) > 0 {
				limit = true
			}
			return
		}
		start = end
	}
	//return quality
	return -1, false
}

func RandomDrawItem(count int, arr []Prize, ignore []int) Prize {
	start, end := 0, 0
	randomNum := rand.Intn(count)
	if randomNum < 1 {
		randomNum = 1
	}
	if ignore != nil && len(ignore) > 0 {
		for _, id := range ignore {
			for k, v := range arr {
				if id == v.Id {
					arr = append(arr[:k], arr[k+1:]...)
				}
			}
		}
	}
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i].Weight < arr[j].Weight
	})
	for _, item := range arr {
		end += item.Weight
		if start <= randomNum && end >= randomNum {
			return item
		}
		start = end
	}
	return Prize{}
}

func RandomDraws(count int, arr []Prize, num int, isUnique bool) (ret []int) {
	if isUnique {
		num = Min(num, len(arr))
	}
	for i := 0; i < num; i++ {
		id := RandomDraw(count, arr, nil)
		if id > 0 {
			ret = append(ret, id)
			if isUnique {
				for k, v := range arr {
					if v.Id == id {
						arr = append(arr[:k], arr[k+1:]...)
						count -= v.Weight
						break
					}
				}
			}
		}
	}
	return ret
}

func Seed(seed int64) {
	rand.Seed(seed)
}

type DiceItem struct {
	Key    int
	Weight int
}

func buildDice(arr map[int]int) []DiceItem {
	data := make([]DiceItem, 0)
	if len(arr) < 1 {
		return data
	}
	for k, v := range arr {
		if v < 1 {
			continue
		}
		data = append(data, DiceItem{
			Key:    k,
			Weight: v,
		})
	}
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Key < data[j].Key
	})
	return data
}

func NewDice(arr map[int]int) (key int, err error) {
	start, max := int32(0), int32(1)
	key = -1
	d := buildDice(arr)
	data := make(map[int][2]int32)
	for _, v := range d {
		end := start + int32(v.Weight)
		data[v.Key] = [2]int32{start, end}
		start = end + 1
		max = end
	}
	rd := RandInterval(0, Max(max, 1))
	for k, v := range data {
		if rd >= v[0] && rd <= v[1] {
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
			log.Error("%s", err)
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
			log.Error("%s", err)
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

func DicesString(arr map[string]int, num int, isUnique bool) (ret []string) {
	if isUnique {
		num = int(math.Min(float64(len(arr)), float64(num)))
	}
	for {
		if num == 0 {
			break
		}
		key, err := DiceString(arr)
		if err != nil {
			log.Error("%s", err)
			continue
		}
		ret = append(ret, key)
		if isUnique {
			delete(arr, key)
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
			log.Error("%s", err)
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

func RandRangeBySlice(list []int) int {
	if len(list) < 1 {
		return 0
	}
	if len(list) == 1 {
		return list[0]
	}
	return int(RandInterval(int32(list[0]), int32(list[1])))
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

func RandIntervalBySource(b1, b2 int32, source int64) int32 {
	if b1 == b2 {
		return b1
	}
	min, max := int64(b1), int64(b2)
	if min > max {
		min, max = max, min
	}
	if source < 1 {
		source = carbon.Now().Timestamp()
	}
	r := rand.New(rand.NewSource(source))
	return int32(r.Int63n(max-min+1) + min)
}

func RandIntRange(list []int) int {
	if len(list) == 0 {
		return 0
	}
	if len(list) < 2 {
		return list[0]
	}
	return int(RandInterval(int32(list[0]), int32(list[1])))
}

func RandFloatval(strList []string, per bool) float64 {
	num := float64(0)
	if len(strList) > 2 {
		return float64(0)
	}
	if len(strList) == 1 {
		min, _ := strconv.ParseFloat(strList[0], 64)
		return min
	}
	min, _ := strconv.ParseFloat(strList[0], 64)
	max, _ := strconv.ParseFloat(strList[1], 64)
	if min == max {
		return min
	}
	if per {
		min *= 100
		max *= 100
		num = float64(RandInterval(int32(min), int32(max))) / 100
	} else {
		num = float64(RandInterval(int32(min), int32(max)))
	}
	return num
}

func RandFloatValByInt(b1, b2 int) float64 {
	num := float64(0)
	rand.Float64()
	return num
}

func RandIntval(strList []string, per bool) int64 {
	num := int64(0)
	if len(strList) > 2 {
		return num
	}
	if len(strList) == 1 {
		min, _ := strconv.ParseInt(strList[0], 10, 64)
		return min
	}
	min, _ := strconv.ParseInt(strList[0], 10, 64)
	max, _ := strconv.ParseInt(strList[1], 10, 64)
	if min == max {
		return min
	}

	//if per {
	//	min *= 100
	//	max *= 100
	//
	//	num = float64(rand.Int63n(max-min-1)+min) / 100
	//} else {
	num = rand.Int63n(max-min-1) + min
	//}
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
