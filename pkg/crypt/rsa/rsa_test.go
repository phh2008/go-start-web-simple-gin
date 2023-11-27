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

func TestEncrypt(t *testing.T) {
	var src = []byte("山不在高，有仙则名。水不在深，有龙则灵。斯是陋室，惟吾德馨。苔痕上阶绿，草色入帘青。谈笑有鸿儒，往来无白丁。")
	var pubKey = "MIGJAoGBAKT1PU6Gw0OiL5WACDq6UT9cCM9ya6fMmKB0n6PbU1jt3kOQCtUnurzXk3QR08fWo+LKxnV1VVul0io4nbYKT+upsIEo1H3LAw9gSqQvv50CIh7KBomgqlsuzea7ehmCTRmhzwRhftzY+eT9rpv036kqw8UkkUje0DfXQDUeAU4rAgMBAAE="
	enc, err := EncryptByPubKey(src, pubKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(enc)
}

func TestDecrypt(t *testing.T) {
	var src = "UrdDbY4IbPOFcdfZqhZs0VWbO6uur8lKy3Gci36WJHfKvDsZNWveBGtI8a4MwcPLQuVYnbXWNroXdegHdyUvm0EOLCVhhDHQFWRlalv2zrbZkmYhVuaOVSsIDDhbizQN1KXgbWhI224IBst2qEorGY8jUWnTZorKYGNkHPGAsstS7dzW/XW6TPQnijRBxcHkt8UpoW0/p4JXwGC6BmmmhV5I1g+vunNFd1ZouU/7HXCLkUruF4TyNEo6vi6JbFTWAVuaOIBl7nAHV+6OK+kAoQ6t+HvhBzg2YdCTb/+ePYr+JrlZF1TrmlZi5sSw9dQ5nRO1I3cHh0Aw7nYJwzM8Hg=="
	var priKey = "MIICXAIBAAKBgQCk9T1OhsNDoi+VgAg6ulE/XAjPcmunzJigdJ+j21NY7d5DkArVJ7q815N0EdPH1qPiysZ1dVVbpdIqOJ22Ck/rqbCBKNR9ywMPYEqkL7+dAiIeygaJoKpbLs3mu3oZgk0Zoc8EYX7c2Pnk/a6b9N+pKsPFJJFI3tA310A1HgFOKwIDAQABAoGAYxKLdJtJsVg6Xf7ccnEulPAwtm8RK2GdFVmV+7Khd1q90DQ13VmUNREAlYiTeoV2PqLs2OTUlEueUw9X4VqGLpveD2Sip9WfR6+b5eMlM3STHdB5plXqCnHIq6o7VH4iueqMG9TxpQPhDLt2VEaZUooCxWcaaSfiLLmmqV8Po+ECQQDFlyrxULS89FHlpWpkEnLLEawJUFS4/Jd7+DW+YSl7ihLIqRnt3qxjDmXw9ZGtrMlRYBkyzikpPyIcJ/rHfz2xAkEA1biTmka0Rezipku47VO4KFyZkhoORlCHWZtjMjb8xwxWn4/6nQ8Zvt+sutGdE/8C7Curnri8RpUAthwx8x80mwJAA7k0ivWdYk4sWOqEFbyvQxpjJo3H+vBvnltwD9Ve5cAVWIivP2dJ0lgMHb1S8HXoGUt5ThbKeceBygwK0sWYEQJBANKs1zt+NRi08ZtSC6JPI7sNxQijjKy9lx66sSb/01/3hrBVworuJsfkP6YNGRVsDRp31f5pRpchLIlX89kgSr8CQAWfIbobVIJfw9xFKJJgfZPNbSKY56w4XgrDrhCWChXSZqBt/ENGgfCLULXuAsdJH0Fr82Rzx32/9xMpC7C+dIU="
	ret, err := DecryptByPriKey(src, priKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ret))
}
