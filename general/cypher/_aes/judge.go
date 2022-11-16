// @author cold bin
// @date 2022/11/16

package _aes

func isKey(key []byte) bool {
	return len(key) == 16 || len(key) == 24 || len(key) == 32
}

func isPadMode(m int) bool {
	return m == PadModePKCS5 || m == PadModePKCS7 || m == PadModeNoPadding
}

func isCodecMode(m int) bool {
	return m == CodecBaseUrl64 || m == CodecModeBase64
}

func isEncryptMode(m int) bool {
	return m == EncryptModeCTR || m == EncryptModeCBC || m == EncryptModeOFB || m == EncryptModeCFB
}
