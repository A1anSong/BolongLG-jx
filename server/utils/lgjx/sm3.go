package lgjx

import (
	"encoding/base64"
	"github.com/tjfoc/gmsm/sm3"
)

func SM3Sign(data string) string {
	h := sm3.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(sum)
}
