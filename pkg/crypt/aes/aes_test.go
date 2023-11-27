package aes

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	var key = []byte("7uYhf14h094&feki")
	var iv = []byte("8AYhf14h094ES4jk")
	var src = "abc"
	enc, err := AesCBCEncrypt([]byte(src), key, iv, PKCS5_PADDING)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(enc))
}

func TestDecrypt(t *testing.T) {
	val, _ := base64.StdEncoding.DecodeString("waq9gD0pCvwb0QqByby4yQ==")
	var key = []byte("7uYhf14h094&feki")
	var iv = []byte("8AYhf14h094ES4jk")
	str, err := AesCBCDecrypt([]byte(val), key, iv, PKCS5_PADDING)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(str))
}

func TestECBEncrypt(t *testing.T) {
	// AES/ECB/PKCS5Padding
	var key = []byte("1kZTLwGpHGSSgWlN1fqYsSVdltygxzRT")
	var src = []byte("君不见黄河之水天上来，奔流到海不复回。君不见高堂明镜悲白发，朝如青丝暮成雪。人生得意须尽欢，莫使金樽空对月。五花马，千金裘，呼儿将出换美酒。与尔同消万古愁。岑夫子丹丘生，将进酒，杯莫停，与君歌一曲，请军为我倾耳听，钟鼓馔玉不足贵，但愿长醉不复醒。")
	enc, err := AesECBEncrypt(src, key, PKCS5_PADDING)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(enc))
}

func TestECBDecrypt(t *testing.T) {
	// AES/ECB/PKCS5Padding
	var key = []byte("1kZTLwGpHGSSgWlN1fqYsSVdltygxzRT")
	var val, _ = base64.StdEncoding.DecodeString("dFt9AaPRilSvH5VDGC1AUHPYnYaF2FmnSgANhZz9ZB23zZcdlhRvROZLdDnM6L3ojf4QUKcl7FzcCCNZxX5vJZBtv5r3t0fARPTNn6w0llWRa4O6f/BhD71ouQIN9bhrQ3/0A7KLRbfysgNZPfa4xsImxnkm+F0sWtHOjxYLmc8yYVmK4iX7hlfLpU0aIpk1QL7ZTqtfRkUfc4yu/LuVHNtkKmZWDKLOgP0h0BZ2CiMY4G7zHgbyt9V0sTOKscxtzkmsgFKAyI49ayNudOTulDD2Mo0Rrm6DH1HxAFbIpWfpwpdBFiYho1eP6A5uyy0AakYOToP7iUH/31VzrvH8131HXV4LqoxT4aQVonlBXZOZHkwdq6zNtkwkJmlo3mm32uVjHCEyhsvXqq+alCN/U1GJ2Wlg0CiK4DuJ337IH1peOA3LKw8wCxF8cNjKcy1hIq6f2+NV58jWHgwjbNZwbK5qnp8FZRzFX76SiWmtzvZyvn5E+ZIjyIV1G74HpRYV")
	str, err := AesECBDecrypt(val, key, PKCS5_PADDING)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(str))
}
