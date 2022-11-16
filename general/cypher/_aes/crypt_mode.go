// @author cold bin
// @date 2022/11/16

package _aes

import (
	"crypto/cipher"
)

func encryptCBC(padData []byte, key []byte, block cipher.Block) (encrypted []byte) {
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()]) // 加密模式
	encrypted = make([]byte, len(padData))                              // 创建数组
	blockMode.CryptBlocks(encrypted, padData)                           // 加密
	return
}

func decryptCBC(encrypted []byte, key []byte, block cipher.Block) (decrypted []byte) {
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	return
}

func encryptCFB(padData []byte, key []byte, block cipher.Block) (encrypted []byte) {
	blockMode := cipher.NewCFBEncrypter(block, key[:block.BlockSize()]) // 加密模式
	encrypted = make([]byte, len(padData))                              // 创建数组
	blockMode.XORKeyStream(encrypted, padData)                          // 加密
	return
}

func decryptCFB(encrypted []byte, key []byte, block cipher.Block) (decrypted []byte) {
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCFBDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.XORKeyStream(decrypted, encrypted)                // 解密
	return
}

func encryptCTR(padData []byte, iv []byte, block cipher.Block) (encrypted []byte) {
	blockSize := block.BlockSize()
	ciphertext := make([]byte, blockSize+len(padData))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[blockSize:], padData)
	return ciphertext[blockSize:]
}

func decryptCTR(encrypted []byte, iv []byte, block cipher.Block) (decrypted []byte) {
	blockSize := block.BlockSize()
	ciphertext := make([]byte, blockSize+len(encrypted))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[blockSize:], encrypted)
	return ciphertext[blockSize:]
}

func encryptOFB(padData []byte, iv []byte, block cipher.Block) (encrypted []byte) {
	blockSize := block.BlockSize()
	ciphertext := make([]byte, blockSize+len(padData))
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(ciphertext[blockSize:], padData)
	return ciphertext[blockSize:]
}

func decryptOFB(encrypted []byte, iv []byte, block cipher.Block) (decrypted []byte) {
	blockSize := block.BlockSize()
	ciphertext := make([]byte, blockSize+len(encrypted))
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(ciphertext[blockSize:], encrypted)
	return ciphertext[blockSize:]
}
