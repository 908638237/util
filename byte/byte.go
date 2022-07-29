package byte

import (
	"bytes"
	"encoding/binary"
	"github.com/908638237/util/common"
)

func DataToBytes[T common.DataType](data T) (rs []byte, err error) {
	bytebuffer := bytes.NewBuffer([]byte{})
	if err = binary.Write(bytebuffer, binary.BigEndian, data); err != nil {
		return
	}
	rs = bytebuffer.Bytes()
	return
}

func BytesToData[T common.DataType](bys []byte) (rs T, err error) {
	bytebuffer := bytes.NewBuffer(bys)
	if err = binary.Read(bytebuffer, binary.BigEndian, &rs); err != nil {
		return
	}
	return
}
