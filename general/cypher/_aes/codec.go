// @author cold bin
// @date 2022/11/16

package _aes

import "github.com/cold-bin/goutil/general/cypher/_codec"

func enBase64(originData []byte) []byte {
	return _codec.EncodeBase64(originData)
}

func deBase64(originData []byte) ([]byte, error) {
	return _codec.DecodeBase64(originData)
}

func enBase64Url(originData []byte) []byte {
	return _codec.EncodeBase64Url(originData)
}

func deBase64Url(originData []byte) ([]byte, error) {
	return _codec.DecodeBase64Url(originData)
}
