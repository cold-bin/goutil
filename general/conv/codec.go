// @author cold bin
// @date 2022/11/11

package conv

import "encoding/base64"

// EncodeBase64 标准的base64编码
func EncodeBase64(src []byte) (dst []byte) {
	dst = make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return
}

// DecodeBase64 标准的base64解码
func DecodeBase64(src []byte) (dst []byte, err error) {
	dst = make([]byte, base64.StdEncoding.DecodedLen(len(src)))
	n, e := base64.StdEncoding.Decode(dst, src)
	return dst[:n], e
}

// EncodeBase64Url 标准的base64 url编码
func EncodeBase64Url(src []byte) (dst []byte) {
	dst = make([]byte, base64.URLEncoding.EncodedLen(len(src)))
	base64.URLEncoding.Encode(dst, src)
	return
}

// DecodeBase64Url 标准的base64 url解码
func DecodeBase64Url(src []byte) (dst []byte, err error) {
	dst = make([]byte, base64.URLEncoding.DecodedLen(len(src)))
	n, e := base64.URLEncoding.Decode(dst, src)
	return dst[:n], e
}
