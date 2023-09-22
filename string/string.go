package _string

import (
	"regexp"
	"strings"
)

func BuilderConcat(str ...string) string {
	if len(str) == 0 {
		return ""
	}
	if len(str) < 2 {
		return str[0]
	}
	var builder strings.Builder
	builder.Grow(len(str))
	for _, v := range str {
		builder.WriteString(v)
	}
	return builder.String()
}

func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	// 返回子串
	return string(rs[begin:end])
}

// ToHump 把str转成驼峰型str 例: role_info => RoleInfo
func ToHump(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArr := []byte(str)
	strArrLen := len(strArr)
	firstStr := strings.ToUpper(string(strArr[0]))
	input := []byte(string(strArr[1:strArrLen]))
	reg := regexp.MustCompile(`([_-][a-z])`)
	out := reg.ReplaceAllFunc(input, replaceUpper)
	return firstStr + string(out)
}

// HumpToString 驼峰型转为字符串, symbolType 符号类型 默认"_", upper 是否把首字符转为大写 默认转为小写
func HumpToString(str, symbolType string, upper bool) string {
	if len(str) < 1 {
		return ""
	}
	if symbolType == "" {
		symbolType = "_"
	}
	strArr := []byte(str)
	strArrLen := len(strArr)
	firstStr := strings.ToLower(string(strArr[0]))
	if upper {
		firstStr = string(strArr[0])
	}
	input := []byte(string(strArr[1:strArrLen]))
	reg := regexp.MustCompile(`[A-Z]`)
	out := make([]byte, 0)
	switch symbolType {
	case "-":
		if upper {
			out = reg.ReplaceAllFunc(input, replaceCrossBarToUpper)
		} else {
			out = reg.ReplaceAllFunc(input, replaceCrossBar)
		}
	default:
		if upper {
			out = reg.ReplaceAllFunc(input, replaceUnderlineToUpper)
		} else {
			out = reg.ReplaceAllFunc(input, replaceUnderline)
		}
	}
	return firstStr + string(out)
}

func replaceUpper(strByte []byte) []byte {
	if len(strByte) < 1 {
		return []byte{}
	}
	return []byte(strings.ToUpper(string(strByte[1])))
}

// replaceCrossBarToUpper 以-分割string(大写)
func replaceCrossBarToUpper(strByte []byte) []byte {
	if len(strByte) < 1 {
		return []byte{}
	}
	return []byte("-" + string(strByte))
}

// replaceUnderlineToUpper 以_分割string(大写)
func replaceUnderlineToUpper(strByte []byte) []byte {
	if len(strByte) < 1 {
		return []byte{}
	}
	return []byte("_" + string(strByte))
}

// replaceCrossBar 以-分割string(小写)
func replaceCrossBar(strByte []byte) []byte {
	if len(strByte) < 1 {
		return []byte{}
	}
	return []byte("-" + strings.ToLower(string(strByte)))
}

// replaceUnderline 以_分割string(小写)
func replaceUnderline(strByte []byte) []byte {
	if len(strByte) < 1 {
		return []byte{}
	}
	return []byte("_" + strings.ToLower(string(strByte)))
}
