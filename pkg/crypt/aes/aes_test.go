package main

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	var key = []byte("7uYhf14h094&feki")
	var iv = "8AYhf14h094ES4jk"
	var src = "abc"
	enc, err := Encrypt(src, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Println(enc)
}

func TestDecrypt(t *testing.T) {
	openId := "waq9gD0pCvwb0QqByby4yQ=="
	var key = []byte("7uYhf14h094&feki")
	var iv = "8AYhf14h094ES4jk"
	str, err := Decrypt(openId, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
