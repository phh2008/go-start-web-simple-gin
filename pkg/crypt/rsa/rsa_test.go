package rsa

import (
	"bytes"
	"crypto"
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestRSAEncrypt(t *testing.T) {
	priBuf := bytes.NewBuffer(nil)
	err := GeneratePrivateKey(2048, priBuf)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("私钥：")
	fmt.Println(string(priBuf.Bytes()))
	pubBuf := bytes.NewBuffer(nil)
	err = GeneratePublicKey(priBuf.Bytes(), pubBuf)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("公钥：")
	fmt.Println(string(pubBuf.Bytes()))
	var src = []byte("山不在高，有仙则名。水不在深，有龙则灵。斯是陋室，惟吾德馨。苔痕上阶绿，草色入帘青。谈笑有鸿儒，往来无白丁。")
	enc, err := EncryptByPubKey(src, pubBuf.Bytes())
	if err != nil {
		panic(err)
	}
	fmt.Println("加密后：", enc)
	val, _ := base64.StdEncoding.DecodeString(enc)
	ret, err := DecryptByPriKey(val, priBuf.Bytes())
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后：", string(ret))
}

func TestRSAEncrypt2(t *testing.T) {
	priBuf := bytes.NewBuffer(nil)
	err := GeneratePrivateKey(2048, priBuf)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("私钥：")
	fmt.Println(string(priBuf.Bytes()))
	pubBuf := bytes.NewBuffer(nil)
	err = GeneratePublicKey(priBuf.Bytes(), pubBuf)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("公钥：")
	fmt.Println(string(pubBuf.Bytes()))
	var src = []byte("山不在高，有仙则名。水不在深，有龙则灵。斯是陋室，惟吾德馨。苔痕上阶绿，草色入帘青。谈笑有鸿儒，往来无白丁。")
	enc, err := EncryptByPriKey(src, priBuf.Bytes())
	if err != nil {
		panic(err)
	}
	fmt.Println("加密后：", enc)
	val, _ := base64.StdEncoding.DecodeString(enc)
	ret, err := DecryptByPubKey(val, pubBuf.Bytes())
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后：", string(ret))
}

func TestRSASign(t *testing.T) {
	priBuf := bytes.NewBuffer(nil)
	err := GeneratePrivateKey(2048, priBuf)
	t.Logf("private key: %s\n", priBuf.Bytes())

	pubBuf := bytes.NewBuffer(nil)
	err = GeneratePublicKey(priBuf.Bytes(), pubBuf)
	assert.NoError(t, err)
	t.Logf("public key: %s\n", pubBuf.Bytes())

	src := []byte("123456")
	sign, err := CreateSign(src, priBuf.Bytes(), crypto.SHA256)
	assert.NoError(t, err)
	t.Logf("sign out: %s\n", base64.RawStdEncoding.EncodeToString(sign))

	err = VerifySign(src, sign, pubBuf.Bytes(), crypto.SHA256)
	assert.NoError(t, err)
}
