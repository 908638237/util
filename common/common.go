package common

import jsoniter "github.com/json-iterator/go"

type (
	DataType interface {
		int | int32 | int64 | uint | uint32 | uint64 | float32 | float64 | string
	}
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func JsonMarshal(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func JsonUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
