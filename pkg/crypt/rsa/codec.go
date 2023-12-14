package rsa

import (
	"encoding/base64"
	"encoding/hex"
	"unsafe"
)

type CodecResult struct {
	Raw   []byte
	Error error
}

func (a CodecResult) String() string {
	return a.ToRawString()
}

// ToRawString outputs as raw string without encoding.
// 输出未经编码的原始字符串
func (a CodecResult) ToRawString() string {
	if len(a.Raw) == 0 || a.Error != nil {
		return ""
	}
	return bytes2string(a.Raw)
}

// ToHexString outputs as string with hex encoding.
// 输出经过 hex 编码的字符串
func (a CodecResult) ToHexString() string {
	if len(a.Raw) == 0 || a.Error != nil {
		return ""
	}
	return hex.EncodeToString(a.Raw)
}

// ToBase64String outputs as string with base64 encoding.
// 输出经过 base64 编码的字符串
func (a CodecResult) ToBase64String() string {
	if len(a.Raw) == 0 || a.Error != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(a.Raw)
}

// ToRawBytes outputs as raw byte slice without encoding.
// 输出未经编码的原始字节切片
func (a CodecResult) ToRawBytes() []byte {
	if len(a.Raw) == 0 {
		return []byte("")
	}
	return a.Raw
}

func interface2bytes(i interface{}) (b []byte) {
	switch v := i.(type) {
	case string:
		b = string2bytes(v)
	case []byte:
		b = v
	}
	return
}

func string2bytes(s string) []byte {
	if len(s) == 0 {
		return []byte("")
	}
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func bytes2string(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&b))
}
