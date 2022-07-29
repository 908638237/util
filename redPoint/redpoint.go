package redPoint

import (
	"github.com/908638237/util/common"
	"github.com/908638237/util/log"
	"strings"
)

type RedPoint struct {
	Parent *RedPoint            `json:"-"`
	SList  map[string]*RedPoint `json:"list,omitempty"`
	IList  map[int32]*RedPoint  `json:"-"`
	Name   string               `json:"name"`
	Val    int32                `json:"val"`
}

func NewRedPoint() *RedPoint {
	r := new(RedPoint)
	r.SList = make(map[string]*RedPoint)
	r.IList = make(map[int32]*RedPoint)
	return r
}

func (t *RedPoint) Add(name string, val int32) *RedPoint {
	item := NewRedPoint()
	item.Name = name
	item.Val = val
	item.Parent = t

	t.SList[name] = item
	t.IList[val] = item

	if val <= 0 {
		log.Error("红点值小于0 %s", name)
	}

	return item
}

func (t *RedPoint) AddChildren(data map[string]int32) *RedPoint {
	for name, val := range data {
		item := NewRedPoint()
		item.Name = name
		if val <= 0 {
			log.Error("红点值小于0 %s", name)
		}
		item.Val = val
		item.Parent = t
		t.SList[name] = item
		t.IList[val] = item
	}
	return t
}

func (t *RedPoint) AddChild(name string, val int32) *RedPoint {
	item := NewRedPoint()
	item.Name = name
	item.Val = val
	item.Parent = t
	if val <= 0 {
		log.Error("红点值小于0 %s", name)
	}
	t.SList[name] = item
	t.IList[val] = item

	return t
}

func (t *RedPoint) GetChildVal() []int32 {
	list := []int32{}
	for _, v := range t.IList {
		list = append(list, v.Val)
	}
	return list
}

func (t *RedPoint) GetChild() []*RedPoint {
	list := []*RedPoint{}
	for _, v := range t.IList {
		list = append(list, v)
	}
	return list
}

func (t *RedPoint) Get(path string) *RedPoint {
	pathList := strings.Split(path, ".")
	node := t
	for _, v := range pathList {
		if node.Has(v) {
			node = node.SList[v]
		} else {
			return nil
		}
	}
	return node
}

func (t *RedPoint) GetByVal(val int32) *RedPoint {
	var node *RedPoint
	if t.Val == val {
		return t
	}
	for _, v := range t.IList {
		if len(t.IList) > 0 {
			if node, ok := v.IList[val]; ok {
				return node
			}
			rs := v.GetByVal(val)
			if rs != nil {
				return rs
			}
		}
	}
	return node
}

func (t *RedPoint) Has(name string) bool {
	_, ok := t.SList[name]
	return ok
}

func (t *RedPoint) GetParents() []*RedPoint {
	list := make([]*RedPoint, 0)
	item := t
	times := 0
	for {
		if times > 10 {
			return list
		}
		times++
		if item.Parent == nil {
			return list
		}
		list = append(list, item.Parent)
	}
}

func (t *RedPoint) GetParentsVal() []int32 {
	list := make([]int32, 0)
	times := 0
	item := t
	if item == nil {
		return list
	}
	for {
		if times > 10 {
			return list
		}
		times++

		if item.Parent == nil || item.Parent.Val <= 0 {
			return list
		}

		list = append(list, item.Parent.Val)

		item = item.Parent

		if item == nil {
			return list
		}
	}
}

func (t *RedPoint) GetJson() []byte {
	v, _ := common.JsonMarshal(t)
	return v
}
