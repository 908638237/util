package md5

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func MD5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}

func Get16MD5(str string) string {
	return MD5(str)[8:24]
}

func Get2MD5(str string) string {
	return MD5(str)[0:2]
}

func Sha256(str string) string {
	ctx := sha256.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}
