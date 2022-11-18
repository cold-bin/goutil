// @author cold bin
// @date 2022/11/16

package _aes

import (
	"crypto/aes"
	"testing"
)

func Test_conf_Crypt(t *testing.T) {
	text := []byte("redrocker501")
	key := []byte("redredshiredred1")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
		return
	}

	for i := 1; i < 5; i++ {
		conf := NewConf(PadModePKCS5, CodecModeBase64, i,
			key, []byte("redredshiredred1"), block)
		decrypt, err := conf.Decrypt(conf.Encrypt(text))
		if err != nil {
			t.Errorf("err: %s", err)
		}
		if err == nil && string(decrypt) != string(text) {
			t.Errorf("not equal. old:%s,new:%s", text, decrypt)
			return
		}
	}
}
