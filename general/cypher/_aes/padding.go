// @author cold bin
// @date 2022/11/15

package _aes

import "bytes"

func paddingPKCS5(cypherText []byte, blockSize int) []byte {
	padLen := blockSize - len(cypherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(cypherText, padText...)
}

func unpaddingPKCS5(originData []byte) []byte {
	length := len(originData)
	padLen := int(originData[length-1])
	return originData[:(length - padLen)]
}

func paddingPKCS7(cypherText []byte, blockSize int) []byte {
	padLen := blockSize - len(cypherText)%blockSize
	padText := bytes.Repeat([]byte{byte(0)}, padLen)
	return append(cypherText, padText...)
}

func unpaddingPKCS7(originData []byte) []byte {
	var index = len(originData)
	for i := len(originData) - 1; i >= 0; i-- {
		if originData[i] != 0 {
			break
		}
		index = i
	}

	return originData[:index]
}
