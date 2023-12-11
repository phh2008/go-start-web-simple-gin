package aes

import (
	"com.gientech/selection/util"
	"testing"
)

func TestAesGcm(t *testing.T) {
	key := []byte("7xhsdawlNNvnbytNyUFmmRLhRmJ86uDr")
	iv := []byte(util.RandStr(16))
	originPlainText := []byte("君不见黄河之水天上来，奔流到海不复回。君不见高堂明镜悲白发，朝如青丝暮成雪。人生得意须尽欢，莫使金樽空对月。五花马，千金裘，呼儿将出换美酒。与尔同消万古愁。岑夫子丹丘生，将进酒，杯莫停，与君歌一曲，请军为我倾耳听，钟鼓馔玉不足贵，但愿长醉不复醒。")
	aesGcm := NewAesGCM(key, iv...)
	cipherText, err := aesGcm.EncryptBase64(originPlainText)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(cipherText))

	plainText, err := aesGcm.DecryptBase64(cipherText)
	if err != nil || string(plainText) != string(originPlainText) {
		t.Error(err)
	}
	t.Log(string(plainText))

	cipherText, err = aesGcm.EncryptBase64WithIV(originPlainText)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(cipherText))

	plainText, err = aesGcm.DecryptBase64WithIV(cipherText)
	if err != nil || string(plainText) != string(originPlainText) {
		t.Error(err)
	}
	t.Log(string(plainText))
}

func TestRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(util.RandStr(32))
	}

}

func TestAesGcm2(t *testing.T) {
	key := []byte("7xhsdawlNNvnbytNyUFmmRLhRmJ86uDr")
	iv := []byte(util.RandStr(16))
	t.Log("iv: ", string(iv))
	gcm := NewAesGCM(key, iv...)
	data := "cjE5UGlYNUdrUFMzYVhyUDB+45NDd0n1ZIuPWJTfgidIxix36qcDBWpFds/nQiIGn2aEDabhERhuGSlAJS82+BcAVHPvfg2Ap5X8jihz6jxUkIXFMDUBzR5IGjQRoALQWJXdk0xflD9+1hWLHAiz5WskU6tOiAAfdwBEpNykaMQ0R0NEkwoeSgbDzskrKBVudKOMh5oyhg8GzhmGrA7xRMeiEsbykvRsukEbQwFBidpC9cwTzG/kTQ13KwIm854BHu8sQiikHZAr5bwYP/kcMBb1uTPk2JvPOtlzNgm+sRfwwtxUBa9wl4xx5KPL7vpdNBNfRGVXde+JEJuUBhBwHw7gMZaEt3jmE9CCrxTLfIyISjkVxZl/DwFL9o+nvQBQXdO6/9+/Wc1dKDd6h++hkO6c9Qr3qK5rrSyHOwHUTspEvEUzcRhinaKN7V6K4XiZFOL9oqrcxDv7S5XXGwTmv0t+M4zY7sGQdZNL4eIe+Azxlb2AydBxC0U3WuHaFccFriygPtog8SapLR7OP00qNpE="
	plainText, err := gcm.DecryptBase64WithIV(data)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(plainText))
}
