package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
)

// GeneratePrivateKey generate RSA private key
func GeneratePrivateKey(bits int, out io.Writer) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	X509PrivateKey, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return errors.New("MarshalPKCS8PrivateKey error")
	}
	privateBlock := pem.Block{Type: "PRIVATE KEY", Bytes: X509PrivateKey}
	return pem.Encode(out, &privateBlock)
}

// GeneratePublicKey generate RSA public key
func GeneratePublicKey(priKey []byte, out io.Writer) error {
	block, _ := pem.Decode(priKey)
	if block == nil {
		return errors.New("key is invalid format")
	}
	// x509 parse
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	pk := privateKey.(*rsa.PrivateKey)
	//publicKey := privateKey.PublicKey
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&pk.PublicKey)
	if err != nil {
		return err
	}
	publicBlock := pem.Block{Type: "PUBLIC KEY", Bytes: X509PublicKey}
	return pem.Encode(out, &publicBlock)
}
