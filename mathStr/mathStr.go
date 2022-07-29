package mathStr

import (
	"bytes"
	"errors"
	"math/big"
)

type (
	Math struct {
		Pre int
	}
	MathCmpType int
)

const (
	// MathCmpTypeEqual 相等
	MathCmpTypeEqual MathCmpType = 0

	// MathCmpTypeNotEqual 不等于
	MathCmpTypeNotEqual MathCmpType = 1

	// MathCmpTypeGreater 大于
	MathCmpTypeGreater MathCmpType = 2

	// MathCmpTypeEqualOrGreater 大于等于
	MathCmpTypeEqualOrGreater MathCmpType = 3

	// MathCmpTypeLessThan 小于
	MathCmpTypeLessThan MathCmpType = 4

	// MathCmpTypeEqualAndLessThan 小于等于
	MathCmpTypeEqualAndLessThan MathCmpType = 5
)

func NewMath(Pre int) *Math {
	t := new(Math)
	t.Pre = Pre
	return t
}

// Add 加法
func (t *Math) Add(strings ...string) string {
	if len(strings) > 0 {
		rs := new(big.Rat)
		for _, v := range strings {
			i := new(big.Rat)
			i.SetString(v)
			rs.Add(rs, i)
		}
		return rs.FloatString(t.Pre)
	}
	return "0"
}

// Sub 减法
func (t *Math) Sub(strings ...string) string {
	if len(strings) > 0 {
		rs := new(big.Rat)
		rs.SetString(strings[0])
		for i := 1; i < len(strings); i++ {
			v := new(big.Rat)
			v.SetString(strings[i])
			rs.Sub(rs, v)
		}
		return rs.FloatString(t.Pre)
	}
	return "0"
}

// Mul 乘法
func (t *Math) Mul(strings ...string) string {
	if len(strings) > 0 {
		rs := new(big.Rat)
		rs.SetString("1")
		for _, v := range strings {
			i := new(big.Rat)
			i.SetString(v)
			rs.Mul(rs, i)
		}
		return rs.FloatString(t.Pre)
	}
	return "0"
}

// Quo 除法
func (t *Math) Quo(first string, strings ...string) (string, error) {
	b := bytes.TrimSpace([]byte(first))
	first = string(b)
	if first == "" || first == "0" {
		return "0", errors.New("first 参数错误")
	}
	if len(strings) > 0 {
		f := new(big.Rat)
		f.SetString(first)
		for _, v := range strings {
			b := bytes.TrimSpace([]byte(v))
			v = string(b)
			if v == "" || v == "0" {
				return "0", errors.New("Quo 错误 参数存在 %s " + v)
			}
			i := new(big.Rat)
			i.SetString(v)
			f.Quo(f, i)
		}
		return f.FloatString(t.Pre), nil
	}
	return "0", errors.New("Quo 错误 列表为空")
}

// Cmp 对比
func (t *Math) Cmp(num1, num2 string, typ MathCmpType) bool {
	p1 := new(big.Rat)
	p1.SetString(num1)
	p2 := new(big.Rat)
	p2.SetString(num2)
	rs := p1.Cmp(p2)
	switch typ {
	case MathCmpTypeEqual:
		return rs == 0
	case MathCmpTypeNotEqual:
		return rs != 0
	case MathCmpTypeEqualOrGreater:
		if rs == 0 {
			return true
		}
		fallthrough
	case MathCmpTypeGreater:
		return rs > 0
	case MathCmpTypeEqualAndLessThan:
		if rs == 0 {
			return true
		}
		fallthrough
	case MathCmpTypeLessThan:
		return rs < 0
	default:
		return false
	}
}

// Min ...
func (t *Math) Min(num1 string, num2 string) (rs string) {
	p1 := new(big.Rat)
	p1.SetString(num1)
	p2 := new(big.Rat)
	p2.SetString(num2)
	if t.Cmp(num1, num2, MathCmpTypeLessThan) {
		return p1.FloatString(t.Pre)
	}
	return p2.FloatString(t.Pre)
}

// Max ...
func (t *Math) Max(num1 string, num2 string) (rs string) {
	p1 := new(big.Rat)
	p1.SetString(num1)
	p2 := new(big.Rat)
	p2.SetString(num2)
	if t.Cmp(num1, num2, MathCmpTypeGreater) {
		return p1.FloatString(t.Pre)
	}
	return p2.FloatString(t.Pre)
}

// GetWholeRemaining 获取整余
func (t *Math) GetWholeRemaining(num1, num2 string) (whole, remaining string) {
	whole = "0"
	remaining = "0"
	if num1 == "" {
		num1 = "1"
	}
	if num2 == "" {
		num2 = "0"
	}
	for {
		if t.Cmp(num1, num2, MathCmpTypeLessThan) {
			//如果num1小于num2则停止
			remaining = num1
			break
		}
		num1 = t.Sub(num1, num2)
		whole = t.Add(whole, "1")
	}
	return whole, remaining
}
