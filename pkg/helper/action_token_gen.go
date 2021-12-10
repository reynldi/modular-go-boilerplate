package helper

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func GenSHA256(str string) string{
	h := sha256.New()
	h.Write([]byte(str))

	return base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(h.Sum(nil))))
}
