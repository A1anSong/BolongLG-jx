package lgjx

import (
	"encoding/base64"
	"github.com/tjfoc/gmsm/sm4"
)

func Sm4Encrypt(key, data []byte) string {
	_ = sm4.SetIV([]byte("0000000000000000"))
	cbcMsg, _ := sm4.Sm4Cbc(key, data, true)
	return base64.StdEncoding.EncodeToString(cbcMsg)
}

func Sm4Decrypt(key, msg []byte) (string, error) {
	_ = sm4.SetIV([]byte("0000000000000000"))
	cbcDec, err := sm4.Sm4Cbc([]byte(key), msg, false)
	return string(cbcDec), err
}
