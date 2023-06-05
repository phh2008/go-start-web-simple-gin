package rsa

import (
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"testing"
)

// TestGenerateRsaKey 测试生成密钥对
func TestGenerateRsaKey(t *testing.T) {
	key, err := GenerateRsaKeyBase64(1024)
	if err != nil {
		panic(err)
	}
	fmt.Println("PrivateKey: ", key.PrivateKey)
	fmt.Println("PublicKey: ", key.PublicKey)
}

// TestPemEncode 生成密钥对 pem格式化
func TestPemEncode(t *testing.T) {
	key, err := GenerateRsaKeyBase64(1024)
	if err != nil {
		panic(err)
	}
	privateBytes, _ := base64.StdEncoding.DecodeString(key.PrivateKey)

	privatePEMBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateBytes,
		},
	)
	fmt.Println("私钥-base64：", key.PrivateKey)
	fmt.Println("私钥：", string(privatePEMBytes))
	//pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	pubBytes, _ := base64.StdEncoding.DecodeString(key.PublicKey)
	pubPEMBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubBytes,
		},
	)
	fmt.Println("公钥-base64：", key.PublicKey)
	fmt.Println("公钥：", string(pubPEMBytes))
}
