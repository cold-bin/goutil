// @author cold bin
// @date 2022/11/15

package cypher

// Cypher 加解密
type Cypher interface {
	Encryptor
	Decryptor
}

// Encryptor 加密器
type Encryptor interface {
	Encrypt(src []byte) (dst []byte)
}

// Decryptor 解密器
type Decryptor interface {
	Decrypt(src []byte) (dst []byte, err error)
}
