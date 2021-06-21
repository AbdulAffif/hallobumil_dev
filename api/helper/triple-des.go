package helper

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

type TripleDES struct {
	key string
	iv  string
}

// @ref: https://blog.csdn.net/xiaoxiao_haiyan/article/details/81320350
func (this *TripleDES) Encrypt(plain string) (string, error) {
	key := []byte(this.key)
	iv := []byte(this.iv)

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}
	input := []byte(plain)
	input = PKCS5Padding(input, block.BlockSize())
	blockMode := cipher.NewOFB(block, iv)
	crypted := make([]byte, len(input))
	blockMode.XORKeyStream(crypted, input)

	return base64.StdEncoding.EncodeToString(crypted), err
}

func (this *TripleDES) Decrypt(secret string) (string, error) {
	key := []byte(this.key)
	iv := []byte(this.iv)

	crypted, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewOFB(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.XORKeyStream(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return string(origData), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// remove the last byte unpadding times
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
