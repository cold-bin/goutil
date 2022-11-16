// @author cold bin
// @date 2022/11/15

package _codec

import "encoding/base64"

func EncodeBase64(src []byte) (dst []byte) {
	dst = make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return
}

func DecodeBase64(src []byte) (dst []byte, err error) {
	dst = make([]byte, base64.StdEncoding.DecodedLen(len(src)))
	n, e := base64.StdEncoding.Decode(dst, src)
	return dst[:n], e
}

func EncodeBase64Url(src []byte) (dst []byte) {
	dst = make([]byte, base64.URLEncoding.EncodedLen(len(src)))
	base64.URLEncoding.Encode(dst, src)
	return
}

func DecodeBase64Url(src []byte) (dst []byte, err error) {
	dst = make([]byte, base64.URLEncoding.DecodedLen(len(src)))
	n, e := base64.URLEncoding.Decode(dst, src)
	return dst[:n], e
}
