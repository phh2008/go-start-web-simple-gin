package rsa

import (
	"bytes"
	"com.gientech/selection/pkg/logger"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"math/big"
	"runtime"
)

func encryptByPublicKey(plainText, publicKey []byte) (cipherText []byte, err error) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				logger.Errorf("runtime err=%v,Check that the key or text is correct", err)
			default:
				logger.Errorf("error=%v,check the cipherText ", err)
			}
		}
	}()
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("key is invalid format")
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pk, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("the kind of key is not a rsa.PublicKey")
	}
	pubSize, plainTextSize := pk.Size(), len(plainText)
	offSet, once := 0, pubSize-11
	buffer := bytes.Buffer{}
	for offSet < plainTextSize {
		endIndex := offSet + once
		if endIndex > plainTextSize {
			endIndex = plainTextSize
		}
		bytesOnce, err := rsa.EncryptPKCS1v15(rand.Reader, pk, plainText[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	cipherText = buffer.Bytes()
	return cipherText, nil
}

func decryptByPrivateKey(cipherText, privateKey []byte) (plainText []byte, err error) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				logger.Errorf("runtime err=%v,Check that the key or text is correct", err)
			default:
				logger.Errorf("error=%v,check the cipherText ", err)
			}
		}
	}()
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("key is invalid format")
	}
	pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return []byte{}, err
	}
	pk := pri.(*rsa.PrivateKey)
	priSize, cipherTextSize := pk.Size(), len(cipherText)
	var offSet = 0
	var buffer = bytes.Buffer{}
	for offSet < cipherTextSize {
		endIndex := offSet + priSize
		if endIndex > cipherTextSize {
			endIndex = cipherTextSize
		}
		bytesOnce, err := rsa.DecryptPKCS1v15(rand.Reader, pk, cipherText[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	plainText = buffer.Bytes()
	return plainText, nil
}

// EncryptByPubKey 加密
func EncryptByPubKey(plainText []byte, pubKey []byte) (base64CipherText string, err error) {
	cipherBytes, err := encryptByPublicKey(plainText, pubKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherBytes), nil
}

// DecryptByPriKey 解密
func DecryptByPriKey(cipherTextBytes []byte, priKey []byte) (plainText []byte, err error) {
	return decryptByPrivateKey(cipherTextBytes, priKey)
}

// EncryptByPriKey 私钥加密（加密内容不能超过密钥长度）
func EncryptByPriKey(plainText []byte, priKey []byte) (base64CipherText string, err error) {
	// 获取私钥
	block, _ := pem.Decode(priKey)
	if block == nil {
		return "", errors.New("key is invalid format")
	}
	// x509 parse
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pk := privateKey.(*rsa.PrivateKey)
	// 签名
	cipherBytes, err := rsa.SignPKCS1v15(rand.Reader, pk, crypto.Hash(0), plainText)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherBytes), nil
}

// DecryptByPubKey 公钥解密
func DecryptByPubKey(cipherBytes []byte, publicKey []byte) (plainText []byte, err error) {
	// 获取公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("key is invalid format")
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pk, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("the kind of key is not a rsa.PublicKey")
	}
	c := new(big.Int)
	m := new(big.Int)
	m.SetBytes(cipherBytes)
	e := big.NewInt(int64(pk.E))
	c.Exp(m, e, pk.N)
	out := c.Bytes()
	skip := 0
	for i := 2; i < len(out); i++ {
		if i+1 >= len(out) {
			break
		}
		if out[i] == 0xff && out[i+1] == 0 {
			skip = i + 2
			break
		}
	}
	return out[skip:], nil
}

// CreateSign RSA sign
func CreateSign(src []byte, priKey []byte, hash crypto.Hash) ([]byte, error) {
	block, _ := pem.Decode(priKey)
	if block == nil {
		return nil, errors.New("key is invalid format")
	}
	// x509 parse
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pk := privateKey.(*rsa.PrivateKey)
	h := hash.New()
	_, err = h.Write(src)
	if err != nil {
		return nil, err
	}
	bt := h.Sum(nil)
	sign, err := rsa.SignPKCS1v15(rand.Reader, pk, hash, bt)
	if err != nil {
		return nil, err
	}
	return sign, nil
}

// VerifySign RSA verify
func VerifySign(src, sign, pubKey []byte, hash crypto.Hash) error {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return errors.New("key is invalid format")
	}
	// x509 parse
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		return errors.New("the kind of key is not a rsa.PublicKey")
	}
	h := hash.New()
	_, err = h.Write(src)
	if err != nil {
		return err
	}
	bt := h.Sum(nil)
	return rsa.VerifyPKCS1v15(publicKey, hash, bt, sign)
}
