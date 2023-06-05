package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
)

// Key 密钥对
type Key struct {
	PrivateKey string
	PublicKey  string
}

// GenerateRsaKeyHex 生成 rsa key, bits is 1024 or 2048
func GenerateRsaKeyHex(bits int) (Key, error) {
	if bits != 1024 && bits != 2048 {
		return Key{}, ErrRsaBits
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return Key{}, err
	}
	return Key{
		PrivateKey: hex.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey)),
		PublicKey:  hex.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)),
	}, nil
}

// GenerateRsaKeyBase64 生成 rsa key, bits is 1024 or 2048
func GenerateRsaKeyBase64(bits int) (Key, error) {
	if bits != 1024 && bits != 2048 {
		return Key{}, ErrRsaBits
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return Key{}, err
	}
	return Key{
		PrivateKey: base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey)),
		PublicKey:  base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)),
	}, nil
}
