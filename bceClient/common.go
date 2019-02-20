package bceClient

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func hmacSha256Hex(key, content string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(content))
	return hex.EncodeToString(hash.Sum(nil))
}