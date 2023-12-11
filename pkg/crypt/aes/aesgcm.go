package aes

import (
	"com.gientech/selection/pkg/crypt"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// AesGCM aes gcm
type AesGCM struct {
	Key []byte
	IV  []byte
}

// NewAesGCM new aes gcm
func NewAesGCM(key []byte, iv ...byte) *AesGCM {
	if len(iv) == 0 {
		iv = key[:12]
	}
	return &AesGCM{
		Key: key,
		IV:  iv,
	}
}

// encrypt 加密
func (obj *AesGCM) encrypt(plainText []byte, withIV bool) ([]byte, error) {
	if len(obj.Key) != 16 && len(obj.Key) != 24 && len(obj.Key) != 32 {
		return nil, crypt.ErrKeyLength
	}

	block, err := aes.NewCipher(obj.Key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	if len(obj.IV) != aesgcm.NonceSize() {
		aesgcm, err = cipher.NewGCMWithNonceSize(block, len(obj.IV))
		if err != nil {
			return nil, err
		}
	}

	var cipherText []byte
	if withIV {
		cipherText = aesgcm.Seal(obj.IV, obj.IV, plainText, nil)
	} else {
		cipherText = aesgcm.Seal(nil, obj.IV, plainText, nil)
	}

	return cipherText, nil
}

// EncryptWithIV 加密,带有iv前缀
func (obj *AesGCM) EncryptWithIV(plainText []byte) ([]byte, error) {
	return obj.encrypt(plainText, true)
}

// Encrypt 加密，无前缀
func (obj *AesGCM) Encrypt(plainText []byte) ([]byte, error) {
	return obj.encrypt(plainText, false)
}

// EncryptBase64WithIV 加密base64，带iv前缀
func (obj *AesGCM) EncryptBase64WithIV(plainText []byte) (string, error) {
	cipherText, err := obj.EncryptWithIV(plainText)
	if err != nil {
		return "", err
	}
	ret := base64.StdEncoding.EncodeToString(cipherText)
	return ret, nil
}

// EncryptBase64 加密base64
func (obj *AesGCM) EncryptBase64(plainText []byte) (string, error) {
	cipherText, err := obj.Encrypt(plainText)
	if err != nil {
		return "", err
	}
	ret := base64.StdEncoding.EncodeToString(cipherText)
	return ret, nil
}

// decrypt 解密
func (obj *AesGCM) decrypt(cipherText []byte, withIV bool) ([]byte, error) {
	if len(obj.Key) != 16 && len(obj.Key) != 24 && len(obj.Key) != 32 {
		return nil, crypt.ErrKeyLength
	}
	if len(cipherText) < aes.BlockSize {
		return nil, crypt.ErrCipherText
	}

	block, err := aes.NewCipher(obj.Key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	if len(obj.IV) != aesgcm.NonceSize() {
		aesgcm, err = cipher.NewGCMWithNonceSize(block, len(obj.IV))
		if err != nil {
			return nil, err
		}
	}

	var plainText []byte
	if withIV {
		iv, cipherText := cipherText[:aesgcm.NonceSize()], cipherText[aesgcm.NonceSize():]
		plainText, err = aesgcm.Open(nil, iv, cipherText, nil)
	} else {
		plainText, err = aesgcm.Open(nil, obj.IV, cipherText, nil)
	}
	if err != nil {
		return nil, err
	}

	return plainText, err
}

// DecryptWithIV 解密，带有iv前缀
func (obj *AesGCM) DecryptWithIV(cipherText []byte) ([]byte, error) {
	return obj.decrypt(cipherText, true)
}

// Decrypt 解密
func (obj *AesGCM) Decrypt(cipherText []byte) ([]byte, error) {
	return obj.decrypt(cipherText, false)
}

// DecryptBase64WithIV 解密base64，带有iv前缀
func (obj *AesGCM) DecryptBase64WithIV(cipherStr string) ([]byte, error) {
	cipherText, err := base64.StdEncoding.DecodeString(cipherStr)
	if err != nil {
		return nil, err
	}
	return obj.DecryptWithIV(cipherText)
}

// DecryptBase64 解密base64
func (obj *AesGCM) DecryptBase64(cipherStr string) ([]byte, error) {
	cipherText, err := base64.StdEncoding.DecodeString(cipherStr)
	if err != nil {
		return nil, err
	}
	return obj.Decrypt(cipherText)
}
