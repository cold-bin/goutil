// @author cold bin
// @date 2022/11/15

package _aes

import (
	"crypto/cipher"
	"crypto/rand"
	"github.com/cold-bin/goutil/general/cypher"
)

const (
	EncryptModeCBC = iota + 1
	EncryptModeCTR
	EncryptModeCFB
	EncryptModeOFB
)

const (
	// PadModePKCS5 是指分组数据缺少几个字节，就在数据的末尾填充几个字节的几，比如缺少5个字节，就在末尾填充5个字节的5
	PadModePKCS5 = iota + 1 + EncryptModeOFB
	// PadModePKCS7 是指分组数据缺少几个字节，就在数据的末尾填充几个字节的0，比如缺少7个字节，就在末尾填充7个字节的0
	PadModePKCS7
	// PadModeNoPadding 是指不需要填充，也就是说数据的发送方肯定会保证最后一段数据也正好是16个字节
	PadModeNoPadding
)

const (
	CodecModeBase64 = iota + 1 + PadModeNoPadding
	CodecBaseUrl64
)

type conf struct {

	// padMode 填充数据的模式: PadModePKCS5\ PadModePKCS7 \ PadModeNoPadding
	padMode int

	// codecMode 编码方式: CodecModeBase64 \ CodecBaseUrl64
	codecMode int

	// encryptMode 加密模式:
	// 	CBC(密码分组链接模式)、CTR(计算器模式)、CFB(密码反馈模式)、OFB(输出反馈模式)
	encryptMode int

	// key 密钥，加解密都一样
	key []byte

	// iv 初始向量，与密钥块长度相同，加解密时一样
	iv []byte

	// block 密钥
	block cipher.Block
}

// GetIv 获取加密安全随机数
func GetIv(blockSize int) (iv []byte, err error) {
	iv = make([]byte, blockSize)
	_, err = rand.Read(iv)
	return
}

// New AES加密算法配置。 GetIv 可以快捷生成iv
func New(padMode, codecMode, encryptMode int, key []byte, iv []byte, block cipher.Block) cypher.Cypher {
	if !isKey(key) {
		panic("key len must be 16,24 or 32 ")
	}

	if isPadMode(padMode) && isCodecMode(codecMode) && isEncryptMode(encryptMode) {
		return &conf{
			padMode:     padMode,
			codecMode:   codecMode,
			encryptMode: encryptMode,
			key:         key,
			iv:          iv,
			block:       block,
		}
	}
	return nil
}

// padding 根据不同的填充模式填充
func (a *conf) padding(cypherText []byte, blockSize int) []byte {
	switch a.padMode {
	case PadModePKCS5:
		return paddingPKCS5(cypherText, blockSize)
	case PadModePKCS7:
		return paddingPKCS7(cypherText, blockSize)
	case PadModeNoPadding:
		// 不填充，在此填充下原始数据必须是分组大小的整数倍，非整数倍时无法使用该模式。
		if len(cypherText)%blockSize != 0 {
			panic("not support the padding mode: PadModeNoPadding. cypherText must be integer multiple of blockSize")
		}
		return cypherText
	default:
		panic("not support the padding mode")
	}
}

// unpadding 根据不同的填充模式去填充
func (a *conf) unpadding(originData []byte) []byte {
	switch a.padMode {
	case PadModePKCS5:
		return unpaddingPKCS5(originData)
	case PadModePKCS7:
		return unpaddingPKCS7(originData)
	case PadModeNoPadding:
		return originData
	default:
		panic("not support the padding mode")
	}
}

func (a *conf) enBase(originData []byte) []byte {
	switch a.codecMode {
	case CodecModeBase64:
		return enBase64(originData)
	case CodecBaseUrl64:
		return enBase64Url(originData)
	default:
		panic("not support the codec mode")
	}
}

func (a *conf) deBase(originData []byte) ([]byte, error) {
	switch a.codecMode {
	case CodecModeBase64:
		return deBase64(originData)
	case CodecBaseUrl64:
		return deBase64Url(originData)
	default:
		panic("not support the codec mode")
	}
}

// Encrypt 加密
func (a *conf) Encrypt(originData []byte) (encrypted []byte) {
	padData := a.padding(originData, a.block.BlockSize()) // 填充

	switch a.encryptMode {
	case EncryptModeCBC:
		encrypted = encryptCBC(padData, a.key, a.block)
	case EncryptModeCFB:
		encrypted = encryptCFB(padData, a.key, a.block)
	case EncryptModeCTR:
		encrypted = encryptCTR(padData, a.iv, a.block)
	case EncryptModeOFB:
		encrypted = encryptOFB(padData, a.iv, a.block)
	default:
		panic("not support the encrypt mode")
	}
	return a.enBase(encrypted)
}

// Decrypt 解密
func (a *conf) Decrypt(encrypted []byte) (decrypted []byte, err error) {
	if encrypted, err = a.deBase(encrypted); err != nil {
		return
	}
	switch a.encryptMode {
	case EncryptModeCBC:
		decrypted = decryptCBC(encrypted, a.key, a.block)
	case EncryptModeCFB:
		decrypted = decryptCFB(encrypted, a.key, a.block)
	case EncryptModeCTR:
		decrypted = decryptCTR(encrypted, a.iv, a.block)
	case EncryptModeOFB:
		decrypted = decryptOFB(encrypted, a.iv, a.block)
	default:
		panic("not support the encrypt mode")
	}

	decrypted = a.unpadding(decrypted) // 去除补全码
	return
}
