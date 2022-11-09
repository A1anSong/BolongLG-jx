package lgjx

import (
	"encoding/base64"
	"github.com/tjfoc/gmsm/sm4"
)

func Sm4Encrypt(key, data string) string {
	_ = sm4.SetIV([]byte("0000000000000000"))
	cbcMsg, _ := sm4.Sm4Cbc([]byte(key), []byte(data), true)
	return base64.StdEncoding.EncodeToString(cbcMsg)
}

func Sm4Decrypt(key, msg string) string {
	_ = sm4.SetIV([]byte("0000000000000000"))
	cbcMsg, _ := base64.StdEncoding.DecodeString(msg)
	cbcDec, _ := sm4.Sm4Cbc([]byte(key), cbcMsg, false)
	return string(cbcDec)
}
